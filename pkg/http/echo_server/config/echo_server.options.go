package config

type EchoOptions struct {
	Port                int64    `mapstructure:"port"                validate:"required" env:"PORT"`
	Development         bool     `mapstructure:"development"                             env:"Development"`
	BasePath            string   `mapstructure:"basePath"            validate:"required" env:"BasePath"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"                     env:"DebugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
	Timeout             int64    `mapstructure:"timeout"                                 env:"Timeout"`
	Name                string   `mapstructure:"name"                                    env:"Name"`
}

// func (c *EchoOptions) Address() string {
// 	return fmt.Sprintf("%s%s", c.Host, c.Port)
// }

// func (c *EchoOptions) BasePathAddress() string {
// 	path, err := url.JoinPath(c.Address(), c.BasePath)
// 	if err != nil {
// 		return ""
// 	}
// 	return path
// }
