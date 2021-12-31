# Go Client SDK for the CyberSource REST API

[![Documentation](https://godoc.org/github.com/tooolbox/cybersource-rest-client-go?status.svg)](https://godoc.org/github.com/tooolbox/cybersource-rest-client-go)

The CyberSource Go client provides convenient access to the [CyberSource REST API](https://developer.cybersource.com/api/reference/api-reference.html) from your Go application.

## Requirements  
* Go version 1.13 or higher
* A CyberSource account (see _Registration & Configuration_ section below)


## Installation

Go Modules:
```shell
# In your project:
go get github.com/fauziasbrankas/cybersource-rest-client
```


## Registration & Configuration
Use of this SDK and the CyberSource APIs requires having an account. You can find details of getting a test account and creating your keys [here](https://developer.cybersource.com/api/developer-guides/dita-gettingstarted/registration.html)

Remember this SDK is for use in server-side Go applications that access the CyberSource REST API and credentials should always be securely stored and accessed appropriately. 


## SDK Usage Examples and Sample Code
You can find details and examples of how the API is structured in the API Reference Guide:
* [Developer Center API Reference](https://developer.cybersource.com/api/reference/api-reference.html)

The API Reference Guide provides examples of what information is needed for a particular request and how that information would be formatted. Using those examples, you can easily determine what methods would be necessary to include that information in a request using this SDK.


### Switching between the sandbox environment and the production environment
Cybersource maintains a complete sandbox environment for testing and development purposes. This sandbox environment is an exact duplicate of the production environment with the transaction authorization and settlement process simulated. By default, this SDK is configured to communicate with the sandbox environment.

API credentials are different for each environment, so be sure to switch to the appropriate credentials when switching environments.



## License
This repository is distributed under a proprietary license. See the provided [`LICENSE.txt`](/LICENSE.txt) file.
