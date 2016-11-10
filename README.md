# Go-Reddit by email

Un binaire simple qui m'envoie par email la liste des Tops des subreddits configurés, toutes les semaines.

## Todo

* Concurrency: Prendre N subreddits en entrée
* `UserAgent` configurable

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

