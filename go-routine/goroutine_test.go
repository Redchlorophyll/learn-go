package go_routine

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello World")
}

func TestHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(num int) {
	fmt.Println("display", num)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(20 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dhonni"
		fmt.Println("channel uploading complete")
	}()

	data := <-channel
	fmt.Println(data)

	defer close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dhonni"
	fmt.Println("channel uploading complete")
}

func TestCreateChannelWithParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	defer close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dhonni"
	fmt.Println("channel uploading complete")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	OnlyOut(channel)

	defer close(channel)

}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	channel <- "Dhonni"
	channel <- "ari"
	channel <- "hendra"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	defer close(channel)

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke-" + strconv.Itoa(i)
		}

		defer close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel := make(chan string)
	otherChannel := make(chan string)

	go GiveMeResponse(channel)
	go GiveMeResponse(otherChannel)

	counter := 0

	for {
		select {
		case data := <-channel:
			fmt.Println("data from channel 1 ", data)
			counter++
		case data := <-otherChannel:
			fmt.Println("data from channel 2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel := make(chan string)
	otherChannel := make(chan string)

	go GiveMeResponse(channel)
	go GiveMeResponse(otherChannel)

	counter := 0

	for {
		select {
		case data := <-channel:
			fmt.Println("data from channel 1 ", data)
			counter++
		case data := <-otherChannel:
			fmt.Println("data from channel 2 ", data)
			counter++
		default:
			fmt.Println("waiting data...")
		}

		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total data ", counter)
}

func TestFixRaceConditionWithMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total data ", counter)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) ReadBalance() int {
	account.RWMutex.RLock()
	result := account.Balance
	account.RWMutex.RUnlock()

	return result
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.ReadBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("last balance ", account.ReadBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Dhonni",
		Balance: 10000,
	}

	user2 := UserBalance{
		Name:    "Ari",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)

	time.Sleep(5 * time.Second)

	fmt.Println("name: ", user1.Name, ", balance: ", user1.Balance)
	fmt.Println("name: ", user2.Name, ", balance: ", user2.Balance)
}

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(group)
	}

	group.Wait()
	fmt.Println("====complete====")
}

func OnlyOnce(num *int) {
	*num++
}

func TestOnce(t *testing.T) {
	counter := 0
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(func() {
				OnlyOnce(&counter)
			})
		}()
	}

	group.Wait()
	fmt.Println("Counter => ", counter)
}

func TestPool(t *testing.T) {
	pool := sync.Pool{}
	group := sync.WaitGroup{}

	pool.Put("Dhonni")
	pool.Put("Ari")
	pool.Put("Hendra")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("====complete====")
}

func AddMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, " : ", value)
		return true
	})
}

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total data ", counter)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())

	group.Wait()
}

func TestTicker(t *testing.T) {
	// this test will fail and it expected because
	// this test only check if ticker will stop when ticker.Stop()
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	// this test will fail and it expected because
	// this test only check if ticker will stop when ticker.Stop()
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

func TestGetGomaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread)
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine", totalGoroutine)
}

func TestGetAdjustedThread(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(12)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread)
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine", totalGoroutine)
}
