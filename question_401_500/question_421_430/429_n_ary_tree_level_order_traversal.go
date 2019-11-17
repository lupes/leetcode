# 429. N叉树的层序遍历
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
// Topics: 树 广度优先搜索

"""
# Definition for a Node.
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        
