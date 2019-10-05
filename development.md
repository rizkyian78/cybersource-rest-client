# Development

### Specification
The swagger spec provided by CyberSource has several fields with invalid parameters:
- String fields with `minimum` validators (a numeric property)
- Boolean fields with `maxLength` validators (a string property)

Although the spec validates, the generated code does not compile due to mismatched types.  To fix this, execute `go run fix-spec.go` and a modified spec will be generated.

### Generation
First, install [go-swagger](https://github.com/go-swagger/go-swagger):
```
brew tap go-swagger/go-swagger
brew install go-swagger
```

Then execute `./generator/cybersource_go_sdk_gen.sh` to generate the client & models.
