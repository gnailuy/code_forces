// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/778/

import java.util.Arrays;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

class GroupAnagrams {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> result = new LinkedList<List<String>>();

        HashMap<String, List<String>> resultMap = new HashMap<String, List<String>>();
        for (String s : strs) {
            char[] sa = s.toCharArray();
            Arrays.sort(sa);
            String key = new String(sa);
            if (!resultMap.containsKey(key)) {
                resultMap.put(key, new LinkedList<String>());
            }
            resultMap.get(key).add(s);
        }
        for (Map.Entry<String, List<String>> e: resultMap.entrySet()) {
            result.add(e.getValue());
        }

        return result;
    }

    public static void main(String[] args) {
        String[][] tests = {
            {"eat", "tea", "tan", "ate", "nat", "bat"},
        };
        for (String[] test : tests) {
            GroupAnagrams ga = new GroupAnagrams();
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();

            List<List<String>> groups = ga.groupAnagrams(test);
            for (List<String> group : groups) {
                for (String s : group) {
                    System.out.print(s + " ");
                }
                System.out.println();
            }
        }
    }
}
