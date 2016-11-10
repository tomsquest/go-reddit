# Go-Reddit by email

Un binaire simple qui m'envoie par email la liste des Tops des subreddits configurés, toutes les semaines.

## Todo

* Titre cliquable dans la template
* Texte cliquable
* Date du crawl en date+time (template)
* Param "fakeReddit" pour ne pas appeler reddit
* Config SMTP
* Config `UserAgent`
* Concurrency: Prendre N subreddits en entrée

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


## Thanks

* HTML template courtesy of Mailchimp: https://github.com/mailchimp/email-blueprints