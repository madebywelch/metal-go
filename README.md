# Unofficial Metal SDK in Go

This project provides an unofficial Go SDK for Metal, a production ready, fully-managed, ML Storage Platform. Use Metal to find meaning in your unstructured data with embeddings. The SDK makes it easy to interact with the Metal API in Go applications. For more information about Metal, including API documentation, visit the official [Metal documentation.](https://docs.getmetal.io/introduction)

[![GoDoc](https://godoc.org/github.com/madebywelch/metal-go?status.svg)](https://pkg.go.dev/github.com/madebywelch/metal-go)

## API Coverage

As of April 8th 2023, this Metal SDK for Go provides full API coverage for Metal Technologies Inc. All available endpoints for Metal API have been implemented in this SDK.

## Installation

You can install the Metal SDK in Go using go get:

```go
go get github.com/madebywelch/metal-go
```

## Usage

To use the Metal SDK, you'll need to initialize a client and make requests to the Metal API. Here's an example of initializing a client and performing a search:

```go
import "github.com/madebywelch/metal-go/pkg/metal"

func main() {
    client := metal.NewClient(apiKey, clientID)

    searchReq := metal.SearchRequest{
        App:  "your_app_name",
        Text: "your_search_text",
    }

    searchResp, err := client.Search(searchReq)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Search results: %v\n", searchResp.Data)
}
```

## Automatic Retries

This unofficial Metal API client SDK supports automatic retries for requests that fail or return unexpected status codes. By using the `WithMaxRetries` and `WithRetryDelay` options, you can configure the client to automatically retry failed requests up to a specified number of times with a specified delay between attempts. This can help improve the reliability of your API calls, particularly in case of temporary network issues or server-side errors.

**Example:**

```go
client, err := metal.NewClient(
	apiKey,
	clientID,
	metal.WithMaxRetries(5),               // Retry up to 5 times
	metal.WithRetryDelay(3 * time.Second), // 3-second delay between retries
)
```

## Sending Hundreds of Requests Concurrently with WorkerPool

The WorkerPool utility can be used to efficiently send a large number of requests concurrently. This example demonstrates how to use the WorkerPool to send hundreds of Index requests to the Metal API.

```go
func main() {
	client, err := metal.NewClient(apiKey, clientID)
	if err != nil {
		panic(err)
	}

	workerCount := 10 // Number of workers
	requestCount := 100 // Number of requests to send
	wp := utils.NewWorkerPool(workerCount)

	// Add tasks to the worker pool
	for i := 0; i < requestCount; i++ {
		indexReq := metal.IndexRequest{
			App:  "sample_app",
			Text: fmt.Sprintf("Text %d", i),
		}

		wp.AddTask(func() {
			indexResp, err := client.Index(indexReq)
			if err != nil {
				fmt.Printf("Error indexing request %d: %s\n", i, err)
			} else {
				fmt.Printf("Indexed request %d, created document with ID: %s\n", i, indexResp.Data.ID)
			}
		})
	}

	// Wait for all tasks to complete and close the worker pool
	wp.CloseAndWait()
}
```

## Contributing

Contributions to this project are welcome. To contribute, follow these steps:

- Fork this repository
- Create a new branch (`git checkout -b feature/my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push the branch (`git push origin feature/my-new-feature`)
- Create a new pull request

## License

This project is licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details.
