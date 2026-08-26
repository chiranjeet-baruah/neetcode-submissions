/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	// First sort the intervals based on the starting times
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i].start < intervals[j].start
	})

	// Check for overlaps
	for i:=0;i<len(intervals)-1;i++{
		if intervals[i].end > intervals[i+1].start{
			return false
		}
	}
	return true
}