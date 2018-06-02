// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/572/

class BuyAndSellOnce {
    public int maxProfit(int[] prices) {
        int currentMinPrice = Integer.MAX_VALUE;
        int maxProfit = 0;

        for (int i = 0; i < prices.length; i++) {
            if (prices[i] < currentMinPrice) {
                currentMinPrice = prices[i];
            }
            if (prices[i] - currentMinPrice > maxProfit) {
                maxProfit = prices[i] - currentMinPrice;
            }
        }

        return maxProfit;
    }

    public static void main(String[] args) {
        BuyAndSellOnce bs = new BuyAndSellOnce();

        int[][] tests = {
            {1, 2, 3, 4, 5, 6},
            {7, 1, 5, 3, 6, 4},
            {7, 6, 4, 3, 1},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(bs.maxProfit(test));
            System.out.println();
        }
    }
}
