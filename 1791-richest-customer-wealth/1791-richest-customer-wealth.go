import "fmt"

func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    wealth := 0

    for subarray := range accounts {
        for element := range accounts[subarray] {
            wealth += accounts[subarray][element]
        }

        if (wealth > maxWealth) {
            maxWealth = wealth
        }

        wealth = 0
    }

    return maxWealth
}