package utils

import "time"

// ฟังก์ชันคำนวณเวลาที่เหลือจนถึงเที่ยงคืนของวันในหน่วยวินาที
func GetTimeMinsToNewDay() time.Duration {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	timeRemaining := endOfDay.Sub(now)

	// Convert duration to minutes
	remainingMinutes := int(timeRemaining.Minutes())
	remainingDuration := time.Duration(remainingMinutes) * time.Minute

	return remainingDuration
}

func DurationMS(start time.Time) int64 {
	timeEnd := time.Now()
	duration := timeEnd.Sub(start)
	return duration.Milliseconds()
}
