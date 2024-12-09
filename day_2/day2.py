
def read_data_from_file(filename):
    """Reads data from a file and returns a list of lists"""
    result = []
    with open(filename, 'r') as file:
        for line in file:
            line_data = []
            line = line.split()
            for word in line:
                line_data.append(int(word))
            result.append(line_data)
    return result

def number_check(right, left):
    difference = right - left
    return 1 <= difference <= 3

def safety_test(line: list) -> bool:

    if len(line) > 2:
        if line[0] > line[1]:
            index = 0
            while index < len(line)-1:
                if line[index]-line[index+1] > 3 or line[index]-line[index+1] < 1:
                    # print("Line is not safe because", line[index], "is not larger than", line[index+1], "by the minimum of 3")
                    # print("Diference is", line[index]-line[index+1])
                    return False
                index += 1
        elif line[0] < line[1]:
            index = 0
            while index < len(line)-1:
                if line[index+1]-line[index] > 3 or line[index+1]-line[index] < 1:
                    # print("Line is not safe because", line[index], "is not smaller than", line[index+1], "by the minimum of 3")
                    # print("Diference is", line[index+1]-line[index])
                    return False
                index += 1
        else:
            # print("Line is not safe because", line[0], "is equal to", line[1])
            return False

    # print("Safe line found! Line is:\n", line)
    return True

def safety_test_with_dampner(line: list, retry: bool=False):
    coments = False
    descending = True
    if 6 in line and 8 in line and 10 in line and 13 in line and 15 in line:
        print(line)
        coments = True
    if len(line) > 2:
        if line[0] == line[1]:
            if retry:
                #print("Line is", line, "and is not safe because", line[0], "is equal to", line[1])
                if coments: print("Failed in",line, line[0])
                return 0
            else:
                #print("Used the extra life! Line was", line)
                newline = line.copy()
                newline.pop(0)
                
                if safety_test_with_dampner(newline,True):
                    return 2
                else:
                    if coments: print("Failed in",line, line[0])
                    return 0
        elif line[0] < line[1]:
            descending = False
    else:
        return 1
            
    index = 0
    while index < len(line)-1:
        if descending:
            passed = number_check(line[index],line[index+1])
        else:
            passed = number_check(line[index+1],line[index])
        if not passed:
            if retry:
                #print("Line is not safe because", line[index], "and", line[index+1], 
                #        "have a diference smaller than 1 of", diference,)
                if coments: print("Failed in",line,line[index])
                return 0
            else:
                # Attempt to remove elements at `index - 1`, `index`, and `index + 1`, if they exist
                for offset in [-1, 0, 1]:
                    # Check if the offset leads to a valid index
                    target_index = index + offset
                    if 0 <= target_index < len(line):
                        # Make a copy of the line and remove the element at the calculated index
                        modified_line = line.copy()
                        modified_line.pop(target_index)

                        # Test the modified line
                        if safety_test_with_dampner(modified_line, True):
                            if coments: 
                                print("Exited True", line, "after removing", line[target_index])
                            return 2

                # If no valid modification works, return 0
                if coments:
                    print("Failed in", line, line[index])
                return 0

        else:
            index += 1

    # print("Safe line found! Line is: ", line)
    if coments: print("Extited True", line)
    return 1

def main():
    # Read the data from the file and put the information into two different lists
    data = read_data_from_file('data.txt')
    #print(data)

    print("Advent of code day 1!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query? ")
    print("")

    match query:
        case "1":
            result = 0

            for line in data:
                if safety_test(line):
                    result += 1
            
            print("There are", result, "stable lines!")

        case "2":
            result = 0
            trash = []
            saved = []
            for index in range(len(data) - 1, -1, -1):
                line = data[index]
                match safety_test_with_dampner(line):
                    case 0:
                        trash.append(data.pop(index))  
                    case 1:
                        result += 1
                    case 2:
                        result += 1
                        #saved.append(data.pop(index))  
                    case _:
                        pass

            #print(trash)

            #for index in range(len(data) - 1, -1, -1):
            #    line = data[index]
            #    if safety_test_with_dampner(line):
            #        result += 1
            #    else:
            #        trash.append(data.pop(index))

            #print("Verifing mistakes in the aproved items:")
            #for line in data:
            #    if not check_ordered_difference_with_tolerance(line):
            #        print(line)
            #print("Verifing mistakes in the deleted items:")
            #for line in trash:
            #    if check_ordered_difference_with_tolerance(line):
            #        print(line)
            print("There are", result, "stable lines!")
            
        case _:
            print("Invalid option!")


    print("That's all folks!")
    return 0


if __name__ == '__main__':
    main()
