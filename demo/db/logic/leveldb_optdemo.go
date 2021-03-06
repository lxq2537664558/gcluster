package logic

import (
	"github.com/gosrv/gcluster/gbase/gdb/gleveldb"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

type serviceLevelDBOptDemo struct {
	gioc.IBeanCondition
	// 通过bean标签自动注入leveldb驱动
	leveldb gleveldb.ILevelDBDriver `bean:""`
}

func NewServiceLevelDBOptDemo() *serviceLevelDBOptDemo {
	return &serviceLevelDBOptDemo{
		// 依赖配置，只有配置了gleveldb.ILevelDBDriverType之后，这个bean才会生效
		IBeanCondition: gioc.NewConditionOnBeanType(gleveldb.ILevelDBDriverType, true),
	}
}

func (this *serviceLevelDBOptDemo) BeanStart() {
	this.leveldb.Set("acc", "12345")
	data, _ := this.leveldb.Get("acc")
	util.Assert(data == "12345", "assert failed")
}

func (this *serviceLevelDBOptDemo) BeanStop() {

}
