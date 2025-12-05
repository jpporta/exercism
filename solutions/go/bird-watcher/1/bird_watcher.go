package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, i := range birdsPerDay {
		sum += i
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	weekStartDay := (week - 1) * 7
	for _, i := range birdsPerDay[weekStartDay : weekStartDay+7] {
		sum += i
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for ind, _ := range birdsPerDay {
		if ind%2 == 0 {
			birdsPerDay[ind] += 1
		}
	}
	return birdsPerDay
}
