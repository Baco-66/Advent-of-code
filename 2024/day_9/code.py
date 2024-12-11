import cProfile

def read_file_character_by_character(filename):
    with open(filename, 'r') as file:
        while (char := file.read(1)):  # Read one character at a time
            if char == "\n":
                break
            yield char  # Use yield to return the character for iteration

def parse_disk_map(file_path):
    data = []
    is_file = True
    id = 0
    id_char = ""
    for char in read_file_character_by_character(file_path):
        if is_file:
            id_char = str(id)
            id += 1
        else:
            id_char = "."
        is_file = not is_file  
        count = int(char)
        data.extend([id_char] * count)
    return data

def organize_files_one(data):
    i = 0
    j = len(data)-1
    while i != j:
        if data[i] == ".":
            if data[j] != ".":
                data[i] = data[j]
                data[j] = "."
                i += 1
            j -= 1
        else:
            i += 1
            if data[j] == ".":
                j -= 1 
    return data

def calculate_checksum(data):
    result = 0

    for index, char in enumerate(data):
        if char != ".":
            result += index * int(char)
    return result

#@profile
def move_file(data, id):
    # id = [start_index, end_index, value_at_end_index]  
    free = 0
    files = id[1] - id[0]
    for index in range(len(data) - 1):
        if data[index] == ".":
            free += 1
        else:
            free = 0  # Reset free when a non-empty spot is encountered
        
        if free > files:  # If enough free spaces are found
            # Start the move process
            file = id[1]
            if index > file:
                break
            # Move the range of elements from the `id[1]` to `id[0]`
            for move_index in range(index, index - files - 1, -1):
                data[move_index] = data[file]
                data[file] = "."
                file -= 1  # Move file backward
            break  # Exit after the move is done
    return data

def organize_files_two(data):

    for file in find_files_reverse(data.copy()):
        data = move_file(data, file)
    return data

def find_files_reverse(data):
    last_id = None
    for index in range(len(data) - 1, -1, -1):
        if data[index] != ".":
            if last_id is None:  
                last_id = [index, index, data[index]]  # First case
            elif data[index] == last_id[2]: 
                last_id[0] = index  # Extend the range backwards
            else:
                yield last_id
                last_id = [index, index, data[index]]  # Start a new range
    if last_id:  # Yield the last found sequence if exists
        yield last_id

def main():
    # Example usage
    file_path = '2024/day_9/test_data.txt'
    file_path = '2024/day_9/data.txt' 

    print("Advent of code day 9!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = "2" #input("What is the query?\n")

    data = parse_disk_map(file_path)

    match query:
        case "1":
            data = organize_files_one(data)

            result = calculate_checksum(data)

            print(result)
        case "2":
            data = organize_files_two(data)
            #print(data)
            result = calculate_checksum(data)

            #print(result)
        case _:
            print("Invalid option!")

    print("That's all folks!")
    return 0

if __name__ == '__main__':
    main()
#    cProfile.run('main()')

#main()