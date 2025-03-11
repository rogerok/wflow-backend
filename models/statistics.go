package models

import (
	"github.com/google/uuid"
	"time"
)

type UserStatistics struct {
	UserID                uuid.UUID `json:"user_id" db:"user_id"`
	TotalWords            float64   `json:"total_words" db:"total_words"`
	TotalBooks            int       `json:"total_books" db:"total_books"`
	TotalGoals            int       `json:"total_goals" db:"total_goals"`
	CompletedGoals        int       `json:"completed_goals" db:"completed_goals"`
	TotalReports          int       `json:"total_reports" db:"total_reports"`
	AverageWordsPerDay    float64   `json:"average_words_per_day" db:"average_words_per_day"`
	AverageWordsPerReport float64   `json:"average_words_per_report" db:"average_words_per_report"`
	AverageDaysToComplete float64   `json:"average_days_to_complete" db:"average_days_to_complete"`
	MostProductiveDay     time.Time `json:"most_productive_day" db:"most_productive_day"`
	MaxWordsInDay         float64   `json:"max_words_in_day" db:"max_words_in_day"`
	CurrentStreak         int       `json:"current_streak" db:"current_streak"`
	LongestStreak         int       `json:"longest_streak" db:"longest_streak"`
	TotalDaysWithActivity int       `json:"total_days_with_activity" db:"total_days_with_activity"`
}

type GoalStatistics struct {
	GoalID                uuid.UUID  `json:"goal_id" db:"goal_id"`
	BookID                uuid.UUID  `json:"book_id" db:"book_id"`
	TotalWordsWritten     float64    `json:"total_words_written" db:"total_words_written"`
	PercentageComplete    float64    `json:"percentage_complete" db:"percentage_complete"`
	RemainingWords        float64    `json:"remaining_words" db:"remaining_words"`
	DailyWordsRequired    float64    `json:"daily_words_required" db:"daily_words_required"`
	DaysElapsed           int        `json:"days_elapsed" db:"days_elapsed"`
	DaysRemaining         int        `json:"days_remaining" db:"days_remaining"`
	AverageWordsPerDay    float64    `json:"average_words_per_day" db:"average_words_per_day"`
	ReportsCount          int        `json:"reports_count" db:"reports_count"`
	DailyProgress         []DayStats `json:"daily_progress" db:"daily_progress"`
	TrendComparedToTarget float64    `json:"trend_compared_to_target" db:"trend_compared_to_target"`
}

type DayStats struct {
	Date       time.Time `json:"date" db:"date"`
	WordCount  float64   `json:"word_count" db:"word_count"`
	ReportIDs  []string  `json:"report_ids" db:"report_ids"`
	DailyGoal  float64   `json:"daily_goal" db:"daily_goal"`
	Difference float64   `json:"difference" db:"difference"`
}
