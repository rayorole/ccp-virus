@echo off

echo Compiling...

go build -ldflags -H=windowsgui -o build/ccp.exe