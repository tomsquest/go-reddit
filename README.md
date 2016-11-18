# Go-Reddit by email

Un binaire simple qui m'envoie par email la liste des Tops des subreddits configurés, toutes les semaines.

## Todo

### Must before public release

* Fix: requêter les TOP: `https://www.reddit.com/r/golang/top/?sort=top&t=week`
* Github: Readme, badges, Travis, Cov (?), Vet (?)
* Simplifier les loggers pour ne pas avoir N packages
* Always use `logxi` instead of GO `log`
* Vérifier si vraiment besoin de `errwrap`

### Nice to have (after release)

#### Functional

* Config param for http timeout
* Reduce template size
* Param pour prendre les tops du jours, de la semaine, du mois...
* Limiter le nombre de posts (5, 10...) ou plus (possible ?)
* Auth User/Password...
* Get front page
* Alternative text au mail html

#### Technical 

* No dependency on Viper: use plain Json config file
* Concurrency: Prendre N subreddits en entrée
* Explore Go mocking libs
* Explore `fasthttp`: https://godoc.org/github.com/valyala/fasthttp

## Thanks

* HTML template courtesy of Mailchimp: https://github.com/mailchimp/email-blueprints