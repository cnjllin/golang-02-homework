# Homework-4

实现一个RESTful的KV存储，能持久化保存数据，并且在启动的时候载入之前保存的数据

## Set接口

```bash
http://192.168.1.8:4000/set/Key/Value
```
写入Key:Value

## Get接口
http://192.168.1.8:4000/get/xxxxxx
获取xxxxx对应的Value，如果不存在就返回HTTP Code 404

## 持久化
每次Set操作都进行存盘操作

