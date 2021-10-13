# Notification Service and Notification Gateway

# How to run

```
make build
docker-compose up -d
```

# How to stop

```
docker-compose down
```

# Send a request to the service

```
curl -X POST http://localhost:8080/notifications -H "Content-Type: application/json" -d @test/notification.json
```