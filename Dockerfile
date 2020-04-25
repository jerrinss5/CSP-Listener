FROM golang:1.14-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/csplistener

# Now copy it into our base image.
# FROM gcr.io/distroless/base-debian10
FROM scratch
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /go/bin/csplistener /

# Use as an unprivileged user.
USER appuser:appuser

# Run the csplistener binary.
CMD ["/csplistener"]