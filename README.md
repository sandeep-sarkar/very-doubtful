# Statistics Processing
This project is to process expense report per company id. It will also reject columns that has too many variations. This service listens on port 50061. It accepts
following parameters:
- Document
- Columns include
- Columns exclude
- Max variation
It returns the file name that gets created after processing, which can be downloaded from port 50071. For example :

```
curl -GET localhost:50071/result-0510202112325200000.csv
```

# Build
To build this project following prerequisites need to be met:

- go 1.10 and above
- gRPC installed from here https://grpc.io/docs/languages/go/quickstart/
- docker and kubernetes (optional)

## Build process

Simply running below command will start the service:

```
go run main.go
```

To build an image : 
```
go build -o statisticsProcessing main.go

docker build -t docker.io/sansark1/stat-processing:v1.1 -f Dockerfile .
```
This image can be run in the following way:
```
docker run -ti -p 50061:50061 -p 50071:50071 --name stat-processing docker.io/sansark1/stat-processing:v1.1
```

To run this in kubernetes environment `manifest/api.yaml` can be built in the following way :
```
kubectl create -f api.yaml
```
This creates a service of type `loadbalancer` and deployment with two replicasets.
