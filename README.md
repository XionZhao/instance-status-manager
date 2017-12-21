# 通过tag对aws ec2实例进行stop和start

> 所在机器需要配置`aws_access_key_id`和`aws_secret_access_key`

## 依赖

```
go get -u github.com/aws/aws-sdk-go
```

## 修改tag
进入到`stop_tag_instance`和`start_tag_instance`
```
$ vim main.go
const (
	tagkey   string = "Env"
	tagvalue string = "SystemTest"
)

tagkey对应aws tag的key
tagvalue对应aws tag的value
```
## Build

```
$ cd start_tag_instance
$ go build .

$ cd stop_tag_instance
$ go build .
```

## 执行crontab

可以根据需求进行stop和start
