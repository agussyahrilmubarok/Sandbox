public class SimpleSum {
    public static void main(String[] args) {
        int sum = 0;

        for (int i = 1; i <= 100; i++) {
            sum += i;
            System.out.println("Adding " + i + ": Sum so far is " + sum);
        }

        System.out.println("Final sum from 1 to 100 is: " + sum);
    }
}
