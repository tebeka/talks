(define (fact n)
  (let loop ((n n) (acc 1))
    (if (zero? n)
      acc
      (loop (- n 1) (* n acc)))))
