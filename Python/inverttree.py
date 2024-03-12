class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        if root:
           invert = self.invertTree
           root.left, root.right = invert(root.right), invert(root.left) 
           return root




# Create an instance of the Solution class
solution = Solution()

