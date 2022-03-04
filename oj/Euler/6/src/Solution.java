import java.util.Scanner;

public class Solution {

    public static long draft(int n) {
        long v = (long) n;
        long result = 0;

        for (long i = 1; i < v; i++) {
            result += i * (v - i) * (i + 1 + v);
        }

        return result;
    }

    public static long process(int n) {
        long v = (long) n;
        return (v * (v + 1) >>> 1) * (v - 1) / 3 * (3 * v + 2) >>> 1;
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int testCases;

            int n;

            while (input.hasNext()) {
                testCases = input.nextInt();
                for (int i = 0; i < testCases; i++) {
                    n = input.nextInt();
                    System.out.println(process(n));
                }
            }
        }
    }
}
