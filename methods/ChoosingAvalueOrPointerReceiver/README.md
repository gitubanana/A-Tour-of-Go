### Choosing a value or pointer receiver
`pointer receiver`를 써야하는 두 가지 이유가 있다.<br>
1. `receiver`가 가리키는 값을 변경할 수 있다.
2. `method` 호출 시, 복사본을 안 만들어도 된다.
    그래서, 더 효율적이다.

전반적으로, 모든 `method`는 `value` 혹은 `pointer` `receiver`만 가져야 한다.<br>
(둘이 섞이면 안 된다.)<br>

### go run
![image](https://github.com/user-attachments/assets/23398962-e66c-4c4f-b1a2-2542ab51faf8)


### Reference
[Methods/Choosing a value or pointer receiver](https://go.dev/tour/methods/8)<br>
