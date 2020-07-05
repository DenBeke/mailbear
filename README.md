# 🐻 MailBear: Forms Backend

[![Build Status](https://travis-ci.com/DenBeke/mailbear.svg?branch=master)](https://travis-ci.com/DenBeke/mailbear)
[![Go Report Card](https://goreportcard.com/badge/github.com/DenBeke/mailbear)](https://goreportcard.com/report/github.com/DenBeke/mailbear)
![![Docker Image Size (latest by date)](https://hub.docker.com/r/denbeke/mailbear)](https://img.shields.io/docker/image-size/denbeke/mailbear?sort=date)


MailBear is an open source, self hosted forms backend.
Just do a post request to the API with some form data, and MailBear will make sure the submission is sent to you via mail!

MailBear will always hide the email address of the recepient, since the forms are accessed by a unique key.


## Run with Docker

You can easily run MailBear with Docker:

Copy `config_sample.yml` to `config.yml` and run the server:

    docker run -v $(PWD)/config.yml:/mailbear/config.yml denbeke/mailbear

For your convenience I created a [docker-compose.yml](./docker-compose.yml) file.


## Run in Development

Copy `config_sample.yml` to `config.yml` and run the server:

    go run cmd/mailbear/main.go



## Configuration

Configuration is very simple. Just create as many forms as you want in `config.yml`:


    global:
        smtp:
            host: smtp.example.com
            port: 25
            user:
            password:
            disable_tls: true
            from_email: no-reply@example.com
            from_name: MailBear
        http:
            address: ":1234"
    

    forms:
        some-form-name:
            key: some-random-key
            allowed_domains:
                - localhost:8080
                - example.com
            to_email:
                - recepient@example.com



## Usage

Once MailBear is running you can send requests with form data in the JSON body:

    curl \
        -X POST \
        http://localhost:1234/api/v1/form/some-random-key \
        -H 'Content-Type: application/json' \
        -H 'Origin: http://localhost:8080' \
        -d '{"name":"Joe","email":"joe@example.com", "subject": "Some subject", "content": "Maecenas faucibus mollis interdum. Sed posuere consectetur est at lobortis."}'


## Acknowledgements

* [github.com/labstack/echo](https://github.com/labstack/echo)
* [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [github.com/go-yaml/yaml](https://github.com/go-yaml/yaml)
* [github.com/go-mail/mail](https://github.com/go-mail/mail)
* [github.com/badoux/checkmail](https://github.com/badoux/checkmail)



## Author

[Mathias Beke](https://denbeke.be)