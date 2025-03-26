

## simple hello world for event queues SQS

to run
```bash 
sudo docker run --rm -it -p 127.0.0.1:4566:4566 -p 127.0.0.1:4556:4556 -v /var/run/docker.sock:/var/run/docker.sock localstack/localstack
```

```bash
./init-localstack # creates the first queue
```

then 
`go run producer/main.go`

and
`go run consumer/main.go`


next steps:
- add better setup with docker compose or something
- increase scale of messages and prevent ordering/duplicate cases
