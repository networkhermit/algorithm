import java.util.Scanner;

public class Main {

    public static void process(int m, int n) {
        if (n == 1) {
            return;
        }

        if (m <= 2) {
            System.out.println(2);
            m = 3;
        }

        if (n == 2) {
            return;
        }

        m |= 1;
        n = (n - 1) | 1;

        int cache = (int) Math.sqrt((double) n);

        boolean[] sievedFactors = new boolean[(cache + 1) >>> 1];

        for (int i = 3, bound = cache + 1; i < bound; i += 2) {
            if (sievedFactors[i >>> 1]) {
                continue;
            }
            for (int j = i * i; j < bound; j += i << 1) {
                if (!sievedFactors[j >>> 1]) {
                    sievedFactors[j >>> 1] = true;
                }
            }
        }

        boolean[] nums = new boolean[((n - m) >>> 1) + 1];

        int factor;
        int target;

        boolean boost = m <= 9;

        for (int i = 1, length = sievedFactors.length; i < length; i++) {
            if (sievedFactors[i]) {
                continue;
            }

            factor = (i << 1) + 1;

            if (boost) {
                target = factor * factor;
            } else if (m <= factor * factor) {
                boost = true;
                target = factor * factor;
            } else {
                if (m % factor == 0) {
                    target = m;
                } else {
                    target = m + (factor - m % factor);
                    if ((target & 1) == 0) {
                        target += factor;
                    }
                }
            }

            for (int j = target; j <= n; j += factor << 1) {
                if (!nums[(j - m) >>> 1]) {
                    nums[(j - m) >>> 1] = true;
                }
            }
        }

        for (int i = 0, length = nums.length; i < length; i++) {
            if (!nums[i]) {
                System.out.println(m + (i << 1));
            }
        }
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int testCases;

            int m;
            int n;

            while (input.hasNext()) {
                testCases = input.nextInt();
                for (int i = 0; i < testCases; i++) {
                    m = input.nextInt();
                    n = input.nextInt();
                    if (i != 0) {
                        System.out.println();
                    }
                    process(m, n);
                }
            }
        }
    }
}
