package common

import "log"

type ConnLimiter struct {
	currentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		currentConn: cc,
		bucket: make(chan int,cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {

	if len(cl.bucket) >= cl.currentConn {
		log.Printf("api limited")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn()  {
	c := <- cl.bucket  //释放频道
	log.Printf("New connection comming: %d", c)
}