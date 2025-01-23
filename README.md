# Qrat

Qrat is a simple Remote Administration Tool (RAT) written in Go.

It lives as HTTP server and respond to JSON. Default port is 226.


### Example

Build the program, deploy to "Server" and run (I assume Windows)
```
go build -ldflags="-H windowsgui"
```


**On the Server**

Run the binary


**On Client**

```
curl -X POST localhost:226/c -d "{\"command\":\"ipconfig\"}"

// if you need to add arguments, use comma to separate
curl -X POST localhost:226/c -d "{\"command\":\"ipconfig,/all\"}"
```