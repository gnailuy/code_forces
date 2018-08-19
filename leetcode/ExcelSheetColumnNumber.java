// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/817/

class ExcelSheetColumnNumber {
    public int titleToNumber(String s) {
        int result = 0;

        for (int i = 0; i < s.length(); i++) {
            int v = s.charAt(i) - 'A' + 1;
            result *= 26;
            result += v;
        }

        return result;
    }

    public static void main(String[] args) {
        ExcelSheetColumnNumber escn = new ExcelSheetColumnNumber();

        String[] tests = {
            "A",
            "B",
            "AB",
            "ZY",
        };
        for (String test : tests) {
            System.out.println(test);
            System.out.println(escn.titleToNumber(test));
            System.out.println();
        }
    }
}
