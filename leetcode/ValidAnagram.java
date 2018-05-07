// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/882/

class ValidAnagram {
    public boolean isAnagram(String s, String t) { // All lower case alphabets
        if (s.length() != t.length())
            return false;

        int[] cnt = new int[26];// All zero, for 26 letters
        for (int i = 0; i < s.length(); i++) {
            int idx = (int)s.charAt(i) - (int)'a';
            cnt[idx]++;
        }
        for (int i = 0; i < t.length(); i++) {
            int idx = (int)t.charAt(i) - (int)'a';
            cnt[idx]--;
        }
        for (int i = 0; i < 26; i++) {
            if (0 != cnt[i])
                return false;
        }
        return true;
    }

    public static void main(String[] args) {
        ValidAnagram va = new ValidAnagram();

        String[][] tests = {
            {"anagram", "nagaram"},
            {"rat", "car"},
        };
        for (String[] test : tests) {
            System.out.println(test[0] + " " + test[1]);
            System.out.println(va.isAnagram(test[0], test[1]));
        }
    }
}
