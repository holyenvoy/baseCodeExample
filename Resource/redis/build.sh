CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  baseCode/baseCodeExample/redis

scp -P 34185 redis  wdb@192.168.129.220:/home/wdb

