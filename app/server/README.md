# Server Package

This package is being used to store Server object constructor,
where the server object should have its dependencies.

Also, each request handler whether it's an gRPC request or REST request.
It should only call Use Case Package and Proto Package.

Each file other than `server.go` should represent services or resources of the application resources.
