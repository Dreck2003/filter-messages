
############################
# STEP 1 WEB ... Missing -_(°v°) _-
############################
# Install node


############################
# STEP 2 Golang ... Missing -_(°v°) _-
############################
FROM golang:1.19-alpine

WORKDIR /app

COPY ./indexer/ ./

# RUN go mod download 

# Build the binary
RUN go build -o indexer

# RUN go run .
EXPOSE 4002

# Run the binary
CMD ["./indexer"]