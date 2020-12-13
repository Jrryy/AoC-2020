position = [0, 0]
position_2 = [0, 0]
waypoint = [10, 1]
facing = 1
directions = ((0, 1), (1, 0), (0, -1), (-1, 0))

moves = {
    'N': 0,
    'E': 1,
    'S': 2,
    'W': 3,
}

rotations = {
    'R': 1,
    'L': -1,
}

waypoint_rotations = [
    lambda x: [x[0], x[1]],
    lambda x: [x[1], -x[0]],
    lambda x: [-x[0], -x[1]],
    lambda x: [-x[1], x[0]],
]


def move(p, d, a):
    return [p[0] + d[0] * a, p[1] + d[1] * a]


with open('input.txt') as f:
    for action in f.readlines():
        direction = action[0]
        amount = int(action[1:])
        if direction in moves:
            position = move(position, directions[moves[direction]], amount)
            waypoint = move(waypoint, directions[moves[direction]], amount)
        elif direction == 'F':
            position = move(position, directions[facing], amount)
            position_2 = move(position_2, waypoint, amount)
        else:
            facing = (facing + (amount//90)*rotations[direction]) % 4
            waypoint = waypoint_rotations[(amount//90)*rotations[direction] % 4](waypoint)
    print(abs(position[0]) + abs(position[1]))
    print(abs(position_2[0]) + abs(position_2[1]))
