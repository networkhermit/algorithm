if __FILE__ == $PROGRAM_NAME
    FIRST_GLYPH = 33 # '!'
    LAST_GLYPH = 126 # '~'
    GLYPHS_PER_LINE = 72

    cursor = FIRST_GLYPH

    i = 0
    loop do
        i = cursor
        GLYPHS_PER_LINE.times do
            putc(i)
            if i > LAST_GLYPH
                i = FIRST_GLYPH
            else
                i += 1
            end
        end
        puts()
        if cursor > LAST_GLYPH
            cursor = FIRST_GLYPH
        else
            cursor += 1
        end
    end
end
