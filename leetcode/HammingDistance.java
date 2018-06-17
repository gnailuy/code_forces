// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/762/

class HammingDistance {
    public int hammingDistance(int x, int y) {
        int d = 0;
        int mask = 1;
        for (int i = 0; i < 32; i++) {
            int xm = (x&mask);
            int ym = (y&mask);
            if (xm != ym) d++;
            mask <<= 1;
        }
        return d;
    }

    public static void main(String[] args) {
        HammingDistance hd = new HammingDistance();

        int[][] tests = {
            {1, 4},
        };

        for (int[] test: tests) {
            System.out.println(test[0] + " " + test[1]);
            System.out.println(hd.hammingDistance(test[0], test[1]));
            System.out.println();
        }
    }
}

