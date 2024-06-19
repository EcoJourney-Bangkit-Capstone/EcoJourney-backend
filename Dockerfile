FROM golang:latest
WORKDIR /app
COPY . .
ENV HOST 0.0.0.0
RUN go mod download
RUN go build -o main ./server
EXPOSE 8080
CMD [ "./main" ]

# FROM golang:latest
# WORKDIR /app
# COPY . .

# ENV HOST 0.0.0.0

# # Environment variables which CompileDaemon requires to run
# ENV PROJECT_DIR=/app \
#     GO111MODULE=on \
#     CGO_ENABLED=0

# RUN go mod download
# # Get CompileDaemon
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go install github.com/githubnemo/CompileDaemon

# EXPOSE 8080
# ENTRYPOINT CompileDaemon -build="go build -o main ./server" -command="./main"