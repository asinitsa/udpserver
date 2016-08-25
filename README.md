# UDP socket listener 

To build:
```
go build
```
To run tests
```
go test 
```
To run executable
```
./udpserver 
```
To use 
```
$ echo -n '[17/06/2016 12:30] Time to move' | nc -u localhost 1234
# Server Output:
{"timestamp":1466166600,"message":"Time to move"}
```

# Docker image 
To build docker image 
```
go build 
docker build -t udpserver .
```

To run docker image 
```
docker run -p 1234:1234/udp udpserver 
```

To use docker container 
```
$ echo -n '[17/06/2016 12:30] Time to move' | nc -u localhost 1234
# Server Output:
{"timestamp":1466166600,"message":"Time to move"}
```
