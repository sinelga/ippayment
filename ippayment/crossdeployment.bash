

rm pkg/linux_386/*.a
rm bin/ippayment
rm bin/hitshandler
rm bin/allcollection
rm bin/showcollection
rm bin/start
rm bin/logconverter
rm bin/sendsms
rm bin/tstip

git pull

GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/ippayment.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/hitshandler.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/allcollection.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/showcollection.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/start.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/logconverter.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/sendsms.go
GOBIN=./bin GOPATH=$(pwd):/home/juno/workspace/gocode go install src/tstip.go


