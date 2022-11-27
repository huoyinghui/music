server:
  address:     ":8000"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"
  serverRoot: "./resource"
  LogPath: /tmp/log/admin/server
logger:
  Path: /tmp/log/admin
  Level: all
  Stdout: true
  debug:
    path: /tmp/log/admin/debug
    level: dev
    stdout: true

database:
  default:
    link:  "mysql:root:mysecretpassword@tcp(127.0.0.1:3306)/godb"
    debug:   true
  logger:
    Path:  /tmp/log/admin/sql
    Level: all
    Stdout: true

redis:
  default:
    address: 127.0.0.1:6379
    db: 9

app:
  musicTask:
    dirPath: /Users/huoyinghui/Music/网易云音乐/
    todo:
      100:
        user: 陈奕迅
        name: 孤勇者
        id: 1901371647
        link: https://music.163.com/#/song?id=1901371647
      101:
        user: 周杰伦
        name: 屋顶
        link: https://music.163.com/#/song?id=298317