package rem

type Config struct {
	DefaultHandler Handler
	BodyProcessor  IBodyProcessor
}

var config *Config = nil

func SetConfig(s *Config) {
	config = s
}

func DefaultConfig() *Config {
	return &Config{
		DefaultHandler: func(ctx *Context) IResponse { return Success(nil) },
		BodyProcessor:  &JSONBodyProcessor{},
	}
}
