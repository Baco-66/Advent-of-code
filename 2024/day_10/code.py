import itertools

def parse_data_from_file(filename):
    """Reads data from a file and returns a list of lists"""
    with open(filename, 'r') as file:
        return [[int(char) for char in line if char != "\n"] for line in file]

def find_positions_in_matrix(matrix, search):
    return [[x, y] for x, line in enumerate(matrix) for y, value in enumerate(line) if value == search]

def print_matrix(matrix):
    for row in matrix:
        print(' '.join(map(str, row)))  # Convert each element to a string and join with spaces

def hike(map, size, position):
    # Create an infinite iterator
    directions = [
        [-1, 0],  # Up
        [0, 1],   # Right
        [1, 0],   # Down
        [0, -1],  # Left
    ]
    hight = map[position[0]][position[1]]

    result = set()

    if hight == 9:
        result.add(position)
        return result
    
    for direction in directions:
        # Continue walking while the step is valid (1) and stop if out of bounds (2)
        nx, ny = position[0] + direction[0], position[1] + direction[1]
        if 0 <= nx < size[0] and 0 <= ny < size[1] and map[nx][ny] == (map[position[0]][position[1]] + 1):
            result = result.union(hike(map, size, (nx, ny)))

    return result

def hike2(map, size, position):
    # Create an infinite iterator
    directions = [
        [-1, 0],  # Up
        [0, 1],   # Right
        [1, 0],   # Down
        [0, -1],  # Left
    ]
    hight = map[position[0]][position[1]]

    result = []

    if hight == 9:
        result.append(position)
        return result
    
    for direction in directions:
        # Continue walking while the step is valid (1) and stop if out of bounds (2)
        nx, ny = position[0] + direction[0], position[1] + direction[1]
        if 0 <= nx < size[0] and 0 <= ny < size[1] and map[nx][ny] == (map[position[0]][position[1]] + 1):
            result.extend(hike2(map, size, (nx, ny)))

    return result

def main():
    # Example usage
    file_path = '2024/day_10/test_data.txt'
    file_path = '2024/day_10/data.txt' 

    print("Advent of code day 10!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")

    data = parse_data_from_file(file_path)

    match query:
        case "1":
            #result = []
            score = 0
            trailheads = find_positions_in_matrix(data,0)
            for trailhead in trailheads:
                #result.append(hike(data,(len(data),len(data[0])),trailhead))
                trail = hike(data,(len(data),len(data[0])),trailhead)
                score += len(trail)
            print(score)
        case "2":
            score = 0
            trailheads = find_positions_in_matrix(data,0)
            for trailhead in trailheads:
                #result.append(hike(data,(len(data),len(data[0])),trailhead))
                trail = hike2(data,(len(data),len(data[0])),trailhead)
                #print(trail, len(trail))
                score += len(trail)
            print(score)
        case _:
            print("Invalid option!")

    print("That's all folks!")
    return 0

if __name__ == '__main__':
    main()
