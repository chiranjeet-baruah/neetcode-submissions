func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    var result [][]int
    result = append(result, intervals[0])
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] > result[len(result)-1][1] {
            result = append(result, intervals[i])
        } else {
            result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
        }
    } 
    
    return result
}