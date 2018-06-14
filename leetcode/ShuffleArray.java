// https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/670/

import java.util.Random;

class ShuffleArray {
    Random rand;
    private int[] store;

    public ShuffleArray(int[] nums) {
        rand = new Random();
        store = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            store[i] = nums[i];
        }
    }

    public int[] reset() {
        return store;
    }

    public int[] shuffle() {
        int[] shuffle = new int[store.length];
        for (int i = 0; i < store.length; i++) {
            shuffle[i] = store[i];
        }
        for (int i = 0; i < shuffle.length; i++) {
            int random = rand.nextInt(shuffle.length-i) + i;
            if (random != i) {
                int tmp = shuffle[random];
                shuffle[random] = shuffle[i];
                shuffle[i] = tmp;
            }
        }
        return shuffle;
    }

    public static void main(String[] args) {
        int[][] tests = {
            {1},
            {1, 2},
            {1, 2, 3},
        };
        for (int[] test : tests) {
            ShuffleArray sa = new ShuffleArray(test);
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            for (int i = 0; i < 10; i++) {
                int[] result = sa.shuffle();
                for (int j = 0; j < result.length; j++) {
                    System.out.print(result[j] + " ");
                }
                System.out.println();
            }
            int[] result = sa.reset();
            for (int j = 0; j < result.length; j++) {
                System.out.print(result[j] + " ");
            }
            System.out.println();
        }
    }
}
