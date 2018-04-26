// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/

class BestTimeBuySell {
    public int maxProfit(int[] prices) {
        int profit = 0;
        int stock_price = -1;
        for (int i = 0; i < prices.length-1; i++) {
            if (stock_price < 0 && prices[i] < prices[i+1]) {
                stock_price = prices[i]; // Buy
            } else if (stock_price >= 0 && prices[i] > prices[i+1]) {
                profit += prices[i] - stock_price; // Sell
                stock_price = -1;
            }
        }

        if (stock_price >= 0) {
            profit += prices[prices.length-1] - stock_price;
        }

        return profit;
    }

    public static void main(String[] args) {
        BestTimeBuySell bt = new BestTimeBuySell();

        int[][] tests = {
            {},
            {1},
            {7, 1, 5, 3, 6, 4},
            {1, 2, 3, 4, 5},
            {7, 6, 4, 3, 1},
            {2, 1, 2, 0, 1}
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(bt.maxProfit(test));
            System.out.println();
        }
    }
}
