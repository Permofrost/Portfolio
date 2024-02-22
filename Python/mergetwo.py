# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
         self.val = val
         self.next = next
    
    def append(self, val):
        current = self
        while current.next is not None:
            current = current.next
        current.next = ListNode(val)

    def print_linked_list(node):
        while node is not None:
            print(node.val, end=" ")
            node = node.next
        print()

class Solution:
    def merging(self, a: ListNode, b: ListNode) -> ListNode:
        head = ListNode(0)
        handler = head
        
        while a is not None and b is not None:
            if a.val < b.val:
                handler.next = a
                a = a.next
            else:
                handler.next = b
                b = b.next
            handler = handler.next

        if a is not None:
            handler.next = a
        elif b is not None:
            handler.next = b

        return head.next
    






solution = Solution()

input_one = [1,4,7]
a = ListNode(input_one[0])  
for i in input_one[1:]:
    a.append(i)

input_two = [2,3,5]
b = ListNode(input_two[0])
for j in input_two[1:]:
    b.append(j)

result = solution.merging(a, b)


ListNode.print_linked_list(result)