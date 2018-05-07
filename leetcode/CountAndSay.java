// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/886/

import java.lang.StringBuilder;

class CountAndSay {
    public String countAndSay(int n) {
        String sequence = "1";
        StringBuilder sb = new StringBuilder();
        for (int i = 2; i <= n; i++) {
            char c = sequence.charAt(0);
            int c_cnt = 1;
            for (int j = 1; j < sequence.length(); j++) {
                if (sequence.charAt(j) == c) {
                    c_cnt++;
                } else {
                    sb.append(c_cnt + "" + c);
                    c = sequence.charAt(j);
                    c_cnt = 1;
                }
            }
            sb.append(c_cnt + "" + c);

            sequence = sb.toString();
            sb.setLength(0);
        }
        return sequence;
    }

    public static void main(String[] args) {
        CountAndSay cas = new CountAndSay();

        for (int i = 1; i <= 10; i++) {
            System.out.println(cas.countAndSay(i));
        }
        System.out.println(cas.countAndSay(30));
    }
}
