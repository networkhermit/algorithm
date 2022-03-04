import java.util.Scanner;

public class Solution {

    public static long process(long n) {
        if (n < 2) {
            return 1;
        }

        while ((n & 1) == 0) {
            n >>>= 1;
        }

        if (n == 1) {
            return 2;
        }

        for (long i = 3, bound = (long) Math.sqrt((double) n) + 1; i < bound; i += 2) {
            while (n % i == 0) {
                if (n == i) {
                    return n;
                }
                n /= i;
            }
        }

        return n;
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int testCases;

            long n;

            while (input.hasNext()) {
                testCases = input.nextInt();
                for (int i = 0; i < testCases; i++) {
                    n = input.nextLong();
                    System.out.println(process(n));
                }
            }
        }
    }
}
