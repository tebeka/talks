(defn fibs
  ([] (fibs 1 1))
  ([a b] (lazy-seq (cons a (fibs b (+ a b))))))
