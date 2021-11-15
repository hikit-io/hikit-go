package hkgo

import "context"

func Async(def func(), fs ...func(ctx context.Context) error) error {
	err := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())
	for _, f := range fs {
		go func(f func(context.Context) error) {
			err <- f(ctx)
		}(f)
	}

	for e := range err {
		if e != nil {
			cancel()
			def()
			return e
		}
	}

	def()
	return nil
}
