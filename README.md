# Go-Reddit by email

Simple binary sending an email the list of the Top posts of Reddit's subs.

## Todo

### Func

- [ ] Oauth: I actually got ban running too many auth
- [ ] Speed: Fetch subreddits concurrently
- [ ] Pimp with some badges
- [ ] Config param for http timeout
- [ ] Reduce template size
- [ ] Configuration for tops of the day/week/month/all
- [ ] Configuration for the number of posts
- [ ] Send text with the html
- [ ] Command line parameter to select the config file
- [ ] Command line parameters to override config file
- [ ] Rename `output.SmtpOutput` to `output.Smtp`
- [ ] No dependency on Viper: use plain Json config file
- [ ] Explore `fasthttp`: https://godoc.org/github.com/valyala/fasthttp
- [ ] Review error handling in `mail.Embed` of `SmtpOutput` 

## Thanks

- HTML template courtesy of Mailchimp: https://github.com/mailchimp/email-blueprints
