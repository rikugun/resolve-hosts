# Auto Resolve server to ip

每隔10分钟 解析servers.txt 中的域名到 hosts 文件中，通过 http://ip:3000/hosts 获取最新文件。

通过替换容器文件(-v /path/servers.txt:/app/servers.txt) 实现解析不同的域名
