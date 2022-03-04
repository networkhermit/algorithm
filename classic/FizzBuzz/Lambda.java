public class Lambda {

    public static String fizzBuzz(int n) {
        String word = "";

        if (n % 3 == 0) {
            word += "Fizz";
        }
        if (n % 5 == 0) {
            word += "Buzz";
        }

        if (word.isEmpty()) {
            word = String.valueOf(n);
        }

        return word;
    }

    public static void main(String[] args) {
        for (int i = 1; i <= 100; i++) {
            System.out.println(fizzBuzz(i));
        }
    }
}
