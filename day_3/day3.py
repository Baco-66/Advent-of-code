import re


def parse_data_from_file(filename):
    """Reads data from a file and returns a list of lists"""
    result = []
    pattern = r"mul\((\d+),(\d+)\)"

    with open(filename, 'r') as file:
            # Perform search
        for line in file:
            result += re.findall(pattern, line)
    return result

def parse_data_from_file_2(filename):
    """Reads data from a file and returns a list of lists"""
    data = ""
    searching = True
    with open(filename, 'r') as file:
            # Perform search
        for line in file:
            line, searching = process_line_with_regex(line, searching)
            data += line
    print(data)

    pattern = r"mul\((\d+),(\d+)\)"
    return re.findall(pattern, data)


def process_line_with_regex(line, register=True):
    # Initialize the output and a cursor to keep track of positions
    output = []
    
    # Use regex to match all occurrences of "don't()" or "do()"
    matches = re.finditer(r"(don't\(\)|do\(\))", line)
    last_pos = 0

    for match in matches:
        print(match)
        if register:
            # Append the part of the line before the match if registering
            output.append(line[last_pos:match.end()])
        
        # Update the register state based on the match
        if match.group() == "don't()":
            register = False
        elif match.group() == "do()":
            register = True

        # Update the last position
        last_pos = match.start()

    # Append the remainder of the string if registering
    if register:
        output.append(line[last_pos:])

    # Combine the registered parts
    return ''.join(output),register

def main():
    print("Advent of code day 1!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query? ")
    print("")

    match query:
        case "1":
            data = parse_data_from_file('data.txt')

            result = 0
            for (right, left) in data:
                result += (int(right)*int(left))
            print(result)
        case "2":
            data = parse_data_from_file_2('data.txt')

            result = 0
            for (right, left) in data:
                result += (int(right)*int(left))
            print(result)            
        case _:
            print("Invalid option!")


    print("That's all folks!")
    return 0


if __name__ == '__main__':
    main()
