
def main():
    # Example usage
    file_path = '2024/day_/test_data.txt'
    #file_path = '2024/day_/data.txt' 

    print("Advent of code day !")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")



    match query:
        case "1":
            pass
        case "2":
            pass
        case _:
            print("Invalid option!")

    print("That's all folks!")
    return 0

if __name__ == '__main__':
    main()
