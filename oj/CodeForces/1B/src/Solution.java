import java.util.Scanner;
import java.util.Stack;

public class Solution {

    public static void process(String coordinate) {
        int column = 0;
        int row = 0;

        int length = coordinate.length();

        int k = 1;
        if (coordinate.charAt(0) == 'R' && !Character.isLetter(coordinate.charAt(k))) {
            while (k < length && Character.isDigit(coordinate.charAt(k))) {
                k++;
            }
        }

        if (k == 1 || k == length) {
            k = 0;
            while (Character.isLetter(coordinate.charAt(k))) {
                column = column * 26 + ((int) coordinate.charAt(k) - 64);
                k++;
            }
            while (k < length) {
                row = row * 10 + ((int) coordinate.charAt(k) - 48);
                k++;
            }

            System.out.printf("R%dC%d%n", row, column);
        } else {
            k = 1;
            while (Character.isDigit(coordinate.charAt(k))) {
                row = row * 10 + ((int) coordinate.charAt(k) - 48);
                k++;
            }
            k++;
            while (k < length) {
                column = column * 10 + ((int) coordinate.charAt(k) - 48);
                k++;
            }

            Stack<Character> s = new Stack<>();

            int r;
            while (column != 0) {
                r = column % 26;
                column /= 26;
                if (r == 0) {
                    s.push('Z');
                    column--;
                } else {
                    s.push((char) (r + 64));
                }
            }

            while (!s.empty()) {
                System.out.print(s.peek());
                s.pop();
            }
            System.out.println(row);
        }
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in, "UTF-8")) {
            int testCases;

            String coordinate;

            while (input.hasNext()) {
                testCases = input.nextInt();
                for (int i = 0; i < testCases; i++) {
                    coordinate = input.next();
                    process(coordinate);
                }
            }
        }
    }
}
