# 🐻 MailBear: Forms Backend

**⚠️ WIP ⚠️**

MailBear is an open source, self hosted forms backend.
Just do a post request to API with some form data, and MailBear will make sure the submission is sent to you via mail!


## Development

Copy `config_sample.yml` to `config.yml` and run the server:

    go run cmd/mailbear/main.go

Then send the form data in the JSON body:

    curl \
        -X POST \
        http://localhost:1234/api/v1/forms/some-random-key \
        -H 'Content-Type: application/json' \
        -d '{"name":"Joe","email":"joe@example.com", "subject": "Some subject", "content": "Maecenas faucibus mollis interdum. Sed posuere consectetur est at lobortis."}'



## Acknowledgements

* [github.com/labstack/echo](https://github.com/labstack/echo)
* [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [github.com/go-yaml/yaml](https://github.com/go-yaml/yaml)
* [github.com/go-mail/mail](https://github.com/go-mail/mail)
* [github.com/badoux/checkmail](https://github.com/badoux/checkmail)



## Author

[Mathias Beke](https://denbeke.be)