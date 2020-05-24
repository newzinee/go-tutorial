# go-tutorial
https://academy.nomadcoders.co/ 노마드코더 go 시작하기

실행

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