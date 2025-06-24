# Go HTTP Server with Worker Pool

This Go program demonstrates an HTTP server using the **worker pool pattern**. 
It handles HTTP requests concurrently, managing multiple connections efficiently 
through channels and a fixed number of worker goroutines.

## Key Concepts

- **Worker Pool Pattern**: A set of pre-created goroutines (workers) handles 
  tasks from a queue, ensuring the server processes connections without overloading.
  
- **Concurrency Control**: The server uses a fixed number of worker goroutines 
  to process incoming requests. Workers are either idle, waiting for a task, 
  or actively handling requests.

- **Channels for Communication**: The program uses channels to pass connections 
  from the main goroutine to workers, ensuring smooth communication between 
  the two.

## Program Flow

1. **Main Goroutine**:
   - Listens on TCP port `8080` for incoming connections.
   - If workers are available, the connection is passed to an idle worker via 
     the channel (`incomingConnections`).
   - If no workers are available, the server responds with a `429 Too Many 
     Requests` error.

2. **Worker Goroutines**:
   - A fixed number of workers (3 in this case) are created at the start.
   - Workers listen for tasks on the `incomingConnections` channel, process 
     the requests, and send responses back to clients.

3. **Concurrency Management**:
   - The workers handle requests concurrently, but the number of workers is 
     limited to avoid overloading the system.
   - If all workers are busy, the server returns a `429 Too Many Requests` 
     error to new connections.

4. **Channel-Based Communication**:
   - Channels act as a task queue, distributing incoming connections to idle 
     workers.
   - This ensures that only a limited number of requests are processed at any 
     given time.

## Triggering "Server Busy" Error

Due to the small worker pool size (3 goroutines), it's easy to hit the "server 
is busy" scenario. If all workers are busy, new connections are rejected with 
an HTTP `429 Too Many Requests` error.

You can simulate this by opening many simultaneous connections to the server 
using the following command:

```bash
$ seq 1 2000 | xargs -Iname -P100 curl -s "http://localhost:8080/index.xhtml" | 
  grep Busy
