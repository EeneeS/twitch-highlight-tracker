package config

type Config struct {
  Server string
  Channel string
  Keywords []string
}

func LoadConfig() *Config {
  return &Config{
    Server: "irc.chat.twitch.tv:6667",
    Channel: "caedrel",
    Keywords: []string{"LOL", "OOOO"},
  }
}
