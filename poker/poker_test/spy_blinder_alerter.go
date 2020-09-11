package poker

import "time"

type SpyBlindAlerter struct {
	Alerts []struct {
		ScheduledAt time.Duration
		Amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, struct {
		ScheduledAt time.Duration
		Amount      int
	}{duration, amount})
}
