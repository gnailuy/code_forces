// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/601/

import java.util.List;
import java.util.LinkedList;

class PascalTriangle {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> result = new LinkedList<List<Integer>>();
        if (1 <= numRows) {
            result.add(new LinkedList<Integer>() {{
                add(1);
            }});
        }
        if (2 <= numRows) {
            result.add(new LinkedList<Integer>() {{
                add(1);
                add(1);
            }});
        }
        for (int i = 3; i <= numRows; i++) {
            List<Integer> line = new LinkedList<Integer>();
            line.add(1);
            for (int j = 0; j < i-2; j++) {
                line.add(result.get(i-2).get(j) + result.get(i-2).get(j+1));
            }
            line.add(1);
            result.add(line);
        }
        return result;
    }

    public static void main(String[] args) {
        PascalTriangle pt = new PascalTriangle();

        int[] tests = {
            0,
            1,
            2,
            3,
            4,
            5,
            10,
        };

        for (int test: tests) {
            System.out.println(test);
            List<List<Integer>> triangle = pt.generate(test);
            for (List<Integer> line: triangle) {
                for (Integer n: line) {
                    System.out.print(n + " ");
                }
                System.out.println();
            }
            System.out.println();
        }
    }
}

