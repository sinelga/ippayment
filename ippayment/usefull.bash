

#GOBIN=./bin GOPATH=$PWD:/home/juno/junoworkspace/gows go install src/taskoutmarkserv/taskoutmarkserv.go

GOBIN=./bin GOPATH=$(pwd) go install src/ippayment.go
GOBIN=./bin GOPATH=$(pwd) go install src/hitshandler.go
GOBIN=./bin GOPATH=$(pwd) go install src/allcollection.go
GOBIN=./bin GOPATH=$(pwd) go install src/showcollection.go
GOBIN=./bin GOPATH=$(pwd) go install src/start.go
GOBIN=./bin GOPATH=$(pwd) go install src/logconverter.go



test


curl -G -H "X-UP-CALLING-LINE-ID: 358451202801" -d "id=111111122" -d "site=test.com" -d "resource=mobilephone" -d "themes=adult" -d "provider=sonera" test.com/sonera?callback=lslslslsl


echo "SET MONIT-TEST value\r\n" | socat - UNIX-CONNECT:/var/run/redis/redis.sock

echo "EXISTS MONIT-TEST" | socat - UNIX-CONNECT:/var/run/redis/redis.sock
