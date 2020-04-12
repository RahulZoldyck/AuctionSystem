FROM golang

# Copy the local package files to the container's workspace.
ADD bidder /go/bin/bidder

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/bidder


ENV BIDDER_NAME Rahul
ENV BIDDER_DELAY 40
ENV BIDDER_PORT 8081
ENV BIDDER_VALUE 400

# Document that the service listens on port 8080.
EXPOSE 8081