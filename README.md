# Go-Reddit by email

## Todo

* Concurrency: Prendre N subreddits en entrée
* Garder une date de crawl (dans `Subreddit`)
* `UserAgent` configurable

## Learn

* [x] net/http
* [x] fake http (simulate Reddit)
* [x] Logging
* [x] Testing
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

