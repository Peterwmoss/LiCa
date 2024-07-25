package web

type ServeOptions struct {
	port int
}

type OptionsFunction func(opts *ServeOptions)

func defaultOptions() ServeOptions {
	return ServeOptions{
		port: 3000,
	}
}

func WithPort(port int) OptionsFunction {
	return func(options *ServeOptions) {
		options.port = port
	}
}
