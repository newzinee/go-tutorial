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

## 상수
```go
func main() {
    // const 변수명 타입 = 값
    const name string = "bear"
}
```

## 변수
```go
func main() {
    // var 변수명 타입 = 값
    var name string = "bear"
}
```

## 변수 축약형
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

## naked return 
function return 부분에 미리 변수를 세팅할 수 있다. 

length와 uppercase는 return 영역에서 이미 선언된 상태.
```go
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
```

## defer
함수가 실행된 후, 마지막에 실행되는 코드

여기서, 함수 호출이 끝난 다음 "I'm done"이라는 메시지가 나온다. 
```go
func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
```

# for
반복문 사용을 위해서는 for를 쓴다. 
```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

## range
array 같은 배열을 반복할떄는 range를 쓴다. 

반환은 순서대로 index, vaule
```go
for index, number := range numbers {
    fmt.Println(index, number)
}
```

여기서도, 쓰지 않을 변수라면 _를 사용한다. 
```go
for _, number := range numbers {
    fmt.Println(number)
}
```

# if
조건문을 사용하기 위해서 if-else를 사용한다.
```go
func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}
```

조건문 내에서만 사용할 변수를 만들 수 있다. 
```go
func canIDrink2(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
```

# switch
switch를 이용해서도 조건문을 만들 수 있다. 

switch 변수 { case 값: }
```go
func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
```

case 문 안에 조건을 넣을 수도 있다. 
```go
func canIDrink2(age int) bool {
	switch {
	case age > 18:
		return true
	case age == 18:
		return true
	}
	return false
}
```

switch 안에서 사용할 변수도 생성해줄 수 있다. 
```go
func canIDrink3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
```

# pointer

## &
주소값을 알기 위해서는 & 사용
```go
func main() {
	a := 2
    b := &a
    fmt.Println(b)
}
```
b에는 a의 주소값이 담기고, 출력하면 a의 주소가 출력됨.

## *
해당 주소값의 값을 접근하기 위해서는 * 사용 
```go
func main() {
	a := 2
	b := &a
	a = 5
	fmt.Println(&a, b)
	fmt.Println(*b)
	*b = 20
	fmt.Print(a)
}
```
*b 를 통해 b의 저장된 주소의 값, 즉 a의 값을 볼 수 있음.

- fmt.Println(*b) -> 출력: 5

*b 를 통해 b의 저장된 주소의 값, 즉 a의 값을 변경할 수 있음.

- *b = 20 -> b의 저장된 주소의 값을 20으로 수정 -> b의 저장된 주소=a의 주소, b의 저장된 주소의 값=a에 저장된 값 -> a가 20으로 수정됨