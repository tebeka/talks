; One step in Collatz conjecture
(define collatz
  (lambda (n)
    (if (even? n)
	(/ n 2)
	(+ (* n 3) 1))))


(collatz 7) ; 22
