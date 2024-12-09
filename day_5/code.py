
    
def parse_data_from_file(filename):
    """
    Reads a file and separates content into two parts:
    - Lines before the first newline ('\\n').
    - Lines after the first newline.
    """
    rules = []
    pages = []
    found_newline = False

    with open(filename, 'r') as file:
        for line in file:
            if not found_newline:
                if line == '\n':  # Found the line with only a newline
                    found_newline = True
                else:
                    rules.append([int(token) for token in line.split("|")])  # Process lines before newline
            else:
                pages.append([int(token) for token in line.split(",")])  # Process lines after newline

    return rules, pages

def test_rule(rule: list, pages:list):
    try:
        index1 = pages.index(rule[0])  # Get the index of num1
        index2 = pages.index(rule[1])  # Get the index of num2
        return index1 < index2    # Check if num1 appears before num2
    except ValueError:
        return True  # If one of the numbers is not in the list, the rule does not matter

def test_all_rules(rules, pages):
    for rule in rules:
        if not test_rule(rule, pages):
            return False
    return True

def correct_rule(rule:list, pages:list):
    try:
        #print("List is ",pages)
        index1 = pages.index(rule[0])  # Get the index of num1
        index2 = pages.index(rule[1])  # Get the index of num2
        #print("Rule to be aplied is", rule[0], "|", rule[1])
        #print("Indexes are ", index1, " ", index2)
        #print("condition is", index1 > index2)

        if index1 > index2:    # Check if num1 appears before num2
            pages.insert(index1,pages.pop(index2))
            #print("changes the pages ", pages)
        return pages
    except ValueError:
        return pages  # If one of the numbers is not in the list, the rule does not matter


def correct_updates(rules, pages):
    for update in pages:
        while test_all_rules(rules,update) is False:
            for rule in rules:
                update = correct_rule(rule, update)


def main():
    # Example usage
    file_path = 'test_data.txt' 
    file_path = 'data.txt' 

    print("Advent of code day 1!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = input("What is the query?\n")
    print("")


    rules, pages = parse_data_from_file(file_path)

    match query:
        case "1":
            result = 0
            for update in pages:
                if test_all_rules(rules, update):
                    result += update[len(update)//2]
            print(result)
        case "2":
            errors = []
            for update in pages:
                if test_all_rules(rules, update) is False:
                    errors.append(update)

            correct_updates(rules, errors)
            result = sum(update[len(update)//2] for update in errors)
            print(result)

        case _:
            print("Invalid option!")


    print("That's all folks!")
    return 0


if __name__ == '__main__':
    main()
