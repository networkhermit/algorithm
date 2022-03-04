import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int n;

            while (input.hasNext()) {
                n = input.nextInt();
                if (n == 42) {
                    break;
                }
                System.out.println(n);
            }
        }
    }
}
