(defun fizz-buzz (n)
  (let ((word ""))

    (when (zerop (rem n 3))
      (setf word (concatenate 'string word "Fizz")))
    (when (zerop (rem n 5))
      (setf word (concatenate 'string word "Buzz")))

    (when (zerop (length word))
      (setf word (write-to-string n)))

    word))

(defun main ()
  (do ((i 1 (1+ i)))
    ((> i 100))
    (write-line (fizz-buzz i))))
