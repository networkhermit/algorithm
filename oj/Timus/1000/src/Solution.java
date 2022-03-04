import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int a;
            int b;

            while (input.hasNext()) {
                a = input.nextInt();
                b = input.nextInt();
                System.out.println(a + b);
            }
        }
    }
}
