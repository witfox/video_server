package taskrunner

type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longlived bool
	Dispatcher fn
	Executor fn
}

//初始化
func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error: make(chan string, 1),
		Data: make(chan interface{}, size),
		dataSize:size,
		longlived: longlived,
		Dispatcher: d,
		Executor: e,
	}
}

func (r *Runner) startDispatch() {
	//释放channel
	defer func() {
		if !r.longlived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()
	for {
		select {
		case c := <- r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_EXECUTE
				}
			}
		case e := <- r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) startAll()  {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
