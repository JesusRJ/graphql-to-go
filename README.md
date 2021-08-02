# Graphql to Go

Extract Graphql schemas and generate Go types using introspection.

# Use

Extract all entities:

```shell
go run main.go https://API_URL "API_TOKEN"
```

Extract one entity:

```shell
go run main.go https://API_URL "API_TOKEN"
```

# Sample

```shell
$ go run main.go https://countries.trevorblades.com/ ""

or

$ go run main.go https://countries.trevorblades.com/ "" Country
```
