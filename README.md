# go-tutorial
https://academy.nomadcoders.co/ 노마드코더 go 시작하기

https://golang.org

# 실행

> go run main.go


실행을 위해서는 main.go 파일이 있어야 하고, 해당 파일은 아래와 같이 작성해야 한다. main() 이 시작 포인트 지점.

```go
package go

func main() {

}
```

함수명을 대문자로 시작하면 export 할 수 있고, 소문자로 시작하면 내부 함수로만 사용된다. 

```go
func main() {
    something.SayHello()    // 가능
    // something.sayBye()   // 불가능
}
```

# 상수와 변수

#### 상수
```go
func main() {
    // const 변수명 타입 = 값
    const name string = "bear"
}
```

#### 변수
```go
func main() {
    // var 변수명 타입 = 값
    var name string = "bear"
}
```

#### 변수 축약형
func 안에서 사용 가능
```go
func main() {
    // 변수명 := 값
	name := "bear"
}
```

# function

func 함수명(매개변수명 타입[,...]) 반환타입 {} 
```go
func multiply(a int, b int) int {
    return a * b
}
```

매개변수 타입이 같다면 마지막에 하나만 써도 된다.
```go
func multiply(a, b int) int {
    return a * b
}
```

여러 개의 매개변수를 받고 싶다면, ... 를 사용한다.
```go
func repeatMe(words ...string) {
    fmt.Println(words)
}

func main() {
    repeatMe("a", "b", "c", "d", "e")
}
```

go문법에 특이한 점 중 하나는 여러 개의 return을 반환할 수 있다.
```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
    totalLength, upperName := lenAndUpper("treasure")
}
```

여러 개의 반환 값이 있지만, 그 중 일부를 무시하고 싶다면 _ 를 사용한다. 
```go
func main() {
    totalLength2, _ := lenAndUpper("treasure")
}
```