import java.util.List;
import java.util.ArrayList;
import java.awt.Color;

public class Test3 {
    public static void main(String[] args) {
        List l = new ArrayList();
        l.add("foo");
        l.add("bar");
        l.add("baz");

        String s = (String)l.get(0);
        System.out.println(s.length());
    }
}

