# Unofficial Metal SDK in Go

This project provides an unofficial Go SDK for Metal, a production ready, fully-managed, ML Storage Platform. Use Metal to find meaning in your unstructured data with embeddings. The SDK makes it easy to interact with the Metal API in Go applications. For more information about Metal, including API documentation, visit the official [Metal documentation.](https://docs.getmetal.io/introduction)

[![GoDoc](https://godoc.org/github.com/madebywelch/metal-go?status.svg)](https://pkg.go.dev/github.com/madebywelch/metal-go)

## Installation

You can install the Metal SDK in Go using go get:

```go
go get github.com/madebywelch/metal-go
```

## Usage

To use the Metal SDK, you'll need to initialize a client and make requests to the Metal API. Here's an example of initializing a client and performing a search:

```go
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
