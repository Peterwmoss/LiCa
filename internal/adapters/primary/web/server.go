package web

import (
	"fmt"
	"net/http"
)

func Serve(optionsFunctions ...OptionsFunction) error {
	options := defaultOptions()
	for _, optionsFunction := range optionsFunctions {
		optionsFunction(&options)
	}

	server := http.NewServeMux()

	setupRoutes(server)

	return http.ListenAndServe(fmt.Sprintf(":%d", options.port), server)
}
