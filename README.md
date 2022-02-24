# Bookstore Service

This repository contains basic implementation of GRPC-Gateway on Golang.

# Dependency

This repository requires this following softwares to work:
- PostgreSQL
- Go
- Docker
- Cmake

That's it, but if somehow `generate` command is not working, do this instead:

```sh
buf generate
```

# How to run?

To run this repository all you need to do is run the make script:

```sh
make setup
```

It will setup the repository, and then change the Makefile environment variable, and run:

```sh
make develop
```
