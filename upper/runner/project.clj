(defproject runner "0.1.1"
  :description "Transforms string that arrives in stdin for uppercase and send to stdout"
  :url "https://github.com/crgimenes/gofn-example/tree/master/upper/runner"
  :license {:name "MIT License"
            :url "https://github.com/crgimenes/gofn-example/blob/master/LICENSE"}
  :dependencies [[org.clojure/clojure "1.8.0"]]
  :main runner.core
  :aot :all)
