### Numeric Constants
- **C의 문제점**
  ```c++
  unsigned int u = 1e9;
  long signed int i = -1;
  ... i + u ...
  ```
  C에서는 서로 다른 타입을 연산 할 때, [암시적 형변환](https://en.cppreference.com/w/cpp/language/usual_arithmetic_conversions)이 일어난다.<br>
  이로 인해, 다양한 버그가 나타날 수 있다.<br>

- **Go의 해결 방법**
  ```go
  var u int
  var i int

  // 명시적으로 타입을 지정해줘야 한다.
  uint(i) + u
  i + int(u)

  // invalid operation: u + i (mismatched types uint and int)
  // i + u
  ```
  `Go`는 이 지뢰밭을 피하기 위해, 서로 다른 타입을 연산하지 않게 만들어졌다.<br>
  하지만, 명시적 형변환으로 인해 코드가 지저분해진다는 단점이 있다.<br>
  
  그렇다면, `상수(constant)`도 형변환을 해줘야 할까?<br>
  `i = int(0)`과 같이 `상수(constant)`의 형변환을 요구하는 것은 불합리적일 것이다.<br>
  따라서, `Go`는 `C`같은 언어들과 다르게 `상수(constant)`를 만들었다.<br>

- **typed constant**
  ```go
  // string-typed constant
  const typedHello string = "Hello, 世界"

  type MyString string
  var m MyString = typedHello // Type error

  // float64-typed constant
  const typedZero float64 = 0.0

  type MyFloat64 float64
  var mf Myfloat64 = typedZero // Type error
  ```
  `typedHello`는 string 타입이고, `m`은 MyString 타입이기 때문에 대입이 불가능하다.<br>
  대입을 하려면, 명시적으로 형변환 해줘야 한다.<br>

- **untyped constant**
  ```go
  // untyped string constant
  const hello = "Hello, 世界"

  type MyString string
  var m MyString = hello // ok!

  // untyped floating-point constant
  const Zero = 0.0

  type MyFloat64 float64
  var mf MyFloat64 = Zero // ok!
  ```
  말그대로, `타입이 정해지지 않은 상수`이다.<br>
  타입이 정해지지 않았기 때문에, 타입 에러가 생기지 않는다.<br>
  이 덕분에, `Go`에서 상수를 자유롭게 사용할 수 있다.<br>

### The largest unsigned int<br>
`uint`의 가장 큰 값을 어떻게 구할까요?<br>
만약 `uint32`라면, `const MaxUint32 = 1 << 32 - 1`로 구할 수 있다.<br>
하지만, 아키텍쳐에 따라 `uint`는 32 혹은 64비트일 수 있기 때문에 위와 같은 방법으로는 구할 수 없다.<br>
- **[2의 보수](http://en.wikipedia.org/wiki/Two's_complement)를 활용한다면?**<br>
  ```go
  const MaxUint uint = -1 // Error: negative value
  const MaxUint uint = uint(-1) // Error: negative value
  ```
  `-1`은 unsigned 범위가 아니기 때문에, 에러가 생긴다. ([constant conversion](https://go.dev/ref/spec#Conversions)의 규칙을 벗어난다.)<br>
  ```go
  var u uint
  var v = -1
  u = uint(v)
  ```
  물론, 위와 같이 변수을 사용해 런타임 환경에서 결정하게 한다면 가능하다.<br>
  하지만, 우리가 원하는 건 가장 큰 상수이다.<br>

- **해결 방법**<br>
  ```go
  const MaxUint uint = ^0 // Error: overflow

   const MaxUint = ^uint(0) // ok!
  ```
  `-1` 대신 `^0`사용해보면 어떨까? (`Go`에서는 `^`이 Not 연산자다.)<br>
  형변환을 안 하고 사용하다면, 타입이 없기 때문에 `^0`은 무한(1111.....)이 되어 버린다. 따라서 overflow 에러가 생긴다. <br>
  이러한 이유 때문에, 우리는 타입을 지정해줄 필요가 있다.<br>
  `uint(0)`을 통해 0의 타입을 지정해준다면, 컴파일이 잘 되는 걸 볼 수 있다!<br>

### go run
![image](https://github.com/user-attachments/assets/da56b6a5-bd9f-4dbe-9e5c-7d2e6fb8f142)

### Reference
[Basics/Numeric Constants](https://go.dev/tour/basics/16)<br>
[Constants](https://go.dev/blog/constants)<br>
