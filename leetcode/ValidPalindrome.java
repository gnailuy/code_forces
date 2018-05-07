// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/883/

class ValidPalindrome {
    public char toLower(char c) {
        if (('a' <= c && c <= 'z') || ('0' <= c && c <= '9')) {
            return c;
        } else if ('A' <= c && c <= 'Z') {
            return (char)(c-'A'+'a');
        }
        return ' ';
    }

    public boolean isPalindrome(String s) {
        if (s.length() <= 1) return true;

        int i = 0, j = s.length()-1;
        while (i < j) {
            char c1 = toLower(s.charAt(i));
            char c2 = toLower(s.charAt(j));
            while (i < j) {
                if (' ' != c1) {
                    break;
                }
                c1 = toLower(s.charAt(++i));
            }
            while (i < j) {
                if (' ' != c2) {
                    break;
                }
                c2 = toLower(s.charAt(--j));
            }
            if (c1 != c2) {
                return false;
            } else {
                i++;
                j--;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        ValidPalindrome vp = new ValidPalindrome();

        String[] tests = {
            "A man, a plan, a canal: Panama",
            "race a car",
            ".,",
            "0P",
        };
        for (String test : tests) {
            System.out.println(test.length());
            System.out.println(test);
            System.out.println(vp.isPalindrome(test));
        }
    }
}
