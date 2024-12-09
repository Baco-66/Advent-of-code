import cProfile
import itertools

def parse_data_from_file(filename):
    """Reads data from a file and returns a list of lists"""
    with open(filename, 'r') as file:
        return [[char for char in line if char != "\n"] for line in file]
    
def print_matrix(matrix):
    for row in matrix:
        print(' '.join(map(str, row)))  # Convert each element to a string and join with spaces

def find_positions_in_matrix(matrix, search):
    return [[x, y] for x, line in enumerate(matrix) for y, value in enumerate(line) if value == search]

def walk(map, size, position, direction): # After optmizing the loop ckeck, the main botleneck is the walk function. 
    # Try 1: avoid cheking len constantly by only doing it once
    # 1.64563 s -> 1.49036 s
    # Try 2: reduce the if statement
    nx, ny = position[0] + direction[0], position[1] + direction[1]

    # Check bounds
#   if 0 <= nx < len(map) and 0 <= ny < len(map[0]): 
    #if 0 <= nx < size[0] and 0 <= ny < size[1]: 
    #    if map[nx][ny] == "#":  # Check if the target is blocked
    #        return 0  # Cannot move to the new position
    #    else:
    #        position[0], position[1] = nx, ny  # Update position
    #        return 1  # Move successful
    #else:
    #    return 2  # Out of bounds
    
    # Boundary check
    if not (0 <= nx < size[0] and 0 <= ny < size[1]):
        return 2  # Out of bounds
    # Check if blocked
    if map[nx][ny] == "#":
        return 0  # Cannot move
    # Move successful
    position[0], position[1] = nx, ny
    return 1

def patrol(map, size, position):
    # Create an infinite iterator
    directions = itertools.cycle([
        [-1, 0],  # Up
        [0, 1],   # Right
        [1, 0],   # Down
        [0, -1],  # Left
    ])
    
    for direction in directions:
        # Continue walking while the step is valid (1) and stop if out of bounds (2)
        while (step := walk(map, size, position, direction)) == 1:
            map[position[0]][position[1]] = "X"  # Mark old position before attempting to walk

        # Stop patrolling if out of bounds
        if step == 2:
            break

def check_for_looped_path(path): # REALY SLOW, 2,5 seconds
    
    # Check if the list has at least 2 elements
    if len(path) < 2:
        return False

    # Get the last two numbers
    last_two = path[-2:]

    # Iterate through the list and check if the last two numbers appear together elsewhere
    for i in range(len(path) - 2):
        if path[i:i+2] == last_two:
            return True

    return False

def check_for_looped_path_faster(path): # one second faster
    # Check if the list has at least 2 elements
    if len(path) < 2:
        return False

    # Get the last two elements
    last_two = path[-2:]

    # Iterate through the list and check if the last two numbers appear together elsewhere
    for i in range(len(path) - 2):
        if path[i] == last_two[0] and path[i+1] == last_two[1]:
            return True

    return False

def check_for_looped_path_FASTER(path): # WTFFF reduzi para 0.023 sec
    # Check if the list has at least 2 elements
    if len(path) < 2:
        return False

    # Get the last two pairs
    last_two = path[-2:]

    # Iterate backwards and compare the pairs
    for i in range(len(path) - 3, -1, -1):
        if path[i] == last_two[0] and path[i+1] == last_two[1]:
            return True

    return False

def find_loops(map, size, position, obstacle):
    directions = itertools.cycle([
        [-1, 0],  # Up
        [0, 1],   # Right
        [1, 0],   # Down
        [0, -1],  # Left
    ])
    loop = set()
    toutched = False # added this and managed to reduce it to 0.040 seconds
    # I was checking the path constantly to see if to tiles repeated. The obvious solution was to go and only check if i have hit a wall or not and verify that

    for direction in directions:
        # Continue walking while the step is valid (1) and stop if out of bounds (2)
        while (step := walk(map, size, position, direction)) == 1:
            pass
        if step == 0:
            if toutched:
                magic = ((position[0], position[1], tuple(direction)))
                if magic in loop :
                    #print("Obstacle", obstacle, "caused the loop",loop)
                    return True
                loop.add(magic)

            elif obstacle == [position[0] + direction[0], position[1] + direction[1]]:
                magic = ((position[0], position[1], tuple(direction)))
                loop.add(magic)
                toutched = True

        # Stop patrolling if out of bounds
        elif step == 2:
            return False

    return False


def slow_function():
    # Example usage

    file_path = 'data.txt'
    map = parse_data_from_file(file_path) 
    size = (len(map), len(map[0]))

    position = find_positions_in_matrix(map, "^")[0]

    patrol(map, size, position.copy())
    map[position[0]][position[1]]= "^"

    path = find_positions_in_matrix(map, "X")

    #result = []
    result = 0


    for i, obstacle in enumerate(path):
        #print("Done ", i, "with obstacle", obstacle)
        map[obstacle[0]][obstacle[1]]= "#"
        if find_loops(map, size, position.copy(), obstacle.copy()):
            #result.append(obstacle)
            result += 1
        map[obstacle[0]][obstacle[1]]= "."
        #if i == 100:
        #    break
    #for x,y in result:
    #    map[x][y]="O"
    #print_matrix(map)
    print("There are",result, "blocks that create a loop!")

#slow_function()

cProfile.run('slow_function()')

# Final one takes about 20 seconds. Comparing to the original 
# which did not actualy finish its a tremendous amout. 
# In order to reduce more i would have to change the way I 
# store the map and change the way the walk function is used 