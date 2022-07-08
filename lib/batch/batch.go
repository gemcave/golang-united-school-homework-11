package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	res = make([]user, n)
	ch_users := make(chan struct{}, pool)

	for i := 0; i < int(n); i++ {
		wg.Add(1)
		ch_users <- struct{}{}

		go func(i int64) {
			defer wg.Done()
			res[i] = getOne(i)
			<-ch_users
		}(int64(i))
	}
	wg.Wait()

	return res
}
