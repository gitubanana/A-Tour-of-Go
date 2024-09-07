### Functions continued
연속된 매개변수가 같은 타입이라면, 마지막 타입 빼고 다 생략할 수 있다.<br>
`C`에서도 매개변수일 때 생략이 불가능했지만, `Go`에서는 생략이 가능하다.<br>
- **C**
  ```c++
  // 생략 못 함
  int add(int x, int y)
  {
    return x + y;
  }
  ```
- **Go**
  ```go
  // 생략 안 한 함수
  func add(x int, y int) int {
    return x + y
  }
  
  // 생략 한 함수
  func add (x, y int) int {
    return x + y
  }
  ```

### go run
![image](https://github.com/user-attachments/assets/697d79e2-6141-4a58-81c9-b33f6a346237)


### Reference
[Basics/Functions continued](https://go.dev/tour/basics/5)<br>
