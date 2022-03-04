(defun main ()
  (defconstant +first-glyph+ 33) ; '!'
  (defconstant +last-glyph+ 126) ; '~'
  (defconstant +glyphs-per-line+ 72)

  (do ((cursor +first-glyph+)

       (i 0 (1+ i)))
    (nil)
    (setf i cursor)
    (dotimes (j +glyphs-per-line+)
      (write-char (code-char i))
      (if (= i +last-glyph+)
        (setf i +first-glyph+)
        (incf i)))
    (terpri)
    (if (= cursor +last-glyph+)
      (setf cursor +first-glyph+)
      (incf cursor))))
