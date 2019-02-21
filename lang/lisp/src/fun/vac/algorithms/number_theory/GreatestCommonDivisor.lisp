(defpackage fun.vac.algorithms.number-theory.greatest-common-divisor
  (:nicknames greatest-common-divisor)
  (:use common-lisp)
  (:export
    iterative-binary-gcd
    recursive-binary-gcd
    iterative-euclidean
    recursive-euclidean
    ))

(in-package greatest-common-divisor)

(defun iterative-binary-gcd (m n)
  (when (< m 0)
    (setf m (- m)))
  (when (< n 0)
    (setf n (- n)))

  (let ((shift 0))

    (do ()
      (nil)
      (when (= m n)
        (return-from iterative-binary-gcd (ash m shift)))
      (when (= m 0)
        (return-from iterative-binary-gcd (ash n shift)))
      (when (= n 0)
        (return-from iterative-binary-gcd (ash m shift)))

      (if (= (logand m 1) 0)
        (if (= (logand n 1) 0)
          (progn
            (setf m (ash m -1))
            (setf n (ash n -1))
            (incf shift))
          (setf m (ash m -1)))
        (if (= (logand n 1) 0)
          (setf n (ash n -1))
          (if (> m n)
            (setf m (ash (- m n) -1))
            (setf n (ash (- n m) -1))))))))

(defun recursive-binary-gcd (m n)
  (when (< m 0)
    (setf m (- m)))
  (when (< n 0)
    (setf n (- n)))

  (when (= m n)
    (return-from recursive-binary-gcd m))
  (when (= m 0)
    (return-from recursive-binary-gcd n))
  (when (= n 0)
    (return-from recursive-binary-gcd m))

  (if (= (logand m 1) 0)
    (if (= (logand n 1) 0)
      (return-from recursive-binary-gcd (ash (recursive-binary-gcd (ash m -1) (ash n -1)) 1))
      (return-from recursive-binary-gcd (recursive-binary-gcd (ash m -1) n)))
    (if (= (logand n 1) 0)
      (return-from recursive-binary-gcd (recursive-binary-gcd m (ash n -1)))
      (if (> m n)
        (return-from recursive-binary-gcd (recursive-binary-gcd (ash (- m n) -1) n))
        (return-from recursive-binary-gcd (recursive-binary-gcd m (ash (- n m) -1)))))))

(defun iterative-euclidean (m n)
  (when (< m 0)
    (setf m (- m)))
  (when (< n 0)
    (setf n (- n)))

  (let (r)
    (do ()
      ((= n 0))
      (setf r (rem m n))
      (shiftf m n r)))

  m)

(defun recursive-euclidean (m n)
  (when (< m 0)
    (setf m (- m)))
  (when (< n 0)
    (setf n (- n)))

  (if (= n 0)
    m
    (recursive-euclidean n (rem m n))))
