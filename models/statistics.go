package models

import (
	"github.com/google/uuid"
	"time"
)

type UserStatistics struct {
	UserID                uuid.UUID `json:"userId" db:"user_id"`
	TotalWords            float64   `json:"totalWords" db:"total_words"`
	TotalBooks            int       `json:"totalBooks" db:"total_books"`
	TotalGoals            int       `json:"totalGoals" db:"total_goals"`
	CompletedGoals        int       `json:"completedGoals" db:"completed_goals"`
	TotalReports          int       `json:"totalReports" db:"total_reports"`
	AverageWordsPerDay    float64   `json:"averageWordsPerDay" db:"average_words_per_day"`
	AverageWordsPerReport float64   `json:"averageWordsPerReport" db:"average_words_per_report"`
	AverageDaysToComplete float64   `json:"averageDaysToComplete" db:"average_days_to_complete"`
	MostProductiveDay     time.Time `json:"mostProductiveDay" db:"most_productive_day"`
	MaxWordsInDay         float64   `json:"maxWordsInDay" db:"max_words_in_day"`
	CurrentStreak         int       `json:"currentStreak" db:"current_streak"`
	LongestStreak         int       `json:"longestStreak" db:"longest_streak"`
	TotalDaysWithActivity int       `json:"totalDaysWithActivity" db:"total_days_with_activity"`
}

type GoalStatistics struct {
	GoalID                uuid.UUID `json:"goalId" db:"goal_id"`
	BookID                uuid.UUID `json:"bookId" db:"book_id"`
	TotalWordsWritten     float64   `json:"totalWordsWritten" db:"total_words_written"`
	PercentageComplete    float64   `json:"percentageComplete" db:"percentage_complete"`
	RemainingWords        float64   `json:"remainingWords" db:"remaining_words"`
	DailyWordsRequired    float64   `json:"dailyWordsRequired" db:"daily_words_required"`
	DaysElapsed           int       `json:"daysElapsed" db:"days_elapsed"`
	DaysRemaining         int       `json:"daysRemaining" db:"days_remaining"`
	AverageWordsPerDay    float64   `json:"averageWordsPerDay" db:"average_words_per_day"`
	ReportsCount          int       `json:"reportsCount" db:"reports_count"`
	TrendComparedToTarget float64   `json:"trendComparedToTarget" db:"trend_compared_to_target"`
}

type GoalsChart struct {
	GoalID                uuid.UUID `json:"goalId" db:"goal_id"`
	BookID                uuid.UUID `json:"bookId" db:"book_id"`
	GoalTitle             string    `json:"goalTitle" db:"goal_title"`
	TotalWordsWritten     float64   `json:"totalWordsWritten" db:"total_words_written"`
	PercentageComplete    float64   `json:"percentageComplete" db:"percentage_complete"`
	RemainingWords        float64   `json:"remainingWords" db:"remaining_words"`
	DailyWordsRequired    float64   `json:"dailyWordsRequired" db:"daily_words_required"`
	DaysElapsed           int       `json:"daysElapsed" db:"days_elapsed"`
	DaysRemaining         int       `json:"daysRemaining" db:"days_remaining"`
	AverageWordsPerDay    float64   `json:"averageWordsPerDay" db:"average_words_per_day"`
	ReportsCount          int       `json:"reportsCount" db:"reports_count"`
	TrendComparedToTarget float64   `json:"trendComparedToTarget" db:"trend_compared_to_target"`
	BookTitle             string    `json:"bookTitle" db:"book_title"`
	CreatedAt             string    `json:"createdAt" db:"created_at"`
	IsFinished            bool      `json:"isFinished" db:"is_finished"`
	IsExpired             bool      `json:"isExpired" db:"is_expired"`
}

type ReportStatistic struct {
	BookId      string  `json:"bookId" db:"book_id"`
	BookName    string  `json:"bookName" db:"book_name"`
	CreatedAt   string  `json:"createdAt" db:"created_at"`
	GoalTitle   string  `json:"goalTitle" db:"goal_title"`
	GoalId      string  `json:"goalId" db:"goal_id"`
	Id          string  `json:"id" db:"id"`
	UpdatedAt   string  `json:"updatedAt" db:"updated_at"`
	WordsAmount float64 `json:"wordsAmount" db:"words_amount"`
	GoalWords   float64 `json:"goalWords" db:"goal_words"`
}

// ChartData represents data for various frontend charts
type ChartData struct {
	DailyProgress        []DailyProgressPoint `json:"daily_progress"`
	CumulativeProgress   []ProgressPoint      `json:"cumulative_progress"`
	MonthlyComparison    []MonthlyStats       `json:"monthly_comparison"`
	GoalCompletion       []GoalCompletionData `json:"goal_completion"`
	ProductivityPatterns ProductivityPatterns `json:"productivity_patterns"`
}

// DailyProgressPoint represents a single data point for daily progress
type DailyProgressPoint struct {
	Date          time.Time `json:"date"`
	WordCount     float64   `json:"word_count"`
	TargetCount   float64   `json:"target_count"`
	MovingAverage float64   `json:"moving_average"`
}

// ProgressPoint represents progress data over time
type ProgressPoint struct {
	Date              string  `json:"date"`
	TotalWords        float64 `json:"totalWords"`
	TargetTotalWords  float64 `json:"targetTotalWords"`
	CompletionPercent float64 `json:"completionPercent"`
	GoalTitle         string  `json:"goalTitle"`
	BookName          string  `json:"bookName"`
	BookId            string  `json:"bookId"`
	GoalId            string  `json:"goalId"`
}

// MonthlyStats represents writing statistics for a specific month
type MonthlyStats struct {
	Month        time.Time `json:"month"`
	TotalWords   float64   `json:"total_words"`
	DailyAverage float64   `json:"daily_average"`
	ActiveDays   int       `json:"active_days"`
}

// GoalCompletionData represents data about goal completion
type GoalCompletionData struct {
	GoalID         string    `json:"goal_id"`
	GoalTitle      string    `json:"goal_title"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	TargetWords    float64   `json:"target_words"`
	CompletedWords float64   `json:"completed_words"`
	CompletionRate float64   `json:"completion_rate"`
	Status         string    `json:"status"` // "completed", "in_progress", "expired"
}

// ProductivityPatterns represents patterns in user productivity
type ProductivityPatterns struct {
	DayOfWeek []DayOfWeekStats `json:"day_of_week"`
	TimeOfDay []TimeOfDayStats `json:"time_of_day"`
}

// DayOfWeekStats represents writing statistics for each day of the week
type DayOfWeekStats struct {
	Day       string  `json:"day"` // Monday, Tuesday, etc.
	WordCount float64 `json:"word_count"`
}

// TimeOfDayStats represents writing statistics for different times of day
type TimeOfDayStats struct {
	TimeSlot  string  `json:"time_slot"` // Morning, Afternoon, Evening, Night
	WordCount float64 `json:"word_count"`
}

type FullProfileChartData struct {
	//DailyProgress      []DailyProgressPoint `json:"daily_progress"`
	CumulativeProgress []ProgressPoint `json:"cumulativeProgress"`
	//MonthlyComparison  []MonthlyStats       `json:"monthly_comparison"`
	GoalCompletion []GoalsChart `json:"goalCompletion"`
	//ProductivityPatterns *ProductivityPatterns `json:"productivity_patterns"`
}
