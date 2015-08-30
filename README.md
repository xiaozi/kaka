### Dependences

1. nsq
2. casperjs
3. phantomjs

### Installation

1. copy .env.example to .env
2. modify .env file. fill the properties.
3. run

``` shell
./kaka
```

### Usage

just put message to nsq <topic>, which you set in the .env file

message struct

``` json
{
	"url": "http://tool.lu/",
	"target": "/data/screenshots/WrTSV5zbkHPCqU6t.png",
	"path": "screenshots/WrTSV5zbkHPCqU6t.png",
	"device": "mac"
}
```

> url: the url you want to take a snapshot
> 
> target: where to save the snaphot
> 
> path: the path you upload to qiniu
> 
> device: use which device to take snapshot, now only support “mac"

### FAQ

1. [使用casperjs截出优雅的图片](http://type.so/linux/casperjs-capture-nice.html)

### Development

``` shell
go get -u github.com/joho/godotenv
go get -u github.com/qiniu/api.v7
go get -u github.com/bitly/go-nsq
```