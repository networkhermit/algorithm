import java.util.Scanner;

public class Solution {

    public static long process(int n, int m, int a) {
        int x = (n - 1) / a + 1;
        int y = (m - 1) / a + 1;

        return (long) x * (long) y;
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int n;
            int m;
            int a;

            while (input.hasNext()) {
                n = input.nextInt();
                m = input.nextInt();
                a = input.nextInt();
                System.out.println(process(n, m, a));
            }
        }
    }
}
