

#GOBIN=./bin GOPATH=$PWD:/home/juno/junoworkspace/gows go install src/taskoutmarkserv/taskoutmarkserv.go

GOBIN=./bin GOPATH=$(pwd) go install src/ippayment.go


test


curl -G -H "X-UP-CALLING-LINE-ID: 358451202801" -d "id=111111122" -d "site=test.com" -d "resource=mobilephone" -d "themes=adult"  test.com/sonera?callback=lslslslsl


echo "SET MONIT-TEST value\r\n" | socat - UNIX-CONNECT:/var/run/redis/redis.sock

echo "EXISTS MONIT-TEST" | socat - UNIX-CONNECT:/var/run/redis/redis.sock
