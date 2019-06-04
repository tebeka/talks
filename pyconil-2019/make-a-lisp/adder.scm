(define make-adder
  (lambda (n)
    (lambda (m) (+ m n))))

(define add-7 (make-adder 7))

(add-7 10) ; 17
