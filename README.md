# Simple Broadcast Server (CLI APP)
this project is a simple soacket broadcast app that is built with Golang with the net Library, in this project there will be two command one for starting the server and another for client to connect to the server and send messages in real time ,this project is built for roadmap.sh on 
[this project link](https://roadmap.sh/projects/broadcast-server)

# How To Run
first you will have to run this command to compile
> go build -o ./bin/broadcast-server .

then to run the server , enter:
```
 #go to bin
 cd ./bin
 ./broadcast-server start
 ```

the same thing for the client, enter:
> ./broadcast-server connect <URL> -p <PORT>


