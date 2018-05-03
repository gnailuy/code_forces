// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/

import java.util.ArrayList;
import java.util.HashMap;

class IntersectionArrays {
    public int[] intersection(int[] nums1, int[] nums2) {
        ArrayList<Integer> al = new ArrayList<Integer>();
        HashMap<Integer, Integer> store = new HashMap<Integer, Integer>();

        int[] s, l;
        if (nums1.length <= nums2.length) {
            s = nums1;
            l = nums2;
        } else {
            s = nums2;
            l = nums1;
        }

        for (int i = 0; i < s.length; i++) {
            if (store.containsKey(s[i])) {
                store.put(s[i], store.get(s[i]) + 1);
            } else {
                store.put(s[i], 1);
            }
        }
        for (int i = 0; i < l.length; i++) {
            if (store.containsKey(l[i]) && store.get(l[i]) > 0) {
                store.put(l[i], store.get(l[i]) - 1);
                al.add(l[i]);
            }
        }

        int[] result = new int[al.size()];
        for (int i = 0; i < al.size(); i++) {
            result[i] = al.get(i);
        }
        return result;
    }

    public static void main(String[] args) {
        IntersectionArrays ia = new IntersectionArrays();

        int[][][] tests = {
            {{1, 2, 2, 1}, {2, 2}},
        };
        for (int[][] test : tests) {
            System.out.println(test[0].length + " " + test[1].length);
            for (int i = 0; i < test[0].length; i++) {
                System.out.print(test[0][i] + " ");
            }
            System.out.println();
            for (int i = 0; i < test[1].length; i++) {
                System.out.print(test[1][i] + " ");
            }
            System.out.println();
            int[] result = ia.intersection(test[0], test[1]);
            for (int i = 0; i < result.length; i++) {
                System.out.print(result[i] + " ");
            }
            System.out.println();
        }
    }
}
