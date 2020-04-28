## Run tests and generate coverage report

```
go test -v -coverprofile ./cover/cover.out .
go tool cover -html=./cover/cover.out -o ./cover/cover.html
```
