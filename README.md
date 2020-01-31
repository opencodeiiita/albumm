# ALBUMM

![Logo](./assets/logo.png)

**Albumm is a simple CLI application to download Albums from Flickr.**


### VERSION:
   1.0.0



### Installation

To get this project up and running on your local machine, you need to do is clone or download this project

To install all the dependencies do 

`sudo apt update && sudo apt install golang`

`sudo snap install go --classic`



Enter the directory and use

`go mod tidy`




You will need an API Key from Flickr to use this CLI.

Find it [here](https://www.flickr.com/services/developer/api/)



Set this API Key as an environment variable using

`export`
 
`FLICKR_API_KEY=<YOUR_KEY_HERE>`

To run the application use

`go run *.go`

or you can build a binary using

`go build`


### Usage commands

`go run *.go id`  - Get User ID of a username

`go run *.go albums`  - List all albums of a user

`go run *.go help`  - Shows a list of commands or help for one command


### GLOBAL OPTIONS:

   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)




### Contributors

See also the list of [contributors](<link to be updated>) who participated in the project.



