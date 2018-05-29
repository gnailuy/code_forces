// https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/774/

class FirstBadVersion {
    private boolean[] is_bad;

    public FirstBadVersion(String versions) {
        is_bad = new boolean[versions.length()+1];
        for (int i = 0; i < versions.length(); i++) {
            is_bad[i+1] = ('1' == versions.charAt(i));
        }
    }

    boolean isBadVersion(int version) {
        return is_bad[version];
    }

    public int firstBadVersion(int n) {
        int s = 1, e = n;
        while (true) {
            if (s == e) {
                if (1 <= s && s <= n && isBadVersion(s)) {
                    return s;
                } else if (s < n && isBadVersion(s+1)) {
                    return s+1;
                } else {
                    return -1;
                }
            } else if (s > e) {
                return s;
            }

            int middle = s + (e-s) / 2;
            if (isBadVersion(middle)) {
                e = middle - 1;
            } else {
                s = middle + 1;
            }
        }
    }

    public static void main(String[] args) {

        String[] tests = {
            "00011",
            "00000",
            "11111",
            "0",
            "01",
            "001",
            "011",
            "111",
            "11",
            "1",
        };
        for (String test : tests) {
            FirstBadVersion fbv = new FirstBadVersion(test);
            System.out.println(test.length());
            System.out.println(test);
            for (int i = 0; i < test.length(); i++) {
                System.out.print(fbv.isBadVersion(i+1) + " ");
            }
            System.out.println();
            System.out.println(fbv.firstBadVersion(test.length()));
            System.out.println();
        }
    }
}
