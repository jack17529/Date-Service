# Basic-Date-Service

Made a basic date service with go-kit-CLI using gorilla mux router to get the current date, status of the service and validate a date.

## Running

1. Download the project in your GOPATH.
2. cd inside the folder and run `docker-compose up` using the `docker-compose.yml` file or using `go run hello/cmd/main.go` 

## Architechture

hello/  
|---cmd/  
|------service/  
|----------server.go          Wire the service.  
|----------server_gen.go      Also wire the service.  
|------main.go                Runs the service  
|---pkg/  
|------endpoints/  
|----------endpoint.go        The endpoint logic along with structures for request and reponse.  
|----------endpoint_gen.go    This will wire the endpoints.  
|----------middleware.go      Endpoint middleware  
|------http/  
|----------handler.go         Transport logic encode/decode data and gorilla mux request reponse routing of the service.  
|----------handler_gen.go     This will wire the transport.  
|------io/  
|----------io.go              The input output structs.  
|------service/  
|----------middleware.go      The service middleware.  
|----------service.go         Business logic.  

## Reproducing the service

1. Use `kit new service hello` to make a service named hello.
2. Edit the business logic file (`hello/pkg/service/service.go`) with the methods in your service.
3. To generate the service with gorilla and default middleware use `kit g s hello -w --gorilla`
4. A basic service that does not handle request and reponse can be run right now using `go run hello/cmd/main.go`
5. NOTE: Our service runs at port 8081 and not ont he default gorilla mux port 8080.
6. Now implement our services in `hello/pkg/service/service.go` file.
7. One can write the test in the same service folder for the business logic, right now.
8. Then edit `hello/pkg/endpoint/endpoint.go` for the request reponse structs.
9. Then edit `hello/pkg/http/handler.go` for the encode and decode.
10. Edit the requests routed by the gorilla mux router in `hello/pkg/http/handler.go`
11. Correct any errors if they arise.
12. Make a dockerfile using `kit g d`
13. We are done just run the service using `docker-compose up`

## References

1. https://medium.com/@kujtimii.h/creating-a-todo-app-using-gokit-cli-20f066a58e1
2. https://github.com/GrantZheng/kit
3. https://dev.to/napolux/how-to-write-a-microservice-in-go-with-go-kit-a66
4. https://medium.com/@shijuvar/go-microservices-with-go-kit-introduction-43a757398183
