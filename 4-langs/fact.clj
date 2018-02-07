(defn fact [n]
  (loop [n n acc 1]
    (if (zero? n)
      acc
      (recur (dec n) (* acc n)))))
