// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/885/

class StrStr {
    public int strStr(String haystack, String needle) {
        if (needle.length() <= 0) {
            return 0;
        } else if (haystack.length() < needle.length()) {
            return -1;
        }

        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            boolean ok = true;
            for (int j = 0; j < needle.length(); j++) {
                if (haystack.charAt(i+j) != needle.charAt(j)) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return i;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        StrStr ss = new StrStr();

        String[][] tests = {
            {"hello", "ll"},
            {"hello", ""},
            {"aaaaa", "bba"},
        };
        for (String[] test : tests) {
            System.out.println(test[0] + " " + test[1]);
            System.out.println(ss.strStr(test[0], test[1]));
            System.out.println();
        }
    }
}
