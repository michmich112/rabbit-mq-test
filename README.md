# rabbit-mq-test

Test for creating a chat using rabbit-mq using a sender (provider) and a receiver (consumer).

## Getting Started
Create your RabbitMQ instance using docker: 
```
docker pull rabbitmq:latest
docker run -d --hostname rabbit-mq-test -p 5672:5672 --name rabbit rabbitmq:latest
```
Then navigate to the respository on 2 terminal sessions. On the first one, launch the receiver by running:
```
./chat receive
```
and on the second one start the sender by running:
```
./chat send
```

On the sender window, write any message and press enter to send it. It should appear automatically on the receiver window!
