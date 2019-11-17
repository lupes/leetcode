# 430. 扁平化多级双向链表
// https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list
// Topics: 深度优先搜索 链表

"""
# Definition for a Node.
class Node:
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child
"""
class Solution:
    def flatten(self, head: 'Node') -> 'Node':
        
