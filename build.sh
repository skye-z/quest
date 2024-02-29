#!/usr/bin/env bash

echo "Start packaging..."

go mod download
go mod tidy

mkdir ./out
cp LICENSE ./out/LICENSE

generate(){
    CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o quest -ldflags '-s -w'
    mv quest ./out/quest-$2-$1
    cd out
    tar -zcf quest-$2-$1.tar.gz quest-$2-$1 LICENSE
    rm -rf ./quest-$2-$1
    cd ../
}

generate_exe(){
    CGO_ENABLED=0 GOOS=windows GOARCH=$1 go build -o quest -ldflags '-s -w'
    mv quest ./out/quest-$1-windows.exe
    cd out
    zip -q quest-$1-windows.zip quest-$1-windows.exe LICENSE
    rm -rf ./quest-$1-windows.exe
    cd ../
}

# echo "[1] MacOS from amd64"
# generate darwin amd64
# echo "[2] MacOS from arm64"
# generate darwin arm64
echo "[3] Linux from amd64"
generate linux amd64
# echo "[4] Linux from arm64"
# generate linux arm64
# echo "[5] Windows from amd64"
# generate_exe amd64
# echo "[6] Windows from arm64"
# generate_exe arm64