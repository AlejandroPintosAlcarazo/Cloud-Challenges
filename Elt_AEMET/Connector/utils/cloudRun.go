package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HandleCloudRun configura el puerto y lanza el servidor HTTP
func HandleCloudRun() {
	port := getPort()

	if os.Getenv("ENV") == "cloud" {
		go func() {
			log.Printf("Listening on port %s", port)
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Service is running")
			})
			log.Fatal(http.ListenAndServe(":"+port, nil))
		}()
	}
}

// getPort maneja la configuración del puerto basado en la variable de entorno
func getPort() string {
	env := os.Getenv("ENV")
	if env == "cloud" {
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("PORT environment variable not set in cloud environment")
		}
		return port
	}
	return "8080" // Puerto predeterminado para entornos locales
}
