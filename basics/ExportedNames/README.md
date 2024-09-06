### Exported names

> [!Important]
> 이름이 대문자로 시작한다면, exported된 것이다.<br>
> (`Pi`는 `math` 패키지에서 exported된 변수다.)<br>

패키지를 `import`하면, 패키지의 `exported names`(대문자로 시작하는 이름)에만 참조할 수 있다.<br>
`unexported names`(대문자로 시작하지 않는 이름)은 패키지 밖에서 참조될 수 없다.<br>

### go run
![image](https://github.com/user-attachments/assets/cb584875-27de-44d8-8017-601e42d07780)

### Reference
[Basics/Exported names](https://go.dev/tour/basics/3)<br>
