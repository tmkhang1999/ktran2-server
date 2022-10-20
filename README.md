# CSC 482

<h3> Create a Go HTTP server that listens for connections and responds to requests. </h3>

The server should run in a Docker container on EC2 and listen on a port of your choosing. It should respond as follows:

- Listen to the /<your_laker_id_name>/status endpoint and respond to GET requests with a JSON body containing the system time and an HTTP status of 200.
- All requests to other endpoints should result in an HTTP status of 404.
- All requests using a method other than GET should return an HTTP status of 405.
- A record of each request including the method type, source IP address, request path, and the resulting HTTP status code should be sent to Loggly.