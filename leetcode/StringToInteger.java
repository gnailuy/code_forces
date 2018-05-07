// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/884/

class StringToInteger {
    public int myAtoi(String str) {
        boolean start = false;
        boolean signed = false;
        boolean met_non_zero = false;
        boolean negative = false;
        int result = 0;
        int digit;

        for (int i = 0; i < str.length(); i++) {
            char c = str.charAt(i);
            if (' ' == c) {
                if (start) {
                    break;
                } else {
                    continue;
                }
            } else if ('-' != c && '+' != c && (c < '0' || c > '9')) {
                if (start) {
                    break;
                } else {
                    return 0;
                }
            } else if ('-' == c) {
                if (start || signed) {
                    break;
                }
                start = true;
                negative = true;
                signed = true;
            } else if ('+' == c) {
                if (start || signed) {
                    break;
                }
                start = true;
                signed = true;
            } else {
                start = true;
                digit = c-'0';
                if (negative && met_non_zero) {
                    if ((Integer.MIN_VALUE + digit) / 10 > result) {
                        return Integer.MIN_VALUE;
                    }
                    result = result * 10 - digit;
                } else {
                    if ((Integer.MAX_VALUE - digit) / 10 < result) {
                        return Integer.MAX_VALUE;
                    }
                    result = result * 10 + digit;
                }
                if (negative && !met_non_zero && digit > 0) {
                    met_non_zero = true;
                    result *= -1;
                }
            }
        }
        return result;
    }

    public static void main(String[] args) {
        StringToInteger sti = new StringToInteger();

        String[] tests = {
            "42",
            "    -42",
            "4193 with words",
            "words and 987",
            "-91283472332",
            "91283472332",
            "+1",
            "+-2",
            "  +0 123",
            "0-1",
            "-5-",
        };
        for (String test : tests) {
            System.out.println(test);
            System.out.println(sti.myAtoi(test));
            System.out.println();
        }
    }
}
