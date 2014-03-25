

rm pkg/linux_386/*.a
rm bin/*
git pull

GOBIN=./bin GOPATH=$(pwd) go install src/ippayment.go
GOBIN=./bin GOPATH=$(pwd) go install src/hitshandler.go
GOBIN=./bin GOPATH=$(pwd) go install src/allcollection.go
GOBIN=./bin GOPATH=$(pwd) go install src/showcollection.go
GOBIN=./bin GOPATH=$(pwd) go install src/start.go
GOBIN=./bin GOPATH=$(pwd) go install src/logconverter.go
GOBIN=./bin GOPATH=$(pwd) go install src/sendsms.go
GOBIN=./bin GOPATH=$(pwd) go install src/tstip.go
