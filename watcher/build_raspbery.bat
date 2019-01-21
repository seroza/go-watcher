echo off

set GOOS=linux
set GOARCH=arm
set GOARM=5

call go build .