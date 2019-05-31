(define collatz
  (lambda (n)
    (let ((even (= (% n 2) 0)))
      (if even
	  (/ n 2)
	  (+ (* n 3) 1)))))
