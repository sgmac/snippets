# RPC simple example

After reading the excellent write-up by Kelsey Hightower in Login Usenix magazine, Spring 2016. I decied to make this similar 
example for running a "remote cat".

```
$ cd server && go build -o servercat
$ cd client  && go build -o rcat
```

Then run both of them separatelly a

```
$ ./rcat  -f /etc/resolv.conf                                       
RPC dialing..                                                                             
RPC call...                                                                               
#                                                                                         
# Mac OS X Notice                                                                         
#                                                                                         
# This file is not used by the host name and address resolution                           
# or the DNS query routing mechanisms used by most processes on                           
# this Mac OS X system.                                                                   
#                                                                                         
# This file is automatically generated.                                                   
#                                                                                         
nameserver 10.2.9.2                                                                       
nameserver 10.3.9.2
```
