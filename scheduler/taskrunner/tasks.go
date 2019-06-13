package taskrunner

import (
	"github.com/pkg/errors"
	"os"
	"sync"
	"video_server/scheduler/models"
)

func deleteVideo(vid string) error {
	err := os.Remove("/videos/" + vid)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := models.ReadDeleteVideoR(3)
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	//将id送入channel
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop: 
		for {
			select {
			case vid := <- dc:
				go func(id interface{}) {
					//删除视频文件
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := models.DelDeleteVideoR(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
			default:
				break forloop
			}
		}

	//错误反馈
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}