### Type conversions
`T(v)`는 값 `v`를 타입 `T`로 변환한다.<br>
`C`와는 다르게, `Go`에서는 명시적 형변환만 있다. (암시적 형변환은 없다.)<br>
- **C**
  ```c++
  int age = 0;
  long long eConversion = (long long)age; // 명시적 형변환 가능
  long long iConversion = age; // 암시적 형변환 가능
  ```
- **Go**
  ```go
  var age int = 0
  var eConversion int64 = int64(age) // 명시적 형변환 가능
  
  // cannot use age (variable of type int) as int64 value in variable declaration
  var iConversion int64 = age // 암시적 형변환 불가능
  ```

### go run
![image](https://github.com/user-attachments/assets/722558b3-333f-4b96-a598-5e777d0e7b71)

### Reference
[Basics/Type conversions](https://go.dev/tour/basics/13)<br>
