Go语言开发的MySql挂码清理工具，只能清理`<script>`挂码，欢迎PR。

## Requirements
- [x] `Golang`

## Build
```sh
go build
```
### Static Build
- Install `musl` to `/usr/local/musl`
```sh
CC=/usr/local/musl/bin/musl-gcc go build -a -ldflags '-s -w -linkmode "external" -extldflags "-static"'
```

## Install
- 修改 `hwsmysqlcleard` 填写你的MYSQL帐号密码
- 运行 `sudo make install`

## Command
如果某位不写，则为默认值
- run
  ```sh
  -u  默认值 "root"  // MySQL用户名
  -p  默认值 ""      // MySQL密码
  -db 默认值 ""      // 数据库
  -t  默认值 10      // 每几秒扫描一次,最小10秒
  -include 默认值 "" // 指定只扫描哪些表,逗号分隔
  -exclude 默认值 "" // 指定排除哪些表,逗号分隔
  ```
- version

## FAQ
- `./hwsmysqlclear` or `./hwsmysqlclear help [command]`
