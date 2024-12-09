import itertools
    
def lines(filename):
    """Process the file line by line, yielding first value and values after ':'."""
    with open(filename, "r") as data:
        for line in data:
            # Split the line by ':'
            parts = line.split(":")
            
            # Extract the first value (before the colon)
            first_value = int(parts[0].strip())
            
            # Extract the list of values after the colon, split by spaces
            values_after_colon = list(map(int, parts[1].strip().split()))
            
            # Yield the result
            yield first_value, values_after_colon

def calculate(values:list):
    if len(values) > 2:
        my_value = values.pop()
        results = calculate(values)
        result = []
        
        for r in results:
            result.append((r + my_value))
            result.append((r * my_value))
    else:
        result = [(values[0] + values[1]), (values[0] * values[1])] 

    return result

def calculate2(values:list):
    if len(values) > 2:
        my_value = values.pop()
        results = calculate2(values)
        result = []
        
        for r in results:
            result.append((r + my_value))
            result.append((r * my_value))
            result.append(int(str(r) + str(my_value)))
    else:
        result = [(values[0] + values[1]), (values[0] * values[1]), (int(str(values[0]) + str(values[1])))] 

    return result


def main():
    # Example usage
    file_path = '2024/day_7/test_data.txt'
    file_path = '2024/day_7/data.txt' 

    print("Advent of code day 7!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")

    result = 0

    match query:
        case "1":
            for value, terms in lines(file_path):
                if value in calculate(terms):
                    result += value
                    print(value, terms)

            
            print("Total calibration are", result)
        case "2":
            for value, terms in lines(file_path):
                if value in calculate2(terms):
                    result += value
                    print(value, terms)

            print("Total calibration are", result)
        case _:
            print("Invalid option!")

    print("That's all folks!")
    return 0

if __name__ == '__main__':
    main()
