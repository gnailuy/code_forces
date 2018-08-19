// https://leetcode.com/explore/interview/card/top-interview-questions-medium/112/design/813/

import java.util.HashMap;
import java.util.Random;

class RandomizedSet {
    int size;
    Random rand;
    HashMap<Integer, Integer> valueIndex;
    HashMap<Integer, Integer> indexValue;

    /** Initialize your data structure here. */
    public RandomizedSet() {
        size = 0;
        rand = new Random();
        valueIndex = new HashMap<Integer, Integer>();
        indexValue = new HashMap<Integer, Integer>();
    }

    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
    public boolean insert(int val) {
        if (valueIndex.containsKey(val)) return false;

        size++;
        valueIndex.put(val, size);
        indexValue.put(size, val);

        return true;
    }

    /** Removes a value from the set. Returns true if the set contained the specified element. */
    public boolean remove(int val) {
        if (!valueIndex.containsKey(val)) return false;

        int index = valueIndex.get(val); // Index of target value
        if (index != size) { // Value is not the last one
            indexValue.put(index, indexValue.get(size)); // Put last value to index
            valueIndex.put(indexValue.get(size), index); // Fix value index map
        }
        indexValue.remove(size); // Delete last index
        valueIndex.remove(val); // Delete val
        size--;

        return true;
    }

    /** Get a random element from the set. */
    public int getRandom() {
        int randomIndex = rand.nextInt(size); // 0~size-1
        return indexValue.get(randomIndex+1);
    }

    public static void main(String[] args) {
        RandomizedSet rs = new RandomizedSet();
        System.out.println("Insert 1: " + rs.insert(1));
        System.out.println("Insert 2: " + rs.insert(2));
        System.out.println("Insert 3: " + rs.insert(3));
        System.out.println("Insert 4: " + rs.insert(4));
        System.out.println("Insert 5: " + rs.insert(5));
        System.out.println("Remove 3: " + rs.remove(3));
        System.out.println("Get Random: " + rs.getRandom());
    }
}

