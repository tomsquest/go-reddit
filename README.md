# Go-Reddit by email

Un binaire simple qui m'envoie par email la liste des Tops des subreddits configurés, toutes les semaines.

## Todo

### Must before public release

* Ne prendre qu'un subreddit en entrée en attendant la concurrence
* Github: Readme, badges, Travis, Cov (?), Vet (?)
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
* Command line parameter to select the config file
* Command line parameters to override config file

#### Technical 

* Rename `output.SmtpOutput` to `output.Smtp`
* No dependency on Viper: use plain Json config file
* Concurrency: Prendre N subreddits en entrée
* Explore Go mocking libs
* Explore `fasthttp`: https://godoc.org/github.com/valyala/fasthttp

## Thanks

* HTML template courtesy of Mailchimp: https://github.com/mailchimp/email-blueprints