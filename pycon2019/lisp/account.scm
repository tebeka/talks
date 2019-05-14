(define make-account
  (lambda (balance)
    (lambda (amt) 
      (begin (set! balance (+ balance amt)) 
             balance))))

(define acct (make-account 100))
