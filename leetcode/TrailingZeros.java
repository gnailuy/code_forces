// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/816/

class TrailingZeros {
    // Every two numbers there is a factor 2
    // Every five numbers there is a factor 5
    // So just count the factor 5
    public int trailingZeroes(int n) {
        int fiveCount = 0;
        // while (n%5 != 0) n--; // Unnecessary count down
        while (n > 0) {
            fiveCount += n/5;
            n /= 5;
        }
        return fiveCount;
    }

    public static void main(String[] args) {
        TrailingZeros tz = new TrailingZeros();

        for (int i = 0; i <= 100; i++) {
            System.out.println(i + ": " + tz.trailingZeroes(i));
        }
        int test = 511663768;
        System.out.println(test + ": " + tz.trailingZeroes(test));
    }
}

