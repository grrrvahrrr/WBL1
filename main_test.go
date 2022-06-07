package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"sync"
	"testing"
	"time"
)

//Task 1
type Human struct {
	name string
}

type Human1 struct {
	name string
}

type Action struct {
	Human
	Human1 Human1
}

func TestTask1(t *testing.T) {

	sam := Human{
		name: "sam",
	}

	tom := Human1{
		name: "tom",
	}

	humanAction := Action{
		sam,
		tom,
	}
	humanAction1 := Action{
		Human{
			name: "jack",
		},
		Human1{
			name: "frost",
		},
	}

	fmt.Println(humanAction)
	fmt.Println(humanAction1)
}

//Task 2
func TestTask2(t *testing.T) {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, v := range array {
		wg.Add(1)
		go func(v int) {
			sqr := v * v
			fmt.Println(sqr)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

//Task 3
func TestTask3(t *testing.T) {
	array := []int{2, 4, 6, 8, 10}

	var sum int

	var wg sync.WaitGroup

	for _, v := range array {
		wg.Add(1)
		go func(v int) {
			sum += v * v
			wg.Done()
		}(v)

	}
	wg.Wait()

	fmt.Println(sum)
}

//Task5
func TestTask5(t *testing.T) {
	var worktime time.Duration = 1

	ch := make(chan int)

	go func() {
		for {
			ch <- 1
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	<-time.After(worktime * time.Second)
}

//Task7
func TestTask7(t *testing.T) {
	testmap := make(map[int]int)

	var wg sync.WaitGroup
	//Mutex to avoid fatal error - rw allows for 1 writer and multiple readers
	var m sync.RWMutex

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			m.Lock()
			testmap[i] = i
			m.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(testmap)

}

//Task8
func TestTask8(t *testing.T) {
	var num int64 = 123456
	bitToChange := 1

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(num))

	fmt.Println(b)

	for i := range b {
		if i == bitToChange {
			b[i] = 1
		}
	}
	fmt.Println(b)
	fmt.Println(int64(binary.LittleEndian.Uint64(b)))
}

//Task9
func TestTask9(t *testing.T) {
	numArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numCh := make(chan int)
	numX2Ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for _, v := range numArray {
			numCh <- v
		}
	}()

	go func() {
		for v := range numCh {
			numX2Ch <- v * 2
		}
	}()

	go func() {
		var counter int
		for v := range numX2Ch {
			fmt.Println(v)
			counter++
			if counter == len(numArray) {
				cancel()
			}
		}
	}()

	<-ctx.Done()
	close(numCh)
	close(numX2Ch)
}

//Task10
func TestTask10(t *testing.T) {
	tempArray := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := make(map[float64][]float32)

	for _, v := range tempArray {
		tempMap[math.Trunc(v/10)*10] = append(tempMap[math.Trunc(v/10)*10], float32(v))
	}
	fmt.Println(tempMap)
}

//Task11
func TestTask11(t *testing.T) {
	array1 := []int{1, 2, 2, 3, 4, 5, 5}
	array2 := []int{2, 2, 5, 6, 7, 8}
	var array3 []int

	m := make(map[int]bool)

	for _, item := range array1 {
		m[item] = true
	}

	for _, item := range array2 {
		if _, ok := m[item]; ok {
			array3 = append(array3, item)
		}
	}

	fmt.Println(array3)
}

//Task12
func TestTask12(t *testing.T) {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	var subarray []string
	var firstElement bool = true

	visited := make(map[string]bool)
	for _, v := range array {
		if visited[v] == true {
			subarray = append(subarray, v)
			if firstElement {
				subarray = append(subarray, v)
				firstElement = false
			}
		} else {
			visited[v] = true
		}
	}

	fmt.Println(subarray)
}

//Task13
func TestTask13(t *testing.T) {
	a, b := 1, 2
	fmt.Printf("a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)
}

//Task14
func TestTask14(t *testing.T) {
	//integer := 15
	//str := "string"
	// boolean := true
	ch := make(chan int)

	func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("It is int: %d\n", i)
		case string:
			fmt.Printf("It is string: %s\n", i)
		case bool:
			fmt.Printf("It is bool: %v\n", i)
		case chan int:
			fmt.Printf("It is chan int: %s\n", i)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}(ch)

}

//Task19
func TestTask19(t *testing.T) {
	hello := "Hello, 世界"
	var rev []rune

	for i := range []rune(hello) {
		rev = append(rev, []rune(hello)[len([]rune(hello))-i-1])
	}

	fmt.Printf("%s\n", string(rev))
}

//Task20
func TestTask20(t *testing.T) {
	str := "snow dog sun"
	var rev []string
	strSlice := strings.Split(str, " ")

	for i := range strSlice {
		rev = append(rev, strSlice[len(strSlice)-i-1])
	}

	revStr := strings.Join(rev, " ")
	fmt.Println(revStr)
}

//Task23
func TestTask23(t *testing.T) {
	//Break order
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	//Set i equal to last element
	a[i] = a[len(a)-1]
	//Delete last element
	a = a[:len(a)-1]

	fmt.Println(a)

	//Maintain Order
	b := []string{"A", "B", "C", "D", "E"}

	//Copy everything after i into i position
	copy(b[i:], b[i+1:])
	//Delete last element
	b = b[:len(b)-1]

	fmt.Println(b)
}

//Task24
type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func TestTask24(t *testing.T) {
	p1 := NewPoint(3, -17)
	p2 := NewPoint(26, 42)

	first := math.Pow(float64(p2.X-p1.X), 2)
	second := math.Pow(float64(p2.Y-p1.Y), 2)
	fmt.Println("Distance: ", math.Sqrt(first+second))
}

//Task25
func TestTask25(t *testing.T) {
	var sleepTime time.Duration = 3

	func(st time.Duration) {
		<-time.After(time.Second * sleepTime)
		fmt.Printf("I've been asleep for %d seconds\n", sleepTime)
	}(sleepTime)

}

//Task26
func TestTask26(t *testing.T) {
	//str := "abcd"
	//str := "abCdefAaf"
	str := "aabcd"

	for i := range []rune(str) {
		for ii := range []rune(str) {
			if i != ii && []rune(str)[i] == []rune(str)[ii] {
				fmt.Printf("%s - false\n", str)
				return
			}
		}
	}
	fmt.Printf("%s - true\n", str)
}
