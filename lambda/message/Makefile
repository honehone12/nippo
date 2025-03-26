.PHONY:

build:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

zip:
	zip nippo_message.zip bootstrap