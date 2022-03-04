import java.util.Scanner;

public class Solution {

    public static long process(int n) {
        if (n < 2) {
            return 1;
        }

        int size = (n + 1) >>> 1;

        boolean[] sievedFactors = new boolean[size];

        for (int i = 3, bound = (int) Math.sqrt((double) n) + 1; i < bound; i += 2) {
            if (sievedFactors[i >>> 1]) {
                continue;
            }
            for (int j = i * i; j <= n; j += i << 1) {
                if (!sievedFactors[j >>> 1]) {
                    sievedFactors[j >>> 1] = true;
                }
            }
        }

        long result = 1;

        for (int i = 2; i <= n; i <<= 1) {
            result <<= 1;
        }

        int factor;
        for (int i = 1; i < size; i++) {
            if (!sievedFactors[i]) {
                factor = (i << 1) + 1;
                for (int j = factor; j <= n; j *= factor) {
                    result *= (long) factor;
                }
            }
        }

        return result;
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
