
def read_data_from_file(filename):
    """Reads data from a file and returns two lists: left and right values."""
    left = []
    right = []
    with open(filename, 'r') as file:
        for line in file:
            left_value, right_value = line.strip().split('   ')
            left.append(left_value)
            right.append(right_value)
    return left, right


def main():
    # Read the data from the file and put the information into two different lists
    left, right = read_data_from_file('2024/day_1/data.txt')

    # print("List with right values: ", right, end="\n")

    # print("List with left values: ", left, end="\n")

    print("Advent of code day 1!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query? ")
    print("")

    match query:
        case "1":
            result = 0

            # Get the smalest number form each list and get the absolute value of one minus the other
            while right:
                right_smallest = min(right)
                left_smallest = min(left)

                right.remove(right_smallest)  
                left.remove(left_smallest)  

                result += abs(int(right_smallest) - int(left_smallest))

            print("Result is ",result, "!")
        case "2":
            result = 0

            # Get the smalest number form each list and get the absolute value of one minus the other
            #for left_elem in left:
            #    value = 0
            #    for right_elem in right:
            #        if left_elem == right_elem:
            #            value += 1
            #    result += value

            for elem in left:
                result += int(elem) * right.count(elem)

            print("Result is ",result, "!")
        case _:
            print("Invalid option!")


    print("That's all folks!")
    return 0


if __name__ == '__main__':
    main()
