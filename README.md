#Gin Webserver - a server control addon
Gin does not natively provide a way to close a server once started. This
addon library with a single file provides a WebServer struct which can
be initialized to a server object:

    engine := gin.Default()
    host := "127.0.0.1:80"
    // Attach routes to engine
    server := network.InitializeWebServer(engine, host)

Once the server has been initialized it can has two methods, start and
stop to control the gin web server.

After initialization the server object must be started.

    server.Start()

If the server needs to be shut down use the stop method.

    server.Stop()

A very simple example is provided to illustrate how it can be used.

### Example

To use the example just launch the example.go file from the example
folder.

    cd example
    go run example.go
