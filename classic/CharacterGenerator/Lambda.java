public class Lambda {

    public static void main(String[] args) {
        final int FIRST_GLYPH = 33; // '!'
        final int LAST_GLYPH = 126; // '~'
        final int GLYPHS_PER_LINE = 72;

        int cursor = FIRST_GLYPH;

        int i;
        while (true) {
            i = cursor;
            for (int j = 0; j < GLYPHS_PER_LINE; j++) {
                System.out.print((char) i);
                if (i == LAST_GLYPH) {
                    i = FIRST_GLYPH;
                } else {
                    i++;
                }
            }
            System.out.println();
            if (cursor == LAST_GLYPH) {
                cursor = FIRST_GLYPH;
            } else {
                cursor++;
            }
        }
    }
}
