## CRYPTO CONVERSION API

### **Setup**

#### Database Setup
First create database in postgres and update the `config.yml` file with db configurations

#### API Setup
Run the following command to download the packages

`go mod tidy`

**NOTE** : Make sure the `go modules` are enabled

If this doesn't workout then use the following command

`go get ./...`


Use either of these commands to run api

`go run app/main.go`

or 

`make run`

Defualt port is 5000 you can access the API at this url 

`localhost:5000/service/api?fsyms=BTC,LINK,MKR&tsyms=USD,EUR,ETH,LTC`

