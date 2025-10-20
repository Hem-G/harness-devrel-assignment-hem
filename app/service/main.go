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

func main() {
	// This function runs first when the program starts

	// Here, we tell the web server what to do when someone visits "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// "w" is how we send information back to the person visiting
		// We send back a friendly text message
		fmt.Fprintf(w, "hello from myservice\n")
	})

	// Another rule: if someone visits "/health", respond with "ok"
	// This is how monitoring systems check if the service is healthy
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	})

	// This rule adds another simple webpage called "/version"
	//
	// When someone visits "/version" (for example: http://localhost:8080/version),
	// the service will reply with the text "version 0.1.0".
	//
	// This is extremely useful in deployments — it helps us easily confirm
	// which version of the application is running in each environment
	// (for example, Dev, QA, or Production).
	//
	// Later, when you update your service and push a new Docker image
	// (say version 0.2.0), you can visit "/version" in the browser or
	// use `curl` to see which version got deployed.
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "version 0.1.0\n")
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
