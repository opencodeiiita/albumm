# ALBUMM

![Logo](./assets/logo.png)

**Albumm is a simple CLI application to download Albums from Flickr.**

### Installation
To install all the dependencies do 

`go mod tidy`

You will need an API Key from Flickr to use this CLI. Find it [here](https://www.flickr.com/services/developer/api/)

Set this API Key as an environment variable using

`export $FLICKR_API_KEY=<YOUR_KEY_HERE>`

To run the application use

`go run *.go`

or you can build a binary using

`go build`

