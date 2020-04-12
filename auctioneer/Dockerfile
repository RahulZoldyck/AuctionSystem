FROM golang

# Copy the local package files to the container's workspace.
ADD auctioneer /go/bin

RUN ./auctioneer

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/auctioneer

# Document that the service listens on port 8080.
EXPOSE 8888