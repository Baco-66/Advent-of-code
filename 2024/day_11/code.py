import cProfile
from collections import Counter


def parse_data_from_file(filename):
    data = []
    with open(filename, 'r') as file:
        for line in file:
            data = ([int(token) for token in line.split(" ")])
    return data

'''
If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.

If the stone is engraved with a number that has an even number of digits, it is replaced by two 
stones. The left half of the digits are engraved on the new left stone, and the right half of the 
digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 
1000 would become stones 10 and 0.)

If none of the other rules apply, the stone is replaced by a new stone; the old stone's number 
multiplied by 2024 is engraved on the new stone.
'''

def blink(stones:list, size):
    for index in range(size-1, -1, -1):
        if stones[index] == 0:
            stones[index] = 1
        elif (mid := len((stone := str(stones[index])))) % 2 == 0:
            mid = mid//2
            stones[index] = int(stone[mid:])
            #stones.insert(index,int(stone[:mid]))
            stones.append(int(stone[:mid]))
            size += 1
        else:
            stones[index] = stones[index] * 2024
    return stones, size

'''
def blink_recursive(stone, iterations):
    size = 1
    if (mid := len((stone := str(stone)))) % 2 == 0:
        mid = mid//2
        stones[index] = int(stone[mid:])
        #stones.insert(index,int(stone[:mid]))
        stones.append(int(stone[:mid]))
        size += 1
    else:
        if stone == 0:
            blink_recursive(1, iterations-1)
        else:
            blink_recursive(stone*2024, iterations-1)

    if stone == 0:
        if iterations == 0:
            return 1
        elif iterations == 1:
            return 1
        elif iterations == 2:
            return 2
        blink_recursive(1)
    elif (mid := len((stone := str(stones[index])))) % 2 == 0:
        mid = mid//2
        stones[index] = int(stone[mid:])
        #stones.insert(index,int(stone[:mid]))
        stones.append(int(stone[:mid]))
        size += 1

    return size
'''
#@profile
def blink_reduced_redunduncy(stones:list):
    for index in range(len(stones)-1, -1, -1):
        if stones[index][0] == 0:
            stones[index][0] = 1
        elif (mid := len((stone := str(stones[index][0])))) % 2 == 0:
            mid = mid//2
            stones[index][0] = int(stone[mid:])
            #stones.insert(index,int(stone[:mid]))
            stones.append([int(stone[:mid]),stones[index][1]])
        else:
            stones[index][0] = stones[index][0] * 2024
    return stones

def reduce_numbers(data):
    """Reduce redundancy in a list of numbers."""
    count = Counter(data)
    return [[num, freq] for num, freq in count.items()]  # Return list of lists

def merge_counts(data):
    """Merge counts from a list of lists (number, frequency)."""
    count = Counter()
    for number, freq in data:
        count[number] += freq
    return [[num, freq] for num, freq in count.items()]  # Return list of lists

def main():
    # Example usage
    file_path = '2024/day_11/test_data.txt'
    file_path = '2024/day_11/data.txt' 

    print("Advent of code day 11!")
    print("For part one type number 1")
    print("For part two type number 2")

    # Ask the user for the query to do
    query = "2"#input("What is the query?\n")

    stones = parse_data_from_file(file_path)

    match query:
        case "1":
            size = len(stones)
            for _ in range(25):
                stones, size = blink(stones, size)
            print("Result is", size)
        case "2":
            stones = reduce_numbers(stones)
            for _ in range(75): # 75
                stones = blink_reduced_redunduncy(stones)
                stones = merge_counts(stones)
            print(stones)
            result = 0
            for stone in stones:
                #print(stone[1], end=" ")
                result += stone[1]
            print("Result is", result)
        case _:
            print("Invalid option!")

    print("That's all folks!")
    return 0

if __name__ == '__main__':
    main()
    #cProfile.run('main()')
