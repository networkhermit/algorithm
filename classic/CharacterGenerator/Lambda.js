"use strict"

if (module === require.main) {
    const FIRST_GLYPH = 33 // '!'
    const LAST_GLYPH = 126 // '~'
    const GLYPHS_PER_LINE = 72

    let cursor = FIRST_GLYPH

    let i = 0
    while (true) {
        i = cursor
        for (let j = 0; j < GLYPHS_PER_LINE; j++) {
            process.stdout.write(String.fromCharCode(i))
            if (i === LAST_GLYPH) {
                i = FIRST_GLYPH
            } else {
                i++
            }
        }
        console.log()
        if (cursor === LAST_GLYPH) {
            cursor = FIRST_GLYPH
        } else {
            cursor++
        }
    }
}
