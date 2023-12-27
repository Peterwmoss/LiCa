# API

## Install

Requires `go >= 1.21` and `docker` - and a user added to the `docker` group for the `makefile` to work as intended

```bash
make get
```

## Run

```bash
make run # (or just: make)
```

To stop the postgres container simply do: `make postgres-stop`

And to remove the container completely, do: `make postgres-rm`
