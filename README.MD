~~~ bash
go vet ./...
go fmt .
~~~

~~~ bash
docker build -t go-calculator .
docker run -p 5000:5000 go-calculator
http://localhost:5000
~~~

~~~ bash
# inside the test/selenium dir
npx mocha runningTests.spec.js
~~~