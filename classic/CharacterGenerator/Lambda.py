def main() -> None:
    FIRST_GLYPH = 33 # '!'
    LAST_GLYPH = 126 # '~'
    GLYPHS_PER_LINE = 72

    cursor = FIRST_GLYPH

    i = 0
    while True:
        i = cursor
        for _ in range(GLYPHS_PER_LINE + 1):
            print(chr(i), end = "")
            if i == LAST_GLYPH:
                i = FIRST_GLYPH
            else:
                i += 1
        print()
        if cursor == LAST_GLYPH:
            cursor = FIRST_GLYPH
        else:
            cursor += 1

if __name__ == "__main__":
    main()
