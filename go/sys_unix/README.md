# TCP server

This is a toy project to learn more about how TCP connections work. Notices that I'm using

```go
golang.org/x/sys/unix
``` 

This package is used by high-level packages such as "os". The recommendation is not 
to use this package.

**Server**

- Create socket
- Bind to a port
- Listen 
- Accept connections
- Read/Write

**Client**

- Create socket
- Connect 
- Read/Write 

### Implementation A

The first implementation was quite straightforward, however it lacked of managing multiple requests.

```go

main(){ 
// more code here..

 for {
 	newFd, _, err := unix.Accept(fd)
	for {
		// read from newFd
	}
 }
}

```
The problem with this approach is that only one client may be connected at a time, with the inner ```for``` loop, we avoid to read from each newFd.

### Implementation B

Now lets provide a way to process multiple requests. If there was a way to process request concunrrently... wait a moment, 
let's use a gorutine.
```go

main(){ 
// more code here..

 for {
 	newFd, _, err := unix.Accept(fd)
	go func() {
		for {
			// read from newFd
		}
	}()	
  }
}
```

