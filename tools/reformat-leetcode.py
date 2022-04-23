import os


def create_packages(path: str, solution_ext: str):
    solutions = [s for s in os.listdir(path) if s.endswith(solution_ext)]
    packages = [s.split(solution_ext)[0] for s in solutions]
    for i in range(len(packages)):
        package_name = packages[i]
        if not os.path.exists(os.path.join(path, packages[i])):
            os.mkdir(os.path.join(path, packages[i]))
        with open(os.path.join(path, solutions[i]), "r") as fin:
            with open(os.path.join(path, packages[i], solutions[i]), "w") as fout:
                for l in fin:
                    fout.write(l)


def main():
    create_packages("../leetcode", ".go")


if __name__ == '__main__':
    main()
