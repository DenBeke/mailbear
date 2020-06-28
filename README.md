# 🐻 MailBear: Forms Backend

**⚠️ WIP ⚠️**

MailBear is an open source, self hosted forms backend.
Just do a post request to API with some form data, and MailBear will make sure the submission is sent to you via mail!

MailBear will always hide the email address of the recepient, since the forms are accessed by a unique key.

## Development

Copy `config_sample.yml` to `config.yml` and run the server:

    go run cmd/mailbear/main.go

Then send the form data in the JSON body:

    curl \
        -X POST \
        http://localhost:1234/api/v1/forms/some-random-key \
        -H 'Content-Type: application/json' \
        -H 'Origin: http://localhost:8080' \
        -d '{"name":"Joe","email":"joe@example.com", "subject": "Some subject", "content": "Maecenas faucibus mollis interdum. Sed posuere consectetur est at lobortis."}'


## Run with Docker

MailBear is published on Dockerhub as [denbeke/mailbear](https://hub.docker.com/repository/docker/denbeke/mailbear). Checkout [docker-compose.yml](./docker-compose.yml) to run it. (Don't forget to copy `config_sample.yml` to `config.yml`)


## Acknowledgements

* [github.com/labstack/echo](https://github.com/labstack/echo)
* [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [github.com/go-yaml/yaml](https://github.com/go-yaml/yaml)
* [github.com/go-mail/mail](https://github.com/go-mail/mail)
* [github.com/badoux/checkmail](https://github.com/badoux/checkmail)



## Author

[Mathias Beke](https://denbeke.be)