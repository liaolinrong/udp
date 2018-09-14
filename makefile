all: server client

# TAG = 0.1
# PREFIX = registry.hundsun.com/hcs/udpServer

server: udpServer.go
	CGO_ENABLED=0 go build -a ./udpServer.go

client: udpClient.go
	CGO_ENABLED=0 go build -a ./udpClient.go

# container: server
# 	docker build -t $(PREFIX):$(TAG) .

# push: container
# 	docker push $(PREFIX):$(TAG)

clean:
	rm -f udpServer
	rm -f udpClient
