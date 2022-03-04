import java.util.Scanner;

public class Solution {

    public static long bruteForce(int sentinel) {
        long sum = 0;
        for (int i = 0; i < sentinel; i++) {
            if (i % 3 == 0 || i % 5 == 0) {
                sum += (long) i;
            }
        }
        return sum;
    }

    public static long arithmeticSeries(int sentinel, int unit) {
        long n = (long) ((sentinel - 1) / unit);

        return (long) unit * (n * (1 + n) >>> 1);
    }

    public static long process(int sentinel) {
        long sum = 0;

        sum += arithmeticSeries(sentinel, 3);
        sum -= arithmeticSeries(sentinel, 15);
        sum += arithmeticSeries(sentinel, 5);

        return sum;
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int testCases;

            int sentinel;

            while (input.hasNext()) {
                testCases = input.nextInt();
                for (int i = 0; i < testCases; i++) {
                    sentinel = input.nextInt();
                    System.out.println(process(sentinel));
                }
            }
        }
    }
}
