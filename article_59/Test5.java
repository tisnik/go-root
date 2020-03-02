import java.util.List;
import java.util.ArrayList;
import java.awt.Color;

public class Test5 {
    public static void main(String[] args) {
        List<String> l = new ArrayList<String>();
        l.add("foo");
        l.add("bar");
        l.add("baz");
        l.add(Integer.toString(42));

        for (String s : l) {
            System.out.println(s.length());
        }
    }
}

