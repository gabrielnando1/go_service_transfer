FROM golang:latest

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get github.com/gorilla/mux
RUN go get github.com/go-delve/delve/cmd/dlv
#RUN go get github.com/derekparker/delve/cmd/dlv

# set the working directory to /go/src
WORKDIR $GOPATH/src

# Copy everything from the current directory to the working directory inside the container
# (as set by WORKDIR)
COPY . .

# 8080 is for the web application
EXPOSE 8000

RUN go build -o main main.go
CMD ["./main"]