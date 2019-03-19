#### Run
`go run src/app/main.go`

#### Hello world
`curl http://localhost:1323`

#### Validation
`curl -X POST http://localhost:1323/users -H 'Content-Type: application/json' -d '{"name":"Joe","email":"joe@valid-domain.com"}'`

#### Custom validation annotation
`curl -X POST http://localhost:1323/cars -H 'Content-Type: application/json' -d '{"brand":"invalid","miles":12}'`
