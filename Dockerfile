FROM golang:1.16

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download & Install all the dependencies
RUN go mod download
RUN go install -v

# Create an executable
RUN go build -o main .

# Run the executable
CMD ["./main"]