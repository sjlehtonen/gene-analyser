# DNA analyser

A small project to practice Golang and RabbitMQ by doing DNA analysis.

## Usage

Start the client and server by `go run main.go` in their respective src folders. In the client the command `analyse-gene` is used to create analysis tasks that are sent to the server. Server picks up tasks and generates reports for the genes. Generated reports can be found in the reports folder of the server.

Example usage: `analyse-gene ESR1` creates an analysis task for the estrogen receptor gene ESR1.
