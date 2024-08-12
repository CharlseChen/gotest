package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"sync/atomic"
	"github.com/panjf2000/ants/v2"
)

func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()
	return c
}
func main() {
	//f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	//defer f.Close()
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	//t := NewTest()
	//println(<-t)
	// 等待 goroutine 结束返回。
	//Signal()
	//Test()
	//MapToStruct()
	//TimeCompare()
	// JsonTest()
	//MapTest()
	//n := 10
	//for i := 0; i < 5; i++ {
	//	nums := generate(n)
	//	bubbleSort(nums)
	//	n *= 10
	//}
	//MemOpt()
	//MyPrint()
	//excel.ExcelHandle.Do()
	//MyArr()
	//PrintFloat()
	// fmt.Printf("%.6f", float64(time.Now().Unix())/1e10)
	//fmt.Printf("%d", gconv.Int64(4)+gconv.Int64(601))
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}

	// Use the MultiPool and set the capacity of the 10 goroutine pools to unlimited.
	// If you use -1 as the pool size parameter, the size will be unlimited.
	// There are two load-balancing algorithms for pools: ants.RoundRobin and ants.LeastTasks.
	mp, _ := ants.NewMultiPool(10, -1, ants.RoundRobin)
	defer mp.ReleaseTimeout(5 * time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mp.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", mp.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the MultiPoolFunc and set the capacity of 10 goroutine pools to (runTimes/10).
	mpf, _ := ants.NewMultiPoolWithFunc(10, runTimes/10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	}, ants.LeastTasks)
	defer mpf.ReleaseTimeout(5 * time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mpf.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", mpf.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500*2 {
		panic("the final result is wrong!!!")
	}
}

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func Signal() {
	var wg sync.WaitGroup
	wg.Add(4)
	s := make(chan int, 1)
	for i := 0; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
			s <- 1
			for j := 0; j < 3; j++ {
				fmt.Printf("id:%v,i:%d\n", id, j)
			}
			<-s
		}(i)
	}
	wg.Wait()

}

func Test() {
	val := "[127565609,4023409]"
	uids := make([]interface{}, 0)

	if err := json.Unmarshal([]byte(val), &uids); err != nil {
		fmt.Printf("%v\n", uids)
	}
	fmt.Printf("%v 2\n", uids)
	for _, u := range uids {
		if gconv.Uint32(u) == 127565609 {
			fmt.Printf("相等")
		}
	}
}

type res struct {
	Test int64 `json:"test"`
	Jso  int64 `json:"jso"`
}

func MapToStruct() {
	var source = map[string]string{"test": "12344", "jso": "23342344"}
	target := &res{}
	var tmap = make(map[string]int64, len(source))
	for k, v := range source {
		c, _ := strconv.Atoi(v)
		tmap[k] = int64(c)
	}
	err := mapstructure.Decode(tmap, target)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Printf("%#v", target)
}

func TimeCompare() {
	res, _ := strconv.Atoi("")
	fmt.Printf("res:%v", res)
	t1, _ := time.Parse("2006-01-02", "2022-08-30")
	fmt.Printf(t1.Add(24*time.Hour - 1).Format("2006-01-02 15:04:05"))
	if t1.After(t1) {
		fmt.Printf("is Today")
		return
	}
	fmt.Printf("is Before")
}

func MapTest() {
	m := make(map[int]int, 3)
	x := len(m)
	m[1] = m[1]
	y := len(m)
	println(x, y)
}

func JsonTest() {
	var target res
	_ = json.Unmarshal([]byte(string(`{"test": 12344, "jso": 23342344}`)), &target)
	fmt.Printf("正确示范 target:%v\n", target)

	var target1 = &res{}
	_ = json.Unmarshal([]byte(string(`{"test": 12344, "jso": 23342344}`)), target1)
	fmt.Printf("正确示范 target1:%v\n", target1)

	var target2 *res
	_ = json.Unmarshal([]byte(string(`{"test": 12344, "jso": 23342344}`)), &target2)
	fmt.Printf("正确示范:target2:%v\n", target2)

	var target3 res
	err := json.Unmarshal([]byte(string(`{"test": 12344, "jso": 23342344}`)), target3)
	fmt.Printf("错误示范 target3:%v err:%v\n", target3, err)

	var target4 *res
	err1 := json.Unmarshal([]byte(string(`{"test": 12344, "jso": 23342344}`)), target4)
	fmt.Printf("错误示范 target4:%v err:%v\n", target4, err1)

	var target5 = new(res)
	_ = json.Unmarshal([]byte(string(`{"test": 12344, "jso": 23342344}`)), target5)
	fmt.Printf("正确示范target5:%v\n", target5)

	var target6 []map[string]interface{}
	str := "[{\"attach\":\"202302/06/139169904_63e09d3675daa9.63637965.jpg?960*1280\",\"tpid\":1675664695693000,\"uid\":139169904},{\"attach\":\"202302/06/139169904_63e077d78c79e1.81689778.jpg?1280*960\",\"tpid\":1675655128663000,\"uid\":139169904},{\"attach\":\"202302/06/139169904_63e0779a5af7a5.73123332.jpg?1280*753\",\"tpid\":1675655068057000,\"uid\":139169904}]"
	_ = json.Unmarshal([]byte(string(str)), target6)
	fmt.Printf("错误示范target6:%v\n", target6)

	var target7 []map[string]interface{}
	str1 := "[{\"attach\":\"202302/06/139169904_63e09d3675daa9.63637965.jpg?960*1280\",\"tpid\":1675664695693000,\"uid\":139169904},{\"attach\":\"202302/06/139169904_63e077d78c79e1.81689778.jpg?1280*960\",\"tpid\":1675655128663000,\"uid\":139169904},{\"attach\":\"202302/06/139169904_63e0779a5af7a5.73123332.jpg?1280*753\",\"tpid\":1675655068057000,\"uid\":139169904}]"
	_ = json.Unmarshal([]byte(string(str1)), &target7)
	fmt.Printf("错误示范target7:%v\n", target7)
}

type demo1 struct {
	A int8
	b int16
	c int32
}

type demo2 struct {
	A int8
	c int32
	b int16
}

// 属性名称顺序对内存的影响
func MemOpt() {
	var str int
	fmt.Println(unsafe.Alignof(demo1{}))
	fmt.Println(unsafe.Alignof(demo2{}))
	fmt.Println(unsafe.Sizeof(demo1{}))
	fmt.Println(unsafe.Sizeof(demo2{}))
	fmt.Println(unsafe.Sizeof(str))
}

func MyPrint() {
	a := &demo1{A: 8}
	fmt.Printf("%v ,%+v, %#v", a, a, a)
}

func MyArr() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func PrintFloat() {
	f4 := decimal.NewFromInt32(1)
	f5 := decimal.NewFromFloat(float64(1000-80) / 10e4)
	score, _ := f4.Add(f5).Float64()
	//fmt.Printf("%v \n", score)
	f1 := decimal.NewFromFloat(score)
	f2 := decimal.NewFromInt32(1000)
	f3 := decimal.NewFromInt32(int32(score))
	//1000-(score-int32(score))*10000   score整数部分代表年兽id,公式代表剩余血量
	remainBlood, _ := f2.Sub(f1.Sub(f3).Mul(decimal.NewFromInt32(10e4))).Float64()
	fmt.Printf("%v", int64(remainBlood))

	score1 := float64(1000-175)/100000 + float64(1)
	remainBlood1 := 1000 - math.Round((score1-float64(int(score1)))*100000)
	fmt.Println("\nremainBlood1", remainBlood1)
}

func Map() {
	res := make(map[int32][]struct{})
	res[12243234] = append(res[12243234], struct{}{})
}

type Option func(*client)
type client struct {
	Name       string
	Port       uint32
	Protocol   string
	MaxConnect uint32
	Timeout    time.Duration
}

func (c *client) NewClient(name string, port uint32, optionFuncs ...Option) *client {
	srv := &client{
		Name: name,
		Port: port,
	}
	for _, option := range optionFuncs {
		option(srv)
	}
	return srv
}

func OptionExcise() {
	c := &client{}
	server := c.NewClient("测试", 2000, Timeout(10*time.Second))
	fmt.Printf("name:%v", server.Name)
}

func Timeout(timeout time.Duration) Option {
	return func(c *client) {
		c.Timeout = timeout
	}
}
