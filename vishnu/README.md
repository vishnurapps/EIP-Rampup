# EIP-Rampup

make a folder inside the src folder

```shell
mkdir demo
```

now change to that directory and initialize go module

```shell
cd demo
go mod init vishnu
```

create our go application

## Dockerfile creation

- Make sure to use the golang image instead of linux images as we need to do many configurations in the later case.
- set the working directory `WORKDIR /app`
- Copy everything from the current directory to the Working Directory inside the container `COPY . .`
- Build the Go app `RUN go build -o main .` Dependencies will be downloaded at this stage
- Expose port 9091 to the outside world `EXPOSE 9091`
- Command to run the executable `CMD ["./main"]`
