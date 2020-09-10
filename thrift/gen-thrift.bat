@echo off

RMDIR /S /Q "thriftmodel"
call "./thrift-0.10.0-patched.exe" -I . -out . -strict -r --gen go payload.thrift

pause
