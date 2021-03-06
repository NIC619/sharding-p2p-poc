# sharding-p2p-poc

A poc of sharding p2p layer with pubsub in libp2p, based on the idea from the [slide](
https://docs.google.com/presentation/d/11a0jibNz0fyUnsWt9fa2MmghHANdHAAABa0TV7EieHs/edit?usp=sharing).


# Getting started: Docker dev environment playground

This require `docker` and `docker-compose`

## Building the image

```
# This builds a image that with all depending go packages.
make build-dev
```

```
# When done developing, this builds a binary file `main`.
make run-dev
```

Run many instances

```
# This runs a private net with 1 bootstrap node and 5 other nodes.
make run-many-dev
# Stop and remove unused container.
make down-dev
```

## Testing

```
make test-dev
```