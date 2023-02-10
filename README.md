##

## 演示

![](https://s3.bmp.ovh/imgs/2023/02/10/350bdf7a09527510.png)

## 上手

你只需要去openai申请一个key，如下图填写到代码即可。<mark>注意！！！！Bearer和中间的空格需要保留</mark>

openai的注册方式网上有很多，这里就不作过多赘述了。

https://openai.com/

![](file://C:\Users\Administrator\AppData\Roaming\marktext\images\2023-02-10-10-55-30-image.png?msec=1675997730977)

## 环境

个人环境是go 1.17.10，没有什么特殊的包，应该没有go版本限制

##

## 编译

```
## 都是go自带的编译方式，有经验的可以直接跳过


windows

## 直接控制台运行
cd cmd
go run main.go

## 打包成exe可执行文件
cd cmd
go build

## 运行
将会在目录生成一个openAi.exe文件，运行即可

linux下
cd cmd
go build

## 运行
./openAi
```