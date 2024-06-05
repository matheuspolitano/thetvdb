## Wrapping up an API service as a client in Go



<!-- ```
/
├── go.mod
├── go.sum
├── README.md
├── client.go
├── client_test.go
├── resources/
│   ├── resource.go
│   ├── resource_test.go
│   └── resource_types.go
├── internal/
│   └── httpclient/
│       ├── httpclient.go
│       └── httpclient_test.go
└── utils/
    ├── errors.go
    └── errors_test.go

``` -->

The project contains client who is responsable for the receive the api-keym, then we have the resources where is mapped different process and services. And we have the internal comunication with