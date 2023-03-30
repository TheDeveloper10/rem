package rem

type Settings struct {
	DefaultHandler Handler
	BodyProcessor  BodyProcessor
}

var settings *Settings = nil

func SetSettings(s *Settings) {
	settings = s
}

func DefaultSettings() {
	settings = &Settings{
		DefaultHandler: func(ctx *Context) Response { return Success(nil) },
		BodyProcessor:  &JSONBodyProcessor{},
	}
}
