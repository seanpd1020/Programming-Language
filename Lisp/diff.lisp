;f1 and f2 are lists to store each line of file1.c and file2.cpp
(defparameter f1 nil)
(defparameter f2 nil)

;open file1.c , readline and put in f1 list
(let((in(open "file1.txt":if-does-not-exist nil)))
    (when in
        (loop for line = (read-line in nil)
            while line do
                (if(not(equal line ""))
                    (setf f1 (cons line f1))
                )
        )
        (setf f1 (reverse f1))
        (close in)
    )
)
;open file2.cpp , readline and put in f2 list
(let((in(open "file2.txt":if-does-not-exist nil)))
    (when in
        (loop for line = (read-line in nil)
            while line do
                (if(not(equal line ""))
                    (setf f2 (cons line f2))
                )
        )
        (setf f2 (reverse f2))
        (close in)
    )
)

;do the diff command between f1 and f2
(dolist(tmp1 f1)
    (let((x 0) (c 1))
        (dolist(tmp2 f2)
            (if (equal tmp1 tmp2)
                (progn
                    (dolist(tmp_before_equal f2)
                        (if(equal tmp1 tmp_before_equal)
                            (return)
                            (progn
                                (format t "+~a~%" tmp_before_equal)
                                (setf c (+ c 1))
                            )
                        )
                    )
                    (format t "~a~%" tmp2)
                    (setf x 1)
                    (do((i 0 (+ i 1)))
                        ((> i (- c 1)) 'done)
                        (setf f2 (cdr f2))
                    )
                )
            )
        )
        (if(equal x 0)
            (format t "-~a~%" tmp1)
        )
    )
)

;output the remaining lines of f2
(dolist(tmp2 f2)
    (format t "+~a~%" tmp2)
)

