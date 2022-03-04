import java.util.Calendar;

public class Lambda {

    public static void main(String[] args) {
        System.out.println("When’s Weekend?");

        int today = Calendar.getInstance().get(Calendar.DAY_OF_WEEK);

        switch (today) {
        case Calendar.SUNDAY:
            System.out.println("Today.");
            break;
        case Calendar.SATURDAY:
            System.out.println("Today.");
            break;
        case Calendar.FRIDAY:
            System.out.println("Tomorrow.");
            break;
        case Calendar.THURSDAY:
            System.out.println("In two days.");
            break;
        default:
            System.out.println("Too far away.");
        }
    }
}
