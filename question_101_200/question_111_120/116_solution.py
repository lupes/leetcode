# 116. 填充每个节点的下一个右侧节点指针
# https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/


class Node:
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

    def __repr__(self):
        val = self.val
        left = self.left
        right = self.right
        next = self.next
        return f"[val:{val} left:{left} right:{right} next:{next}]"


class Solution(object):
    def connect(self, root):
        """
        :type root: Node
        :rtype: Node
        """
        if root is None:
            return
        nodes = [root]
        while len(nodes) > 0:
            next = []
            for i, node in enumerate(nodes):
                if len(nodes) > i + 1:
                    node.next = nodes[i + 1]
                if node.left is not None:
                    next.extend([node.left, node.right])
            nodes = next
        return root


def test():
    node4 = Node(4, None, None, None)
    node5 = Node(5, None, None, None)
    node6 = Node(6, None, None, None)
    node7 = Node(7, None, None, None)

    node2 = Node(2, node4, node5, None)
    node3 = Node(3, node6, node7, None)

    node1 = Node(1, node2, node3, None)
    solution = Solution()
    print(solution.connect(node1))


if __name__ == '__main__':
    test()
