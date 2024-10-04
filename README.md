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