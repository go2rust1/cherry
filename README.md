# Cherry Web Crawling Framework
Cherry is a web crawling framework written in Go. It features a scrapy-like API. If you need fast and good productivity, you can try cherry.

#### The key features of Cherry are:
+ Fast
+ Concise
+ Multiple Topic Parallel
+ Multiple Model Binding
+ Multiple Database Binding
+ Auto Data Persistent Storage
+ Support: DB2、MySQL、Oracle

## Getting started

### Prerequisites
Go Version: >=1.17

### Getting Cherry
With Go module support, simply add the following import to your code.
```
import "github.com/go2rust1/cherry"
```
Otherwise, run the following command to install the cherry package
```
go get -u github.com/go2rust1/cherry
```

## Running Cherry
```
package main

import (
  "github.com/go2rust1/cherry"
)

type Model struct {
	F1 string `db:"F1"`
	F2 string `db:"F2"`
}

func Parser(topic cherry.Topic, response cherry.Response) {
	topic.Send(Model{F1: "", F2: ""})
}

func main() {
	db := cherry.MySQL()
	db.SetDSN("")
	db.SetTable("")
	db.Bind(Model{})
  
  _cherry := cherry.New()

  topic, _ := _cherry.NewTopic("TopicName")
  topic.Request("https://www.baidu.com/", Parser, nil)
  topic.Bind(db)

  _cherry.Start()
}
```

## Learn more examples
Learn and practice more examples, please read the
```
github.com/go2rust1/cherry/examples
```
