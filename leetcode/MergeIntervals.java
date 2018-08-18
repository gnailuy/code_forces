// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/803/

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

class MergeIntervals {
    class IntervalStartComparator implements Comparator<Interval> {
        public int compare(Interval i, Interval j) {
            return Integer.valueOf(i.start).compareTo(j.start);
        }
    }

    public List<Interval> merge(List<Interval> intervals) {
        Collections.sort(intervals, new IntervalStartComparator());
        List<Interval> result = new ArrayList<Interval>();

        Interval current = null;
        for (Interval i: intervals) {
            if (null == current) {
                current = i;
            } else if (current.end >= i.start) {
                current.end = i.end > current.end ? i.end : current.end; // Merge
            } else {
                result.add(current);
                current = i;
            }
        }
        if (null != current) result.add(current);

        return result;
    }

    public static void main(String[] args) {
        MergeIntervals mi = new MergeIntervals();

        int[][][] tests = {
            {{1, 3}, {2, 6}, {8, 10}, {15, 18}},
            {{2, 6}, {8, 10}, {1, 3}, {15, 18}},
            {{1, 4}, {4, 5}},
            {{1, 4}, {2, 3}},
        };
        for (int[][] test : tests) {
            System.out.println(test.length);
            List<Interval> intervals = new ArrayList<Interval>();
            for (int i = 0; i < test.length; i++) {
                System.out.print("(" + test[i][0] + ", " + test[i][1] + ") ");
                intervals.add(new Interval(test[i][0], test[i][1]));
            }
            System.out.println();

            List<Interval> result = mi.merge(intervals);
            for (Interval i: result) {
                System.out.print("(" + i.start + ", " + i.end + ") ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
