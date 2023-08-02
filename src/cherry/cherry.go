package cherry

import (
	"github.com/go2rust1/cherry/src/dbms"
	"github.com/go2rust1/cherry/src/errors"
	"github.com/go2rust1/cherry/src/sauce"
	"github.com/go2rust1/cherry/src/trait"
)

type cherry struct {
	topic map[string]struct{}
	done  chan struct{}
}

// NewTopic 创建主题, 主题不能重复
func (c *cherry) NewTopic(name string) (trait.Topic, error) {
	if _, ok := c.topic[name]; ok {
		return nil, errors.DuplicateTopicError
	}
	topic := sauce.New(name)
	c.topic[name] = struct{}{}
	return topic, nil
}

// Start 启动任务
func (c *cherry) Start() {
	defer dbms.DBMS.Close()
	dbms.DBMS.Initialize()
	go c.RoundRobin()
	go c.HeartBeatDetection()
	<-c.done
}

// New 创建任务
func New() trait.Cherry {
	return &cherry{
		topic: make(map[string]struct{}),
		done:  make(chan struct{}),
	}
}
