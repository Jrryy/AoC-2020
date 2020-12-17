puzzle_input = (14, 8, 16, 0, 1, 17)

said = {number: turn + 1 for turn, number in enumerate(puzzle_input)}

i = len(puzzle_input) + 1
last_spoken = 0

# while i < 2020:
while i < 30000000:
    if last_spoken not in said:
        said[last_spoken] = i
        last_spoken = 0
    else:
        last_spoken_turn = said[last_spoken]
        said[last_spoken] = i
        last_spoken = i - last_spoken_turn
    i += 1

print(last_spoken)
