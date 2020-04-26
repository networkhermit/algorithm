(defpackage fun.vac.algorithms.number-theory.greatest-common-divisor
  (:nicknames greatest-common-divisor)
  (:use common-lisp)
  (:export
    iterative-binary-gcd
    recursive-binary-gcd
    iterative-euclidean
    recursive-euclidean))

(in-package greatest-common-divisor)

(defun iterative-binary-gcd (m n)
  (when (minusp m)
    (setf m (- m)))
  (when (minusp n)
    (setf n (- n)))

  (do ((shift 0))
    (nil)
    (when (= m n)
      (return-from iterative-binary-gcd (ash m shift)))
    (when (zerop m)
      (return-from iterative-binary-gcd (ash n shift)))
    (when (zerop n)
      (return-from iterative-binary-gcd (ash m shift)))

    (if (zerop (logand m 1))
      (if (zerop (logand n 1))
        (setf m (ash m -1)
              n (ash n -1)
              shift (1+ shift))
        (setf m (ash m -1)))
      (if (zerop (logand n 1))
        (setf n (ash n -1))
        (if (> m n)
          (setf m (ash (- m n) -1))
          (setf n (ash (- n m) -1)))))))

(defun recursive-binary-gcd (m n)
  (when (minusp m)
    (setf m (- m)))
  (when (minusp n)
    (setf n (- n)))

  (when (= m n)
    (return-from recursive-binary-gcd m))
  (when (zerop m)
    (return-from recursive-binary-gcd n))
  (when (zerop n)
    (return-from recursive-binary-gcd m))

  (if (zerop (logand m 1))
    (if (zerop (logand n 1))
      (return-from recursive-binary-gcd (ash (recursive-binary-gcd (ash m -1) (ash n -1)) 1))
      (return-from recursive-binary-gcd (recursive-binary-gcd (ash m -1) n)))
    (if (zerop (logand n 1))
      (return-from recursive-binary-gcd (recursive-binary-gcd m (ash n -1)))
      (if (> m n)
        (return-from recursive-binary-gcd (recursive-binary-gcd (ash (- m n) -1) n))
        (return-from recursive-binary-gcd (recursive-binary-gcd m (ash (- n m) -1)))))))

(defun iterative-euclidean (m n)
  (when (minusp m)
    (setf m (- m)))
  (when (minusp n)
    (setf n (- n)))

  (do (r)
    ((zerop n) m)
    (setf r (rem m n))
    (shiftf m n r)))

(defun recursive-euclidean (m n)
  (when (minusp m)
    (setf m (- m)))
  (when (minusp n)
    (setf n (- n)))

  (if (zerop n)
    m
    (recursive-euclidean n (rem m n))))
