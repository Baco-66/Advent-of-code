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

def walk(map, size, position, direction):
    nx, ny = position[0] + direction[0], position[1] + direction[1]
    # Check bounds
    if 0 <= nx < size[0] and 0 <= ny < size[1]: 
        if map[nx][ny] == "#":  # Check if the target is blocked
            return 0  # Cannot move to the new position
        else:
            position[0], position[1] = nx, ny  # Update position
            return 1  # Move successful
    else:
        return 2  # Out of bounds

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

def find_loops(map, size, position, obstacle):
    directions = itertools.cycle([
        [-1, 0],  # Up
        [0, 1],   # Right
        [1, 0],   # Down
        [0, -1],  # Left
    ])
    loop = set()
    toutched = False 

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

def main():
    # Example usage
    file_path = '2024/day_6/test_data.txt' 
    file_path = '2024/day_6/data.txt' 

    print("Advent of code day 6!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")

    map = parse_data_from_file(file_path)
    size = (len(map), len(map[0]))

    position = find_positions_in_matrix(map, "^")[0]
    patrol(map, size, position.copy())
    map[position[0]][position[1]]= "^"



    match query:
        case "1":
            print("Result is ",len(find_positions_in_matrix(map, "X")))
        case "2":
            path = find_positions_in_matrix(map, "X")
            result = 0
            for i, obstacle in enumerate(path):
                print("Done ", i, "with obstacle", obstacle)
                map[obstacle[0]][obstacle[1]]= "#"
                if find_loops(map, size, position.copy(), obstacle.copy()):
                    result += 1
                map[obstacle[0]][obstacle[1]]= "."

            print("There are",result, "blocks that create a loop!")
        case _:
            print("Invalid option!")


    print("That's all folks!")
    return 0


if __name__ == '__main__':
    main()
