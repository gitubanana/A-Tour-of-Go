### Mutating Maps
- **원소 추가 혹은 수정**
  ```go
  m[key] = elem
  ```
- **원소 가져오기**
  ```go
  elem = m[key]
  ```
- **원소 지우기**
  ```go
  delete(m, key)
  ```
- **key가 있는지 테스트하기**
  ```go
  elem, ok := m[key]
  ```
  `key`가 `m`에 있으면, `ok`는 `true` 아니면, `false`<br>
  `key`가 `m`에 없으면, `elem`은 `zero value`가 된다.<br>

### go run
![image](https://github.com/user-attachments/assets/b71fca7a-634a-4cc8-a451-82a431467757)


### Reference
[Moretypes/Mutating Maps](https://go.dev/tour/moretypes/22)<br>
