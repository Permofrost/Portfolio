class Solution:
    def maxProfit(self, prices):
        # max possible and minimum possible price
        min_price = 10000
        max_profit = 0

        # Array iterates forwards, updating min prices
        for i in range(len(prices)):
            if prices[i] < min_price:
                min_price = prices[i]
            else:
                # Max profit is only updated if in runs where min price isn't updated.
                 # This guarantees a max profit only being updated using prices after the min price.
                if prices[i] - min_price > max_profit:
                    max_profit = prices[i] - min_price

        return max(0, max_profit)


if __name__ == "__main__":
    s = Solution()
    prices = [2, 1, 2, 1, 0, 0, 1]
    result = s.maxProfit(prices)
    print(result)