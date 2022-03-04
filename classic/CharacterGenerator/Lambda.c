#include <stdbool.h>
#include <stdio.h>

int main(void) {
    const int FIRST_GLYPH = 33; // '!'
    const int LAST_GLYPH = 126; // '~'
    const int GLYPHS_PER_LINE = 72;

    int cursor = FIRST_GLYPH;

    int i;
    while (true) {
        i = cursor;
        for (int j = 0; j < GLYPHS_PER_LINE; j++) {
            putchar(i);
            if (i == LAST_GLYPH) {
                i = FIRST_GLYPH;
            } else {
                i++;
            }
        }
        putchar('\n');
        if (cursor == LAST_GLYPH) {
            cursor = FIRST_GLYPH;
        } else {
            cursor++;
        }
    }
}
