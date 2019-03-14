def init_knapsack():
    pass


def find_solution_knapsack(w, v, W):
    solutions = [[None for _ in range(W + 1)] for _ in range(len(w) + 1)]

    for i in range(W + 1):
        solutions[0][i] = 0

    for i in range(1, len(w) + 1):
        for j in range(W + 1):
            if w[i - 1] <= j:
                solutions[i][j] = max(
                    solutions[i - 1][j],
                    v[i - 1] + solutions[i - 1][j - w[i - 1]])
            else:
                solutions[i][j] = solutions[i - 1][j]
    for i in solutions:
        print(i)


if __name__ == "__main__":
    find_solution_knapsack([5, 4, 6, 3], [10, 40, 30, 50], 10)
