package config

type Config struct {
	Server   string
	Channel  string
	Keywords []string
}

func LoadConfig() *Config {
	return &Config{
		Server:   "irc.chat.twitch.tv:6667",
		Channel:  "dima_wallhacks",
		Keywords: []string{"LOL", "OOOO", "aga", "waga"},
	}
}
