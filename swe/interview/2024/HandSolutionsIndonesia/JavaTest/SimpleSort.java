import java.util.Scanner;

public class SimpleSort {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Enter x value: ");
        int x = scanner.nextInt();
        System.out.print("Enter y value: ");
        int y = scanner.nextInt();
        System.out.print("Enter z value: ");
        int z = scanner.nextInt();

        int min = Math.min(x, Math.min(y, z));
        int max = Math.max(x, Math.max(y, z));
        int median = x + y + z - min - max;

        System.out.println("Sorted from small numbers: " + min + ", " + median + ", " + max);
    }
}
