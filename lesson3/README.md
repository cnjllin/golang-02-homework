# 作业:

## 要求:

编写一个mini ps，输出两列: pid和命令行

类似 `1234 /bin/bash`


## 思路:

1. 读取/proc目录下的文件
2. 获取所有pid子目录
3. 读取pid子目录下的cmdline文件
