### Basic types
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
`Go`에는 다양한 타입이 있다.<br>
`int`, `uint`, `uintptr` 타입은 32bit 시스템에서는 32bit이고, 64bit 시스템에서는 64bit이다.<br>
정수를 저장할 경우 특별한 이유가 없다면, `int`를 사용해야 한다.<br>

### Factored variable declarations
```go
import (
  "fmt"
  "math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```
`import`와 마찬가지로, `var`도 `factored`될 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/33e5cc95-7d62-4208-a80b-ff8f04519fa9)

### Reference
[Basics/Basic types](https://go.dev/tour/basics/11)<br>
