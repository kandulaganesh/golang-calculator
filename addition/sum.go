package addition

import "fmt"

func Addition(value1, value2 int) int {
    fmt.Printf("Sum of %d and %d is %d\n", value1, value2, value1 + value2)
    return value1 + value2
}
