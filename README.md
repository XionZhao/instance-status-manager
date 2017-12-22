# 通过tag对aws ec2实例进行stop和start

> 所在机器需要配置`aws_access_key_id`和`aws_secret_access_key`

## 依赖

```
go get -u github.com/aws/aws-sdk-go
```

## Build

```
$ go build .
```

## 执行命令
```
$ ./instance-status-manager -h
Usage of ./instance-status-manager:
  -status string
    	Select start or stop
  -tagkey string
    	Choose to stop or start ec2 tag key
  -tagvalue string
    	Choose to stop or start ec2 tag value
```
### 启动实例
```
./instance-status-manager -tagkey=Name -tagvalue=SystemTest -status=start
```
### 停止实例
```
./instance-status-manager -tagkey=Name -tagvalue=SystemTest -status=stop
```
## 执行crontab

可以根据需求进行stop和start
