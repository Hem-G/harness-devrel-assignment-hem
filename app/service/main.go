// ---------------------------------------------
// 🧠 What this program does
// ---------------------------------------------
//
// This is a tiny web service written in Go.
//
// This small service is ideal for demonstrating CI/CD pipelines — it’s lightweight, always runs instantly, and lets you see deployments succeed in different environments.
//
// When it runs, it starts a simple web server
// that listens for visitors (requests) on port 8080.
//
// If someone opens the service in a web browser or sends a request:
//
//   • Going to “/” → shows the message: “hello from myservice”
//   • Going to “/health” → shows “ok” (used by systems like Kubernetes
//     to check that the service is healthy and running)
//
// It also prints a message on startup ("myservice started on :8080")
// so we can see in the logs that the server is running.
//
// If anything goes wrong (like if port 8080 is blocked),
// the service prints an error before stopping.
//
// This kind of service is perfect for testing pipelines,
// deployments, and environments — simple, stable, and predictable.
//

package main // Every Go program starts with a "main" package (entry point)

import (
	"fmt"      // A built-in Go library to print text or format output
	"net/http" // A built-in Go library for creating web servers
)

func main() { // This function runs first when the program starts

	// This rule adds a simple webpage called "/"
	// When someone visits "/" (for example: http://localhost:8080/),
	// the service will reply with the text "hello from myservice".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// "w" is how we send information back to the person visiting
		// We send back a friendly text message
		fmt.Fprintf(w, "hello from myservice\n")
	})

	// This rule adds another simple webpage called "/health"
	// When someone visits "/health" (for example: http://localhost:8080/health),
	// This is how monitoring systems check if the service is healthy
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	})

	// Print this line to the console so we know the service started
	fmt.Println("🚀 myservice started on :8080")

	// Actually start the web server on port 8080
	// If anything goes wrong (e.g., port is busy), capture and print the error
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server failed:", err)
	}
}
