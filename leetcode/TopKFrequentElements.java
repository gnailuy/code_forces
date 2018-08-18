// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/799/

import java.lang.Comparable;
import java.util.Comparator;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;

class TopKFrequentElements {
    class MapValueDescComparator<K, V extends Comparable<V>> implements Comparator<Map.Entry<K, V>> {
        public int compare(Map.Entry<K, V> a, Map.Entry<K, V> b) {
            return b.getValue().compareTo(a.getValue());
        }
    }

    List<Integer> topKFrequent(int[] nums, int k) {
        List<Integer> result = new LinkedList<Integer>();
        if (k <= 0) return result;

        HashMap<Integer, Integer> numCount = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            if (numCount.containsKey(nums[i])) {
                numCount.put(nums[i], numCount.get(nums[i]) + 1);
            } else {
                numCount.put(nums[i], 1);
            }
        }

        PriorityQueue<Map.Entry<Integer, Integer>> pq =
            new PriorityQueue<Map.Entry<Integer, Integer>>(nums.length, new MapValueDescComparator<Integer, Integer>());
        for (Map.Entry<Integer, Integer> entry: numCount.entrySet()) {
            pq.add(entry);
        }

        for (int i = 0; i < k; i++) {
            Map.Entry<Integer, Integer> entry = pq.poll();
            result.add(entry.getKey());
        }

        return result;
    }

    public static void main(String[] args) {
        TopKFrequentElements tkfe = new TopKFrequentElements();

        int[][][] tests = {
            {{1, 1, 1, 2, 2, 3}, {2}},
        };
        for (int[][] test : tests) {
            int[] nums = test[0];
            int k = test[1][0];

            System.out.println(nums.length + " " + k);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();

            List<Integer> result = tkfe.topKFrequent(nums, k);

            for (Integer i: result) {
                System.out.print(i + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
