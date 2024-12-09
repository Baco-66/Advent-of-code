
matches = []

def parse_data_from_file(filename):
    """Reads data from a file and returns a list of lists"""
    with open(filename, 'r') as file:
        return [[char for char in line if char != "\n"] for line in file]

def print_matrix(matrix):
    for row in matrix:
        print(' '.join(map(str, row)))  # Convert each element to a string and join with spaces

def match_xmas_direction(matrix, x, y, dx, dy):
    """Check if the word 'XMAS' is found starting from (x, y) in the direction (dx, dy)."""
    # Positions for "M", "A", "S" after the initial "X"
    positions = [(x + i * dx, y + i * dy) for i in range(1, 4)]  # Checking for "M", "A", "S"

    # Check if all positions are valid and contain the respective letters "M", "A", "S"
    if all(0 <= nx < len(matrix) and 0 <= ny < len(matrix[0]) and matrix[nx][ny] == "MAS"[i-1]
           for i, (nx, ny) in enumerate(positions, 1)):  # Start from i=1 for "M"
        #print(f"Match found at positions: {(x, y)} -> {positions}")  # Debug print
        matches.extend([(x, y)] + positions)  # Store the positions
        return 1
    return 0

def match_xmas_all_directions(matrix, x, y):
    """Check all 8 directions from (x, y) for the word 'MAS'."""
    directions = [
        (-1, 0),  # Up
        (1, 0),   # Down
        (0, 1),   # Right
        (0, -1),  # Left
        (-1, -1), # Up-left
        (-1, 1),  # Up-right
        (1, -1),  # Down-left
        (1, 1),   # Down-right
    ]
    
    matches = 0
    for dx, dy in directions:
        matches += match_xmas_direction(matrix, x, y, dx, dy)
    return matches

def match_mas_direction(matrix, x, y, dx, dy):
    """Check if the word 'MAS' is found starting from (x, y) in the direction (dx, dy)."""
    if (0 <= x + dx < len(matrix) and 0 <= y + dy < len(matrix[0]) and matrix[x + dx][y + dy] == "M"):
        if (0 <= x - dx < len(matrix) and 0 <= y - dy < len(matrix[0]) and matrix[x - dx][y - dy] == "S"):
            return 1
    return 0

def match_mas(matrix, x, y):
    """Check all 4 directions from (x, y) for the word 'MAS' in a cross."""
    positions = [
        (x, y),
        (x + 1, y + 1),
        (x - 1, y + 1),
        (x + 1, y - 1),
        (x - 1, y - 1)
    ]
    #directions = [
    #    (-1, -1), # Up-left
    #    (-1, 1),  # Up-right
    #    (1, -1),  # Down-left
    #    (1, 1),   # Down-right
    #]
    if match_mas_direction(matrix, x, y, -1, -1) or match_mas_direction(matrix, x, y, 1, 1):
        if match_mas_direction(matrix, x, y, 1, -1) or match_mas_direction(matrix, x, y, -1, 1):
            #print(f"Match found at positions: {positions}")  # Debug print
            #matches.extend(positions)
            return 1
    return 0

def debug_matrix(matrix):
    # Replace unmatched characters with dots
    for x in range(len(matrix)):
        for y in range(len(matrix[x])):
            if (x, y) not in matches:
                matrix[x][y] = '.'
    
    return(matrix)

def find_positions_in_matrix(matrix, search):
    return [(x, y) for x, line in enumerate(matrix) for y, value in enumerate(line) if value == search]

def main():
    # Example usage
    file_path = 'data.txt' 

    print("Advent of code day 1!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")
    print("")


    char_matrix = parse_data_from_file(file_path)

    


    match query:
        case "1":
            positions = find_positions_in_matrix(char_matrix, "X")
            result = 0
            for (x,y) in positions:
                result += match_xmas_all_directions(char_matrix, x, y)
            print("There are", result, "matches!")
            
            #with open("output.txt", 'w') as f:
            #    for row in debug_matrix(char_matrix):
            #        f.write(''.join(row) + '\n')  # Join the elements in each row to form a string and write

        case "2":
            positions = find_positions_in_matrix(char_matrix, "A")
            result = 0
            for (x,y) in positions:
                result += match_mas(char_matrix, x, y)
            print("There are", result, "matches!")

            #with open("output.txt", 'w') as f:
            #    for row in debug_matrix(char_matrix):
            #        f.write(''.join(row) + '\n')  # Join the elements in each row to form a string and write
        case _:
            print("Invalid option!")


    print("That's all folks!")
    return 0


if __name__ == '__main__':
    main()
