# 590. N叉树的后序遍历
// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
// Topics: 树

"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        
