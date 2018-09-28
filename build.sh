APPPATH=$(cd `dirname $0`; pwd)
export GOPATH=$APPPATH
go build -o $APPPATH/tomtom $APPPATH/main.go
#echo `$APPPATH/tomtom`
