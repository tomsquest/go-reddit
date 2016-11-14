# Go-Reddit by email

Un binaire simple qui m'envoie par email la liste des Tops des subreddits configurés, toutes les semaines.

## Todo

* Voir DI/IOC idiomatic Go
* Param "fakeReddit" pour ne pas appeler reddit
* Config: SMTP, `UserAgent`
* Gérer les posts sans thumbnails

## Learn

* [x] net/http
* [x] fake http (simulate Reddit)
* [x] Logging
* [x] Testing
* [x] Email
* [x] Templating
* [ ] Config file (Viper ?)
* [ ] Mocking
* [ ] Concurrency

## Maybe

### Technical

* Concurrency: Prendre N subreddits en entrée
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

* Github: Readme, badges, Travis, Cov (?), Vet (?)
* Param pour prendre les tops du jours, de la semaine, du mois...
* Limiter le nombre de posts (5, 10...) ou plus (possible ?)
* User/Password...
* Alternative text au mail html


## Thanks

* HTML template courtesy of Mailchimp: https://github.com/mailchimp/email-blueprints