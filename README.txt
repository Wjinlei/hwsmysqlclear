【软件简述】

软件名称：护卫神·挂马自动清理系统

软件功能：自动清理mysql数据库中的<script></script>恶意代码！




【手工启动】

运行命令：./hwsmysqlclear run -u root -p 密码 -db 数据库名 -t 30

注意：每次重启服务器后需要手工启动该软件




【服务启动】

1、修改hwsmysqlcleard中的数据库连接信息

2、注册并启动服务，运行命令：sudo make install


注意：每次重启服务器后会自动启动该软件

卸载方法：sudo make uninstall




【帮助说明】

运行以下命令查看帮助：./hwsmysqlclear help run

Usage of command "run":

        hwsmysqlclear run [options]

Options:

  -db string
        database name
  -exclude string
        Exclude tables, comma separated
  -include string
        Include tables, comma separated
  -p string
        password
  -t int
        How many seconds between scans (default 10)
  -u string
        username (default "root")


