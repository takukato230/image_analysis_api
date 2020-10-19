# image_analysis_api
画像解析API

## アーキテクチャ
```text
├── Dockerfile.dev
├── README.md
├── adapter
│   └── gateway
│       └── ocr_gateway.go
├── air.conf
├── api
│   ├── server
│   │   ├── app_middleware
│   │   │   └── middleware.go
│   │   ├── handler
│   │   │   ├── analysis_handler.go
│   │   │   └── upload_file_handler.go
│   │   ├── router
│   │   │   └── api_router.go
│   │   └── server.go
│   └── types
│       ├── common_types.go
│       └── file_upload_types.go
├── applications
│   ├── analysis_usecase.go
│   └── upload_file_usecase.go
├── docker-compose.yaml
├── domain
│   ├── model
│   │   └── ocr.go
│   └── service
│       ├── auth_service.go
│       └── file_loading_service.go
├── go.mod
├── go.sum
├── image_analysis_api.iml
├── infrastructure
│   ├── appctx
│   │   └── appctx.go
│   ├── config
│   │   └── config.go
│   └── logger
│       └── logger.go
├── injector
│   └── injector.go
├── main.go
└── pkg
    └── x_request_id_generator.go

```  

### 利用技術  

- go 1.14
- docker
- docker compose
- [tesseract](https://github.com/tesseract-ocr/tesseract)  
## 起動方法
docker-composeを利用して起動する
```shell script
# docker compose build
docker-compose build
# docker compose up
docker-compose up
```