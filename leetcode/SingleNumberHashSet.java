// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
import java.util.HashSet;

class SingleNumberHashSet {
    public int singleNumber(int[] nums) {
        HashSet<Integer> cache = new HashSet<Integer>();
        for (int i = 0; i < nums.length; i++) {
            if (cache.contains(nums[i])) {
                cache.remove(nums[i]);
            } else {
                cache.add(nums[i]);
            }
        }
        assert cache.size() == 1;
        for (Integer i: cache) {
            return i;
        }
        return -1; // This will never happen
    }

    public static void main(String[] args) {
        SingleNumberHashSet sn = new SingleNumberHashSet();

        int[][] tests = {
            {2, 2, 1},
            {4, 1, 2, 1, 2},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(sn.singleNumber(test));
        }
    }
}
