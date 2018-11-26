import java.util.Arrays;

public class Test {
    public static void main(String[] args) {
        int[] a1 = new int[10];
        int[] a2 = a1;

        for (int i=0; i<a1.length; i++) {
            a1[i] = i*2;
        }

        System.out.println(Arrays.toString(a1));
        System.out.println(Arrays.toString(a2));
    }
}
