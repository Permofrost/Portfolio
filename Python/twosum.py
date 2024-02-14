#nums -> array of integers
#target -> target sum

#class Solution:
#    def twoSum(self, nums: List[int], target: int) -> List[int]:

#        """First Iteration"""
#        for i in range(len(nums)):
#            """Second Iteration, one step ahead"""
#            for j in range(i+1, len(nums)):
#                if (
#                    (i!=j) and ((nums[i]+nums[j]) == target)
#                    ):
#                    return [i,j]
#        return nums

    #Brute Force Approach:
    #Time Complexity: O(n^2)
    #Space Complexity: O(1)

#Inputs: nums (List of Integers) + target (Sum of targets)
class Solution:
    #def __init__(self, nums: list[int], target: int):
        #Initializing the values within this class for running in IDE
        #self.nums = nums
        #self.target = target
        #self.num_dict = {}
        #print(self.twoSum())
    
    
    def twoSum(self,nums,target) -> list[int]:
        #Create an empty dictionary
        num_dict = {}
        for i, num in enumerate(nums):
            #Calculate the complement of the current number in the iteration
            complement = target - num
            #If the complement is in the dictionary, return the index of the complement, and the current index
            if complement in num_dict:
                return [num_dict[complement], i]
            #Otherwise add the number to the dictionary
            num_dict[num] = i
        #Base case: Empty List
        return []
    
    #Dictionary Approach
    #Time Complexity: O(n)
    #Space Complexity: O(n)

    