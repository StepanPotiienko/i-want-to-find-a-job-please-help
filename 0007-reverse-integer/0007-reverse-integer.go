import (
    "strconv"
    "fmt"
)

func reverse(x int) int {
    var isNegative bool = false

    if x < 0 {
        x = -x
        isNegative = true
    }

    var xStr =  strconv.Itoa(x)
    var reversed = ""
    
    for i := len(xStr) - 1; i >= 0; i-- {
        reversed += string(xStr[i])
    }

    

    if isNegative {
        reversed = "-" + reversed
    }

    fmt.Println(reversed)

    reversedInt, _ := strconv.Atoi(reversed)
    
    if reversedInt > math.MaxInt32 || reversedInt < math.MinInt32 {
        return 0
    }
    
    return reversedInt
}