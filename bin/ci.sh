#!/bin/bash
#部署到跳板服务器
set -e

SSH1="root@catlabs.cn"
DIR1="$SSH1:~/pet-sim/bin"

deploy() {
    echo "build $1..."
    rm -f $1.gz
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./$1 $2
    echo "copy.."
    gzip $1
    scp ./$1.gz $DIR1/
    rm -f $1.gz
    echo done!
}

remote() {
    ssh $SSH1 << EOF
#!/bin/bash
set -e
cd ~/pet-sim/bin
gzip -df server.gz
supervisorctl restart pet
EOF
}

# build one/many
for arg in "$@"
do
	if [[ $arg == 1 ]]; then
	    deploy server "github.com/zii/pet-sim/cmd"
	elif [[ $arg == 2 ]]; then
	    remote
	else
		echo unknown argument: $arg
	fi
done
