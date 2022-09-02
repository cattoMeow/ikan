FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

# copy 1 file saja
# COPY *.go ./

# copy 1 folder
COPY . ./

RUN go build -o /main

CMD ["/main"]