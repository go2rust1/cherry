package dbms

import (
	"context"
	"github.com/go2rust1/cherry/src/conf"
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/mod/database"
)

type node struct {
	// db 数据库实例
	db database.Database

	// buffer 模型缓冲单元, 当缓冲区满, 就将缓冲区的数据压入数据库
	buffer []interface{}

	// pipeline 模型管道, 发送和接收来自用户的模型数据
	pipeline chan interface{}
}

func (n *node) wander() {
	counter.Increase()
	defer counter.Decrease()
	ctx, cancel := context.WithTimeout(context.Background(), conf.WaitDataTimeout)
	defer cancel()
rust:
	for {
		select {
		case model := <-n.pipeline:
			if len(n.buffer) < conf.BatchSize {
				n.buffer = append(n.buffer, model)
			} else {
				buffer := make([]interface{}, len(n.buffer))
				copy(buffer, n.buffer)
				n.buffer = []interface{}{model}
				go n.db.Insert(buffer...)
			}
		case <-ctx.Done():
			break rust
		}
	}
	buffer := make([]interface{}, len(n.buffer))
	copy(buffer, n.buffer)
	go n.db.Insert(buffer...)
}
