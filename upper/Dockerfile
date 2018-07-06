FROM clojure:alpine

ADD runner /srv/app
WORKDIR /srv/app
RUN lein uberjar

CMD ["java", "-jar", "target/runner-0.1.1-standalone.jar"]
