# go-getting-started

Learn Basic Golang
  
## Session 1  
  
### Khai báo biến  
  
```go  
var a int  
b := "Hello"  
  
// init struct
type Person struct {  
    Name string  
    Age int  
}  
  
var p1 Person  
p1.Name = "John"  
  
p2 := Person{  
    Name: "John",  
    Age: 20,  
}  
  
p3 := &Person{  
    Name: "John",  
    Age: 20,  
}  
```

### for loop

```go  
a := []string{"a", "b", "c"}  
for index, value := range a {  
    fmt.Println(index, value)  
}

# output
0 a
1 b
2 c
```  
  
Lôi hay gặp khi dùng for loop: update biến value sẽ không update giá trị trong mảng hay slice mà phải dùng index để update  
  
```go  
a := []string{"a", "b", "c"}  
for index, value := range a {  
    value = value + "_update_value"
    fmt.Println(index, value)  
}    
fmt.Println(a)

# output
0 a_update_value
1 b_update_value
2 c_update_value
[a b c]
```  
  
muốn update phải dùng biến index  
  
```go  
a := []string{"a", "b", "c"}  
for index, value := range a {  
    a[index] = value + "_update_value"  
    fmt.Println(index, value)  
    fmt.Println(index, a[index])  
}
fmt.Println(a)  

# output
0 a
0 a_update_value
1 b
1 b_update_value
2 c
2 c_update_value
[a_update_value b_update_value c_update_value]
```

### switch-case

```go
var input string
switch input {
	case "value1":
		println("value1")
	case "value2":
		println("value2")
	default:
		println("default")
}
```

Lưu ý: không cần `break` ở các case, input match case nào sẽ chỉ chạy code trong case đó, không match thì chạy ở phần `default`

### routine

Tạo routine bằng cách thêm từ khóa `go` trước hàm cần chạy

```go
func printHello() {
	fmt.Println("hello")
}

func main() {
	go printHello() // go routine
	time.Sleep(1*time.Second)
}
```

#### WaitGroup

khi cần chờ routine chạy xong thì dùng `sync.WaitGroup`

```go
func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start task 1")
	time.Sleep(1 * time.Second)
	fmt.Println("end task 1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start task 2")
	time.Sleep(2 * time.Second)
	fmt.Println("end task 2")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start task 3")
	time.Sleep(3 * time.Second)
	fmt.Println("end task 3")
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go task1(wg)
	go task2(wg)
	go task3(wg)
	wg.Wait()
}
```

#### Worker pool

Dùng pattern này khi cần xử lý nhiều công việc cùng lúc, mà muốn kiểm soát số lượng go-routine tạo ra để tránh mất kiểm soát,
thay vì tạo 1 routine cho 1 việc thì fix số lượng routine rồi dùng channel để điều phối công việc vào các routine

```go
type worker struct {
	id int
	c  chan int
}

func (w *worker) process() {
	for {
		data, ok := <-w.c
		if ok {
			println("worker ", w.id, " processing data ", data)
		} else {
			// channel closed
			fmt.Println("Worker ", w.id, " exiting")
			return
		}
	}
}

func main() {
	c := make(chan int)

	w1 := &worker{id: 1, c: c}
	w2 := &worker{id: 2, c: c}
	w3 := &worker{id: 3, c: c}

	go w1.process()
	go w2.process()
	go w3.process()

	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}

	close(c)
	time.Sleep(1 * time.Second)
}
```