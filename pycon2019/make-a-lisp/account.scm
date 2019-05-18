(define make-account
  (lambda (balance)
    (lambda (amount)
      (begin (set! balance (+ balance amount))
             balance))))

(define acct (make-account 100))
