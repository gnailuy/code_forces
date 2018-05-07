// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/881/A

class FirstUniqueChar {
    public int firstUniqChar(String s) { // All lower case alphabets
        int[] cnt = new int[26];// All zero, for 26 letters
        for (int i = 0; i < s.length(); i++) {
            int idx = (int)s.charAt(i) - (int)'a';
            cnt[idx]++;
        }
        for (int i = 0; i < s.length(); i++) {
            int idx = (int)s.charAt(i) - (int)'a';
            if (1 == cnt[idx]) {
                return i;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        FirstUniqueChar fuc = new FirstUniqueChar();

        String[] tests = {
            "leetcode",
            "loveleetcode",
            "aabbccdd",
        };
        for (String test : tests) {
            System.out.println(test.length());
            System.out.println(test);
            System.out.println(fuc.firstUniqChar(test));
        }
    }
}
