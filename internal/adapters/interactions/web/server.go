package web

import (
	"fmt"
	"net/http"
)

func Serve(router Router, optionsFunctions ...OptionsFunction) error {
	options := defaultOptions()
	for _, optionsFunction := range optionsFunctions {
		optionsFunction(&options)
	}

	server := http.NewServeMux()

	router.SetupRoutes(server)

	return http.ListenAndServe(fmt.Sprintf(":%d", options.port), server)
}
