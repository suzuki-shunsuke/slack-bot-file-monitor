set -e

source `dirname $0`/env.sh
source `dirname $0`/decho.sh

if [ "$ENV" = "docker" ]; then
  which dep > /dev/null 2>&1 || decho go get -u github.com/golang/dep/cmd/dep
fi

decho dep ensure
dexec go run main.go
# echo "+ go run main.go"
# exec go run main.go
