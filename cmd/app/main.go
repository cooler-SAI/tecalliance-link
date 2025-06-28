package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"tecalliance-link/internal/web" // Ensure this package path is correct
)

// openBrowser opens the specified URL in the default browser for the current OS.
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		// For Linux, use xdg-open.
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		// For Windows, use "cmd /c start", this is a more reliable way to open a URL.
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		// For macOS, use the open command.
		err = exec.Command("open", url).Start()
	default:
		// If the OS is not supported, print an error.
		err = fmt.Errorf("Current OS does not support browser opening: %s", runtime.GOOS)
	}
	if err != nil {
		// Print a warning if the browser could not be opened.
		log.Printf("WARNING: Could not open browser: %v", err)
	}
}

func main() {
	addr := "localhost:8095"    // Address on which the server will listen
	fullURL := "http://" + addr // Full URL to open in the browser

	// --- Start of new logic to handle already running server ---
	// Attempt to temporarily create a listener on the port to check if it's already in use.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// If an error occurred, check if it's an "address already in use" error.
		if opErr, ok := err.(*net.OpError); ok && opErr.Op == "listen" && opErr.Err.Error() == "address already in use" {
			// If the port is busy, it means the server is already running.
			log.Printf("Server is already running on %s. Opening browser to existing instance.", fullURL)
			openBrowser(fullURL) // Open the browser to the already running server.
			os.Exit(0)           // Exit the current process, as another one is already running.
		}
		// If it's another error when trying to listen on the port, print it and exit the program.
		log.Fatalf("Error trying to listen on %s: %v", addr, err)
	}
	// If the port is free, immediately close the temporary listener,
	// as the real listener will be created by http.ListenAndServe.
	err2 := listener.Close()
	if err2 != nil {
		log.Fatalf("Error closing temporary listener: %v", err2)
		return
	}
	// --- End of new logic ---

	// If we reached this point, the port is free, and we can start the server.
	log.Printf("Starting server on %s", fullURL)

	// Create a new HTTP server
	server := &http.Server{
		Addr: addr,
	}

	// Configure the links handler from the web package.
	linksHandler := web.LinksHandler()
	http.HandleFunc("/", linksHandler) // Use http.DefaultServeMux

	// --- New: Handler for favicon.ico ---
	// This handler serves the favicon.ico file from the project root.
	// Ensure your favicon.ico file is located at the project root, e.g., D:\Enterprise Development\Go-projects\tecalliance-link\favicon.ico
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// Serve the favicon.ico file. The path "favicon.ico" is relative to the Working Directory,
		// which is configured as the project root.
		http.ServeFile(w, r, "favicon.ico")
	})
	// --- End New ---

	// Start the server in a separate goroutine so the main function can continue execution.
	go func() {
		// Here we use log.Println instead of log.Fatalf to avoid immediate program exit
		// for errors not related to an already busy port.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Critical Server Error: %v", err)
		}
	}()

	// Print message and open the browser.
	log.Printf("Opening browser...")
	openBrowser(fullURL)

	// Block the main goroutine indefinitely to keep the server running.
	// This prevents the program from exiting while the server is active.
	select {}
}
