# 116. 填充每个节点的下一个右侧节点指针
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
// Topics: 树 深度优先搜索

"""
# Definition for a Node.
class Node:
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
class Solution:
    def connect(self, root: 'Node') -> 'Node':
        
