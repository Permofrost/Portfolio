public class bsstock {
    static class Solution {
        public int maxProfit(int[] prices) {
            
            //max possible and minimum possible price
            int minPrice = 10000;
            int maxProfit = 0;

            //Array iterates forwards, updating min prices
            for (int i = 0; i < prices.length; i++) {
                if (prices[i] < minPrice) {
                    minPrice = prices[i];
                } else 
                //Max profit is only updated if in runs where min price **isn't** updated.
                //This guarantees a max profit only being updated using prices after the min price.
                if (prices[i] - minPrice > maxProfit) {
                    maxProfit = prices[i] - minPrice;
                }
            }
            return Math.max(0,maxProfit);
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] prices = {2, 1, 2, 1, 0, 0, 1};
        int result = s.maxProfit(prices);
        System.out.println(result);  
    }
}

