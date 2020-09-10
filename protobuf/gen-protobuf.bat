@echo off

rem RMDIR /S /Q "thriftmodel"
call "./protoc.exe" -I . --go_out . payload.proto

pause
