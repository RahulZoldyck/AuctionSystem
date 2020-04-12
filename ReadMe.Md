#Steps to install

## Setup of Auctioneer
`cd auctioneer`

 `go build`
 
 `docker build --tag auctioneer . `
 
 `docker run --publish 8888:8888 --detach --name auctioneer auctioneer`
 
 ## Setup of Bidder
`cd bidder`

 Input port number, name, delay and value in Dockerfile

 `go build`
 
 `docker build --tag <biddername> . `
 
 `docker run --publish <bidderport>:<bidderport> --detach --name <biddername> <biddername>`
