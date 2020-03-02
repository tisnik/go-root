import java.util.List;
import java.util.ArrayList;
import java.awt.Color;

public class Test4 {
    public static void main(String[] args) {
        List<String> l = new ArrayList<String>();
        l.add(new Object());
        l.add("foobar");
        l.add(42);
        l.add(Color.green);

        for (Object i : l) {
            System.out.println(i);
        }
    }
}

