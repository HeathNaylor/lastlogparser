```
root@ubuntu:~/lastlogparser# go build && ./lastlogparser
2017/10/26 02:21:09 &main.UserInfo{Name:"root", Line:"pts/19", Host:"127.0.0.1", Last:"2017-09-01 09:40:45 -0700 PDT"}
2017/10/26 02:21:09 &main.UserInfo{Name:"daemon", Line:"", Host:"", Last:"**Never logged in**"}
2017/10/26 02:21:09 &main.UserInfo{Name:"bin", Line:"", Host:"", Last:"**Never logged in**"}
2017/10/26 02:21:09 &main.UserInfo{Name:"sys", Line:"", Host:"", Last:"**Never logged in**"}
2017/10/26 02:21:09 &main.UserInfo{Name:"sync", Line:"", Host:"", Last:"**Never logged in**"}
...
2017/10/26 02:21:09 &main.UserInfo{Name:"sshd", Line:"", Host:"", Last:"**Never logged in**"}
2017/10/26 02:21:09 &main.UserInfo{Name:"lolkin", Line:"pts/18", Host:"127.0.0.1", Last:"2017-10-25 08:09:29 -0700 PDT"}
```