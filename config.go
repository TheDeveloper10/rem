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

func ParseBody(ctx *Context, out any) error {
	return config.BodyProcessor.Parse(ctx.Request().Body(), &out)
}

func Serialize(obj any) ([]byte, error) {
	return config.BodyProcessor.Serialize(obj)
}
