import os

SOLUTIONS_DIR = "/Repos/solutions/adventofcode/2022"


class TreeNode:
    def __init__(self, name="", files_size=0, total_size=0, parent=None, ):
        self.name = name
        self.files_size = files_size
        self.total_size = total_size
        self.parent = parent
        self.children = dict()


def read_input(filename: str) -> list:
    with open(os.path.join(os.getcwd(), filename), "r") as f:
        return f.readlines()


def day_seven_1():
    text = read_input("day07_1.input")
    cmds = list()

    # getting list of commands
    i = 1
    while i < len(text):
        l = text[i].replace("\n", "")
        if "cd" in l:
            cmds.append(["cd", l.split(" ")[-1]])
        elif "ls" in l:
            output = list()
            j = i + 1
            while j < len(text) and "$" not in text[j]:
                output.append(text[j].replace("\n", ""))
                j += 1
            cmds.append(["ls", output])
            i = j - 1
        else:
            raise ValueError(l)
        i += 1

    # building filesystem tree
    fstree = TreeNode(name="/")
    cn = fstree
    for cmd in cmds:
        if cmd[0] == "cd":
            target = cmd[1]
            if target == "..":
                cn = cn.parent
            else:
                if target not in cn.children:
                    cn.children[target] = TreeNode(name=target, parent=cn)
                cn = cn.children[target]
        else:
            output = cmd[1]
            files_size = 0
            for o in output:
                if "dir" in o:
                    dir_name = o.replace("dir ", "")
                    cn.children[dir_name] = TreeNode(name=dir_name, parent=cn)
                else:
                    files_size += int(o.split(" ")[0])
            cn.files_size = files_size

    # iterating over the tree to calculate the total size of every directory
    def calculate_total_size(node: TreeNode) -> int:
        node.total_size = node.files_size
        for ch in node.children.values():
            node.total_size += calculate_total_size(ch)
        return node.total_size

    calculate_total_size(fstree)

    # calculating result
    def calc_res(node: TreeNode) -> int:
        res = 0
        if node.total_size <= 100000:
            res = node.total_size
        for ch in node.children.values():
            res += calc_res(ch)
        return res

    print(calc_res(fstree))


def day_seven_2():
    text = read_input("day07_1.input")
    cmds = list()

    # getting list of commands
    i = 1
    while i < len(text):
        l = text[i].replace("\n", "")
        if "cd" in l:
            cmds.append(["cd", l.split(" ")[-1]])
        elif "ls" in l:
            output = list()
            j = i + 1
            while j < len(text) and "$" not in text[j]:
                output.append(text[j].replace("\n", ""))
                j += 1
            cmds.append(["ls", output])
            i = j - 1
        else:
            raise ValueError(l)
        i += 1

    # building filesystem tree
    fstree = TreeNode(name="/")
    cn = fstree
    for cmd in cmds:
        if cmd[0] == "cd":
            target = cmd[1]
            if target == "..":
                cn = cn.parent
            else:
                if target not in cn.children:
                    cn.children[target] = TreeNode(name=target, parent=cn)
                cn = cn.children[target]
        else:
            output = cmd[1]
            files_size = 0
            for o in output:
                if "dir" in o:
                    dir_name = o.replace("dir ", "")
                    cn.children[dir_name] = TreeNode(name=dir_name, parent=cn)
                else:
                    files_size += int(o.split(" ")[0])
            cn.files_size = files_size

    # iterating over the tree to calculate the total size of every directory
    sizes = list()

    def calculate_total_size(node: TreeNode) -> int:
        node.total_size = node.files_size
        for ch in node.children.values():
            node.total_size += calculate_total_size(ch)
        sizes.append(node.total_size)
        return node.total_size

    calculate_total_size(fstree)

    res = fstree.total_size
    for i in sizes:
        if 70_000_000 - fstree.total_size + i >= 30_000_000 and i < res:
            res = i

    print(res)


def main():
    # day_seven_1()
    day_seven_2()


if __name__ == '__main__':
    main()
