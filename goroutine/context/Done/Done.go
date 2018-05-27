package main

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}

func Stream(ctx context.Context, out chan<- Value) error {

	for {
		v, err := DoSomething(ctx)

		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():

			return ctx.Err()
		case out <- v:
		}
	}
}

func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}
}

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}

type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

// A cancelCtx can be canceled. When canceled, it also cancels any children

// that implement canceler.
type cancelCtx struct {
	Context

	done     chan struct{} // closed by the first cancel call.
	mu       sync.Mutex
	children map[canceler]struct{} // set to nil by the first cancel call

	err error // set to non-nil by the first cancel call
}
