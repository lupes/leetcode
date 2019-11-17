# 559. N叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
// Topics: 树 深度优先搜索 广度优先搜索

"""
# Definition for a Node.
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution:
    def maxDepth(self, root: 'Node') -> int:
        
