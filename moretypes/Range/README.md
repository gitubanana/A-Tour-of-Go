### Range
```go
package main

import "fmt"

var people = []string{
  "euiclee",
  "hyeoan",
  "jinheo",
  "nakoo",
  "jinhyeok",
  "jaehjoo",
}


func main() {
  for i, v := range people {
    fmt.Printf("index : %d, name : %v\n", i, v)
  }
}

/*
index : 0, name : euiclee
index : 1, name : hyeoan
index : 2, name : jinheo
index : 3, name : nakoo
index : 4, name : jinhyeok
index : 5, name : jaehjoo
*/
```
`for`에서 `range`를 사용하면, `slice`나 `map`을 순회할 수 있다.<br>

`slice`를 `ranging`할 때는, 두 개의 값이 각 원소를 순회할 때마다 리턴된다.<br>
첫 번째 값은 `index`이고, 두 번째 값은 `그 인덱스에 해당되는 원소의 복사본`이다.<br>

### go run
![image](https://github.com/user-attachments/assets/60853d20-a510-4133-becf-752cbd387430)


### Reference
[Moretypes/Range](https://go.dev/tour/moretypes/16)<br>
