(defconstant +lisp-path+ (sb-ext:posix-getenv "LISP_PATH"))

(defun use (module)
  (let ((*default-pathname-defaults*
         (if +lisp-path+
             (let ((pathname (pathname +lisp-path+)))
               (make-pathname :directory (append (or (pathname-directory pathname)
                                                     (list :relative))
                                                 (list (file-namestring pathname)))
                              :name nil
                              :type nil
                              :defaults pathname))
             *default-pathname-defaults*)))
    (load (merge-pathnames (concatenate 'string module ".lisp")))))
