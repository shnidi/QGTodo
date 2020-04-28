FROM golang:1.14.2


# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
#RUN go get /go/src/github.com/saegewerk/QGTodo/...
RUN go get github.com/saegewerk/QGTodo/...

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/todoServer

# Document that the service listens on port 8080.
EXPOSE 8000