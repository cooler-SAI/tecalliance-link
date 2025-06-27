package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"tecalliance-link/internal/web"
)

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("does't current support")
	}
	if err != nil {
		log.Printf("WARNING: can not open broser: %v", err)
	}
}

func main() {

	addr := "localhost:8088"
	fullURL := "http://" + addr

	linksHandler := web.LinksHandler()
	http.HandleFunc("/", linksHandler)

	go func() {
		log.Printf("Server started. URL: %s", fullURL)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Critical Server Error: %v", err)
		}
	}()

	log.Printf("Opening browser...")
	openBrowser(fullURL)

	select {}

}
