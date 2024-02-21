import java.util.HashMap;
import java.util.Map;

// Solution class
public class twosum {
    // twosum method
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    public int[] twosum(int[] nums, int target) {
        // Create an empty map
        Map<Integer, Integer> numDict = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            // Calculate the complement of the current number in the iteration
            int complement = target - nums[i];
            // If the complement is in the map, return the index of the complement, and the current index
            if (numDict.containsKey(complement)) {
                return new int[]{numDict.get(complement), i};
            }
            // Otherwise add the number to the map
            numDict.put(nums[i], i);
        }
        // Base case: Empty array
        return new int[]{};
    }

    public static void main(String[] args) {
        twosum solution = new twosum();

        // Inputs: nums (array of integers) + target (sum of targets)
        int[] nums = {2, 7, 11, 15};
        int target = 9;
        int[] result = solution.twosum(nums, target);

        // Print the result
        for (int num : result) {
            System.out.print(num + " ");
        }
    }
}