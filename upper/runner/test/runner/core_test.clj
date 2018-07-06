(ns runner.core-test
  (:require [clojure.test :refer :all]
            [runner.core :refer :all]))

(deftest test-upper-defn
  (testing "test abc"
    (is (= "TEST ABC" (upper "test abc"))))
  (testing "test ONE string upper"
    (is (= "TEST ONE STRING UPPER" (upper "test ONE string upper"))))
  (testing "Not equal"
    (is (not= "Not equal" (upper "NOT EQUAL")))))
