#!/usr/bin/python3


def diagonalDifference(arr):
    # Write your code here
    assert len(arr) == len(arr[0])
    last_index = len(arr[0]) - 1
    l_r_diagonal = 0
    r_l_diagonal = 0
    for i in range(len(arr)):
        l_r_diagonal += arr[i][i]
        r_l_diagonal += arr[last_index - i][i]
    print(l_r_diagonal)
    print(r_l_diagonal)
    return abs(l_r_diagonal - r_l_diagonal)


"""
1   2  4
6   8 10
12 14 16

=> |25-24| = 1
"""
diagonalDifference([[1, 2, 4], [6, 8, 10], [12, 14, 16]])
