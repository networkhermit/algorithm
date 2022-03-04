import java.util.Scanner;

public class Solution {

    public static int draft(int n) {
        if (n == 1) {
            return 2;
        }

        int[] primes = new int[n];

        primes[0] = 2;

        int k = 1;

        boolean isPrime;
        int bound;

        for (int i = 3; i < 2147483647; i += 2) {
            isPrime = true;
            bound = (int) Math.sqrt((double) i);
            for (int j = 1; j < k; j++) {
                if (primes[j] > bound) {
                    break;
                }
                if (i % primes[j] == 0) {
                    isPrime = false;
                    break;
                }
            }
            if (isPrime) {
                primes[k] = i;
                k++;
                if (k == n) {
                    break;
                }
            }
        }

        return primes[n - 1];
    }

    public static int process(int n) {
        if (n == 1) {
            return 2;
        }

        int ceiling;

        if (n < 6) {
            ceiling = (n << 1) + 1;
        } else {
            double cache = Math.log((double) n);
            ceiling = (int) ((double) n * (cache + Math.log(cache)));
        }

        int size = (ceiling + 1) >>> 1;

        boolean[] arr = new boolean[size];

        for (int i = 3, bound = (int) Math.sqrt((double) ceiling) + 1; i < bound; i += 2) {
            if (arr[i >>> 1]) {
                continue;
            }
            for (int j = i * i; j <= ceiling; j += i << 1) {
                if (!arr[j >>> 1]) {
                    arr[j >>> 1] = true;
                }
            }
        }

        int k = 1;
        for (int i = 1; i < size; i++) {
            if (!arr[i]) {
                k++;
                if (k == n) {
                    return (i << 1) + 1;
                }
            }
        }

        return 0;
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
