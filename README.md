go mod init gogin

go run main.go

go test ./platform/newsfeed
go test ./...
go test -cover ./...