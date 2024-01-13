# API

## Install

Requires `go >= 1.21`, `air` ([cosmtrek/air](https://github.com/cosmtrek/air)) and `docker` - and a user added to the `docker` group for the `makefile` to work as intended

## Run

```bash
make run # (or just: make)
```

The password for the local postgres container is running on port `5433`, with password: `postgres` and database: `lica`.

To stop the postgres container simply do: `make postgres-stop`.

And to remove the container completely, do: `make postgres-rm`.
