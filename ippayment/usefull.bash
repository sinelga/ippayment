

#GOBIN=./bin GOPATH=$PWD:/home/juno/junoworkspace/gows go install src/taskoutmarkserv/taskoutmarkserv.go

GOBIN=./bin GOPATH=$(pwd) go install src/ippayment.go
GOBIN=./bin GOPATH=$(pwd) go install src/hitshandler.go
GOBIN=./bin GOPATH=$(pwd) go install src/allcollection.go
GOBIN=./bin GOPATH=$(pwd) go install src/showcollection.go
GOBIN=./bin GOPATH=$(pwd) go install src/start.go
GOBIN=./bin GOPATH=$(pwd) go install src/logconverter.go
GOBIN=./bin GOPATH=$(pwd) go install src/sendsms.go
GOBIN=./bin GOPATH=$(pwd) go install src/tstip.go



test


curl -G -H "X-UP-CALLING-LINE-ID: 358451202801" -d "id=111111122" -d "site=test.com" -d "resource=mobilephone" -d "themes=adult" -d "provider=sonera" 192.168.1.100/scanip?callback=lslslslsl


echo "SET MONIT-TEST value\r\n" | socat - UNIX-CONNECT:/var/run/redis/redis.sock

echo "EXISTS MONIT-TEST" | socat - UNIX-CONNECT:/var/run/redis/redis.sock


for i in `seq 1 10`; do sleep 300; bin/hitshandler; bin/sendsms; done



curl -i -d "MSISDN=%2B358451202801&ServiceId=2&Price=100" ippayment.info:8080/MobileCharger?Action=Check

curl -i -d "MSISDN=%2B358451202801&ServiceId=2&Price=100&TransactionId=" ippayment.info:8080/MobileCharger?Action=Charge


GOPATH=$GOPATH:/home/juno/git/ippayment/ippayment go test -v
