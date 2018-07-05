(ns runner.core
  (:gen-class))

(defn upper []
  (print (clojure.string/upper-case (slurp *in*)))
  (flush))

(defn -main []
  (upper))