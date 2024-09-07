### Short variable declarations
- **함수 안에서**<br>
  `:=`는 `var`를 대신해 사용될 수 있다.<br>
- **함수 밖에서**<br>
  함수 밖에서는 모든 구문이 키워드(`var`, `func`, ...)로 시작해야 하기 때문에, `:=`는 사용 불가능하다.<br>
```go
package main

import "fmt"

// syntax error: non-declaration statement outside function body
// hi := "hi"

func main() {
  c, python, java := true, false, "no!"

  fmt.Println(c, python, java)
}
```

### go run
![image](https://github.com/user-attachments/assets/522dd5ae-6474-432c-87ca-89c670ec8f92)


### Reference
[Basics/Short variable declarations](https://go.dev/tour/basics/10)<br>
