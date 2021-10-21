run:
	go run cmd/mailer/main.go
build-linux:
	rm -f build/* && cd cmd/mailer && go build && cd - && mkdir -p build && mv cmd/mailer/mailer build/
