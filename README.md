# Go-Reddit by email

Un binaire simple qui m'envoie par email la liste des Tops des subreddits configurés, toutes les semaines.

## Todo

* Voir DI/IOC idiomatic Go
* Concurrency: Prendre N subreddits en entrée
* Github: Readme, badges, Travis, Cov (?), Vet (?)

* Param "fakeReddit" pour ne pas appeler reddit
* Template: CSS Inline (Douceur) ou à la mano
* Config: SMTP, `UserAgent`

## Learn

* [x] net/http
* [x] fake http (simulate Reddit)
* [x] Logging
* [x] Testing
* [ ] Config file (Viper ?)
* [ ] Mocking
* [ ] Concurrency
* [ ] Email
* [ ] Templating

## Maybe

### Technical

* Fwk de Mock
* Dependency injection
* Switch to `fasthttp`: https://godoc.org/github.com/valyala/fasthttp
* `http.Client` pourrait appeler la méthode `unmarshall`:

```
type Entity interface {
    UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
    return v.UnmarshalHTTP(r)
}
```

### Functional

* Param pour prendre les tops du jours, de la semaine, du mois...
* Limiter le nombre de posts (5, 10...) ou plus (possible ?)
* User/Password...
* Alternative text au mail html


## Thanks

* HTML template courtesy of Mailchimp: https://github.com/mailchimp/email-blueprints