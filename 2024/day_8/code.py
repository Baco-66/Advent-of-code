from collections import defaultdict

def parse_data_from_file(filename):
    """Reads data from a file and returns a list of lists"""
    with open(filename, 'r') as file:
        return [[char for char in line if char != "\n"] for line in file]

def print_matrix(matrix):
    for row in matrix:
        print(' '.join(map(str, row)))  # Convert each element to a string and join with spaces

def separate_by_frequency(values):
    result = defaultdict(list)  # Create a dictionary to store lists
    for a, x, y in values:
        result[a].append((x, y))  # Group by 'a' and append (x, y)
    return dict(result)  # Convert defaultdict to a regular dict for cleaner output

def get_antenna_location(grid):
    # Return all non-dot values as tuples of (value, x, y)
    return [(value, int(x), int(y)) for x, row in enumerate(grid) for y, value in enumerate(row) if value != "."]

def get_vector(point_a, point_b):
    return (point_b[0] - point_a[0], point_b[1] - point_a[1])

def add_vector(point, vector):
    return (point[0] - vector[0], point[1] - vector[1])

def main():
    # Example usage
    file_path = '2024/day_8/test_data.txt'
    file_path = '2024/day_8/data.txt' 

    print("Advent of code day 8!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")

    map = parse_data_from_file(file_path)
    size = (len(map), len(map[0]))

    antenas = separate_by_frequency(get_antenna_location(map))
    antinodes = []


    match query:
        case "1":
            for locations in antenas.values():
                for origin in locations:
                    for location in locations:
                        if origin != location:
                            vector = get_vector(location, origin)
                            antinode = add_vector(location, vector)
                            if  0 <= antinode[0] < size[0] and 0 <= antinode[1] < size[1] and antinode not in antinodes:
                                #map[antinode[0]][antinode[1]] = "#"
                                #print("Vector",vector, "created from points", origin, location, "to generate", antinode)
                                antinodes.append(antinode)
                            #else:
                            #    print("Vector",vector, "created from points", origin, location, "is out of bounds", antinode)
            #print_matrix(map)
            print(len(antinodes))
        case "2":
            for locations in antenas.values():
                print(locations)
                for origin in locations:
                    for location in locations:
                        if origin != location:
                            vector = get_vector(location, origin)
                            antinode = location
                            while 0 <= antinode[0] < size[0] and 0 <= antinode[1] < size[1]:
                                if antinode not in antinodes:
                                    antinodes.append(antinode)
                                antinode = add_vector(antinode, vector)
            print(len(antinodes))
        case _:
            print("Invalid option!")

    print("That's all folks!")
    return 0

if __name__ == '__main__':
    main()
