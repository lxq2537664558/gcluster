package gmongo

import (
	"github.com/globalsign/mgo"
	"github.com/gosrv/gcluster/gbase/gl"
)

type MongoMessageQueue struct {
	queue   string
	session *mgo.Session
}

func NewMongoMessageQueue(queue string, session *mgo.Session) *MongoMessageQueue {
	return &MongoMessageQueue{queue: queue, session: session}
}

func (this *MongoMessageQueue) Pop(num int) []string {
	gl.Panic("not implement")
	return nil
}

func (this *MongoMessageQueue) Push(msg string) bool {
	gl.Panic("not implement")
	return false
}

func (this *MongoMessageQueue) Name() string {
	return "mongo"
}
