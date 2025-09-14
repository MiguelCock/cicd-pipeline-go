~~~ bash
go vet ./...
go fmt .
~~~

~~~ bash
docker build -t go-calculator .
docker run -p 8000:8000 go-calculator
http://localhost:8000
~~~

~~~ bash
# inside the test/selenium dir
npx mocha runningTests.spec.js
~~~