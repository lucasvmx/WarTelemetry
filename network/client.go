package network

import (
	"log"
	"net/http"

	"github.com/lucas-engen/WarTelemetry/model"
)

func RetrieveData() {
	response, err := http.Get(model.URL)
	if err != nil {
		log.Fatalf("[FATAL] failed to send request because of: %v", err)
	}


}
