### Imports
> [!Important]
> ()로 묶인 import문을 `factored import`문라고 한다.<br>
> 물론, ()없이 여러 줄로 import문을 사용할 수 있지만, `factored import`문이 더 좋은 스타일이다.<br>
```go
// factored import statement
import (
  "fmt"
  "math"
)

// multiple import statement
import "fmt"
import "math"
```

### go run
![image](https://github.com/user-attachments/assets/8acfa82a-ab07-4e6e-81bf-186aee52967e)

### Reference
[Basics/Imports](https://go.dev/tour/basics/2)<br>
