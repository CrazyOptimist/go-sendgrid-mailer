# Golang Sendgrid XLSX Mailer

## Description
Twilio Sendgrid provides a free tier so that you can send 100 emails per day for free.  
Create an account at sendgrid.com and get your API key.  
Put the API key in `.env` file.  
Create a XLSX file where the first column contains an array of emails.  
It's enough to put the built binary, env file, XLSX file in the same directory and you are good to go.  

## Run
```bash
make run
```
or
```bash
go run cmd/mailer/*.go
```

## Build
For Linux
```bash
make build-linux
```

For Windows
```bash
make build-windows
```

Created by [CrazyOptimist](https://github.com/crazyoptimist) with :heart:
