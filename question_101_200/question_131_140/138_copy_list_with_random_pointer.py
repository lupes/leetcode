# 138. 复制带随机指针的链表
# https://leetcode-cn.com/problems/copy-list-with-random-pointer
# Topics: 哈希表 链表


class Node:
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random

    def __str__(self):
        return "{val:%s next:%s random:%s}" % (self.val, self.next, id(self.random))


class Solution:
    def copyRandomList(self, head: Node) -> Node:
        new = Node(0, None, None)
        head = Node(0, head, None)
        nn = new
        hn = head.next
        m = {}
        while hn is not None:
            node = Node(hn.val, None, None)
            m[hn] = node
            nn.next = node
            nn = nn.next
            hn = hn.next

        hn = head
        nn = new
        while hn is not None:
            if hn.random is not None:
                nn.random = m[hn.random]
            hn = hn.next
            nn = nn.next
        return new.next


if __name__ == '__main__':
    solution = Solution()

    node2 = Node(2, None, None)
    node2.random = node2
    node1 = Node(1, node2, node2)

    print(node1)
    print(solution.copyRandomList(node1))
