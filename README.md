# Go-Reddit by email

## Todo

* Renommer `reddit.New` en `reddit.NewReddit`
* `Reddit` doit avoir une dépendance vers la func PostUnmarshaller
* Concurrency: Prendre N subreddits en entrée
* `UserAgent` configurable

## Learn

* [x] net/http
* [x] fake http (simulate Reddit)
* [x] Logging
* [x] Testing
* [ ] Concurrency
* [ ] Email
* [ ] Templating

## Maybe

### Technical

* Dependency injection
* Switch to `fasthttp`: https://godoc.org/github.com/valyala/fasthttp

### Functional

* Param pour prendre les tops du jours, de la semaine, du mois...
* Limiter le nombre de posts (5, 10...) ou plus (possible ?)
* User/Password...

