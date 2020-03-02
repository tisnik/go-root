import java.util.Collection;
import java.util.ArrayList;
import java.awt.Color;

public class Test6 {
    public static void main(String[] args) {
        Collection<String> l1 = new ArrayList<String>();
        Collection<Integer> l2 = new ArrayList<Integer>();

        System.out.println(l1.getClass().getName());
        System.out.println(l2.getClass().getName());

        System.out.println(l1.getClass() == l2.getClass());
    }
}

