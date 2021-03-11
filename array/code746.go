package array

func minCostClimbingStairs(cost []int) int {

    var min int

    // dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
    dp0, dp1 := cost[0], cost[1]
    for i := 2; i < len(cost); i++ {
        dp := cost[i]
        if dp0 < dp1 {
            dp += dp0
        } else {
            dp += dp1
        }
        dp0, dp1 = dp1, dp
    }

    if dp0 < dp1 {
        min = dp0
    } else {
        min = dp1
    }

    return min
}
