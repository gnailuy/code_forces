// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/827/

class ProductOfArrayExceptSelf {
    public int[] productExceptSelf(int[] nums) {
        int n = nums.length;

        int[] result = new int[n];
        int after = 1;

        result[0] = 1;
        for (int i = 1; i < n; i++) {
            result[i] = result[i-1]*nums[i-1];
        }
        for (int i = n-2; i >= 0; i--) {
            after *= nums[i+1];
            result[i] *= after;
        }

        return result;
    }

    public static void main(String[] args) {
        ProductOfArrayExceptSelf poa = new ProductOfArrayExceptSelf();

        int[][] tests = {
            {1, 2, 3, 4},
            {1, 2, 3, 4, 5, 6, 7},
            {1, 0},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            int[] result = poa.productExceptSelf(test);
            for (int i = 0; i < result.length; i++) {
                System.out.print(result[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
