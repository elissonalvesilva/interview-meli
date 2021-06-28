# Interview application mercado livre

## Tecnologies
- Golang
- Docker
- Kubernetes
- Istio
- gRPC

## Design Pattern, Architecture
- Factory
- Dependency Injection
- Repository Pattern
- DDD
- Clean Architecture
- Microservices Architecture

## Applications
- api
> responsible for getting dna data and send do analyzer service, and return a stats
- analyzer
> responsible for getting dna and analyze

## Project
### Start Project
#### 1 - Create .env file
 - Create .env file using .env.sample and set the variables there

#### 2 - Step
**If want to run local**
- API
```bash
$ go run ./api/server/cmd/http/main.go
```
- analyzer
```bash
$ go run ./analyzer/cmd/server/main.go
```
> remember you must up a mongodb or pass a external mongodb database in .env

**If want to run in docker**
```bash
$ make build
$ make start
```

### Routes
> Host
 - in local:

`http://localhost:4513/`
 - in cloud Azure AKS
 
`http://52.153.123.208`

> Get Stats 
- GET `/stats`
- Params

> POST Simian
- POST `/simian`
- Params
    - dna: string array
- Example
```json
{
  "dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]
}
```

## Run test
 - Local Test
```bash
$ make test-local
```

## Logs
**If you are using docker**
- Show full logs
```bash
make logs-api
```
```bash
make logs-analyzer
```
- Show tail logs (last 100 lines)
```bash
make logs-tail
```