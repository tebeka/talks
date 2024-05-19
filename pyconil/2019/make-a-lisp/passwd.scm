(define login
  (lambda (user password)
    (let ((given (hash-password password))
	  (stored (load-user-password user)))
      (= given stored))))
