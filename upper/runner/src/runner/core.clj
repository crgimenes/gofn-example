(ns runner.core
  (:gen-class))

(defn upper [stdin]
  (clojure.string/upper-case stdin))

(defn -main []
  (let [stdin (slurp *in*)]
    (print (upper stdin)))
  (flush))
