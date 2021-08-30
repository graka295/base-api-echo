# apgateway

apgateway is a gateway for accessing all service

## Prequesites

1. Golang 1.14.x

## Usage
-> Create database mysql on your system, name is free

1. Clone app from repo
2. Navigate to project folder
4. Execute go mod vendor
5. Adjust your config on conf/config.json
6. Migrate your table using 'make migrate-db-up' in your terminal
7. Create swagger using 'make swagger-doc' in your terminal