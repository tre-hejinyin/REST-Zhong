# zhongdalu-test-api

### install command

> make install

### start command

- local run

> make air

- docker run

> make up

## 目录构成

```shell
.
├──controller   
├── db
│  └── migrations // database migration
├── domain
├── infra 
│  ├── postgres.go 
│  └── router.go 
├── middleware   
│  ├── auth      
│  ├── cache
│  ├── constants
│  ├── jwt
│  ├── response
│  ├── validator
│  ├── errorhandler 
│  │  ├── api_error.go  
│  │  ├── error_bussiness.go
│  │  ├── error_code.go 
│  │  ├── error_handler.go 
│  │  └── error_message.go 
│  └── logger     
│     ├── error_handler.go 
│     └── error_message.go 
├── repository 
├── testdata 
├── usecase   
├── docker-compose.yml 
├── Dockerfile
├── go.mod 
├── go.sum
├── main.go    // entry
├── Makefile   
└── README.md

```


###Tips
The functions of the test API are in the 'testdata' folder