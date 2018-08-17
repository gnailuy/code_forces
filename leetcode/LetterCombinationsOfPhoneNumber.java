// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/793/

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;

class LetterCombinationsOfPhoneNumber {
    HashMap<Character, String> keyMap = new HashMap<Character, String>() {{
        put('2', "abc");
        put('3', "def");
        put('4', "ghi");
        put('5', "jkl");
        put('6', "mno");
        put('7', "pqrs");
        put('8', "tuv");
        put('9', "wxyz");
    }};

    public List<String> letterCombinations(String digits) {
        List<String> next;
        List<String> result = new LinkedList<String>();
        if (digits.length() <= 0) return result;

        result.add("");
        for (int i = 0; i < digits.length(); i++) {
            String letters = keyMap.get(digits.charAt(i));
            next = new LinkedList<String>();
            for (String s: result) {
                for (char c: letters.toCharArray()) {
                    next.add(s + c);
                }
            }
            result = next;
        }

        return result;
    }

    public static void main(String[] args) {
        LetterCombinationsOfPhoneNumber lc = new LetterCombinationsOfPhoneNumber();

        String[] tests = {
            "23",
            "234",
        };
        for (String test : tests) {
            System.out.println(test);
            List<String> result = lc.letterCombinations(test);
            for (String s: result) {
                System.out.print(s + " ");
            }
            System.out.println();
        }
    }
}
