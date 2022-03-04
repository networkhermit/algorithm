<?php
    declare(strict_types=1);

    if (count(debug_backtrace()) == 0) {
        $FIRST_GLYPH = 33; // '!'
        $LAST_GLYPH = 126; // '~'
        $GLYPHS_PER_LINE = 72;

        $cursor = $FIRST_GLYPH;

        $i = 0;
        while (true) {
            $i = $cursor;
            for ($j = 0; $j < $GLYPHS_PER_LINE; $j++) {
                print(chr($i));
                if ($i == $LAST_GLYPH) {
                    $i = $FIRST_GLYPH;
                } else {
                    $i++;
                }
            }
            print("\n");
            if ($cursor == $LAST_GLYPH) {
                $cursor = $FIRST_GLYPH;
            } else {
                $cursor++;
            }
        }
    }
?>
