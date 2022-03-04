(defun main ()
  (write-line "When’s Weekend?")

  (let ((today (nth 6 (multiple-value-list (get-decoded-time)))))

    (cond ((= today 6)
           (write-line "Today."))
          ((= today 5)
           (write-line "Today."))
          ((= today 4)
           (write-line "Tomorrow."))
          ((= today 3)
           (write-line "In two days."))
          (t
            (write-line "Too far away.")))))
