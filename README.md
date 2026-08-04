# LoadBalancer
This project integrates many fundamental concepts in networking and concurrent programming
into a single application. Building a production-style load balancer requires handling multiple
client connections, managing shared resources safely, detecting server failures, and maintaining
high availability under changing conditions.
This project focuses on following core areas of systems programming:
- Socket programming — accepting client connections, communicating with backend
servers, and forwarding network traffic.
- Concurrency — handling thousands of simultaneous connections using Go’s gorou-
tines while safely managing shared resources such as backend health, configuration, and
connection counts.
- Fault tolerance — automatically detecting failed backend servers, removing them from
service, and adding them back when they recover.
- Protocol awareness — understanding the differences between Layer 4 (TCP) and Layer
7 (HTTP) load balancing, and the routing decisions that are possible at each layer.
- System operations— supporting graceful shutdown and health monitoring so the system
can be operated reliably.

