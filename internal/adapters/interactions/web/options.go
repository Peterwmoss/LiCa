package web

type ServeOptions struct {
	Port int
}

type OptionsFunction func(opts *ServeOptions)

func defaultOptions() ServeOptions {
	return ServeOptions{
		Port: 3000,
	}
}

func WithPort(port int) OptionsFunction {
	return func(options *ServeOptions) {
		options.Port = port
	}
}
