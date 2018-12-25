ps aux | grep main_ireader | grep -v grep | awk '{print $2}' | xargs kill -9
export GOPATH=$GOPATH:/data/IReaderServer
rm ./bin/main_ireader
git pull origin dev_1.0
export GIN_MODE=release
go install main_ireader
nohup ./bin /main_ireader  > ret.log &