// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/743/

import java.util.LinkedList;
import java.util.List;

class FizzBuzz {
    public List<String> fizzBuzz(int n) {
        List<String> result = new LinkedList<String>();
        for (int i = 1; i <= n; i++) {
            if (0 == i%15) {
                result.add("FizzBuzz");
            } else if (0 == i%5) {
                result.add("Buzz");
            } else if (0 == i%3) {
                result.add("Fizz");
            } else {
                result.add(String.valueOf(i));
            }
        }
        return result;
    }

    public static void main(String[] args) {
        FizzBuzz fb = new FizzBuzz();

        int[] tests = {
            1,
            3,
            15,
            100,
            1000,
        };

        for (int test: tests) {
            List<String> r = fb.fizzBuzz(test);
            System.out.println(test);
            for (String s: r) {
                System.out.print(s + " ");
            }
            System.out.println("\n");
        }
    }
}

