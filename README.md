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

> url: (required) the url you want to take a snapshot
> 
> target: (required) where to save the snaphot (absolute path)
> 
> path: (optional) the path you upload to qiniu
> 
> device: (optional) use which device to take snapshot, now only support “mac"

### FAQ

1. [使用casperjs截出优雅的图片](http://type.so/linux/casperjs-capture-nice.html)
   
2. Multi network environment?
   
   Just deploy multi kaka instance, and use different nsq channel, that will be OK.
   
3. I don’t want to upload to qiniu
   
   Just let the path empty.

### Development

``` shell
go get -u github.com/joho/godotenv
go get -u github.com/qiniu/api.v7
go get -u github.com/bitly/go-nsq
```