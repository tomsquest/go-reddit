package config

import (
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/mgutz/logxi/v1"
	"github.com/spf13/viper"
)

type SmtpConfig struct {
	From string
	To   string
	Host string
	Port int
	User string
	Pass string
}

type Config struct {
	Subreddits []string
	Output     string
	UserAgent  string
	Smtp       SmtpConfig
	FakeReddit bool
}

func Read() (cfg Config, err error) {
	viper.SetDefault("subreddits", []string{"golang"})
	viper.SetDefault("userAgent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0")
	viper.SetDefault("output", "console")
	viper.SetDefault("fakeReddit", false)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("REDDIT")

	viper.SetConfigName("go-reddit")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config")

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return cfg, errwrap.Wrapf("Unable to parse config file: {{err}}", err)
		}

		return cfg, errwrap.Wrapf("Unable to locate config file: {{err}}", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, errwrap.Wrapf("unable to decode into struct: {{err}}", err)
	}

	log.Info(cfg.String())
	return
}

func (c *Config) String() string {
	return fmt.Sprintf(`Using configuration:
- Subreddits:  %v
- User Agent:  %v
- Output:      %v
- Fake Reddit: %v
- Smtp:
	- From:    %v
	- To  :    %v
	- Host:    %v
	- Port:    %v
	- User:    %v
	- Pass:    %v
	`, c.Subreddits, c.UserAgent, c.Output, c.FakeReddit,
		c.Smtp.From, c.Smtp.To, c.Smtp.Host, c.Smtp.Port, c.Smtp.User, "(hidden)")
}
