# Go Email Scheduler

This is a simple email scheduler that sends emails at a specific time. It is built using Golang, MySQL DB and [MailGun](https://www.mailgun.com/) API.

## Pre-requisites

1. Go > v1.16
2. MySQL
3. MailGun account, activate account and create API KEY and DOMAIN (Follow this [Guide](https://documentation.mailgun.com/docs/mailgun/user-manual/get-started/) for reference)

## Installation

1. Clone the repository
2. Run `go mod tidy` to install the dependencies
3. Create a `cmd/main/.env` file in the root directory and add the following environment variables:
    - `MAILGUN_DOMAIN` - Your MailGun API key
    - `MAILGUN_API_KEY` - Your MailGun domain
    - `TICKER_DURATION` - (in seconds), the duration of the ticker to check db for active emails
    - `DB_CONNECTION_STRING` - The connection string to your MySQL database
4. Run `go run cmd/main/main.go` to start the server

## Usage

1. Create a new email by sending a `POST` request to `http://localhost:8080/email` with the following JSON payload:
    ```json
    {
        "Subject": "sending  email from Go email scheduler",
        "Body": "Hello to cyberpunk, from past",
        "Status" : "not sent",
        "To": "myemailisawesome@gmail.com",
        "Time": "2077-09-15T11:58:24.945Z"
    }
    ```
2. The email will be sent at the specified time
3. To view all emails, send a `GET` request to `http://localhost:8080/email`
4. To view a specific email, send a `GET` request to `http://localhost:8080/email/{id}`
5. To update an email, send a `PUT` request to `http://localhost:8080/email/{id}` with the following JSON payload:
    ```json
    {
        "Subject": "sending  email from Go email scheduler",
        "Body": "Hello to cyberpunk, from past",
        "Status" : "not sent",
        "To": "myemailisawesomeedited@gmail.com",
        "Time": "2077-09-15T11:58:24.945Z"
    }
    ```
6. To delete an email, send a `DELETE` request to `http://localhost:8080/email/{id}`


## Build

To build the project, run `go build cmd/main/main.go`


