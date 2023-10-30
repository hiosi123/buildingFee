package utils

import (
	"fmt"
	"time"
)

var YYYY = "2006"

var YYYYMMDD = "20060102"

var YYYYMMDDhhmmss = "20060102150405"

var YYYYMMDDhhmmssfff = "20060102150405.000"

func GetTodayYear() string {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		// Handle error
		location, err = time.LoadLocation("UTC")
		if err != nil {
			fmt.Printf("failed to load time zone: %v", err)
			return ""
		}
		now := time.Now().In(location).Add(9 * time.Hour)
		return now.Format(YYYY)
	}

	now := time.Now().In(location)
	return now.Format(YYYY)
}

func GetTodayDate() string {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		// Handle error
		location, err = time.LoadLocation("UTC")
		if err != nil {
			fmt.Printf("failed to load time zone: %v", err)
			return ""
		}
		now := time.Now().In(location).Add(9 * time.Hour)
		return now.Format(YYYYMMDD)
	}

	now := time.Now().In(location)
	return now.Format(YYYYMMDD)
}

func GetCurrentTime() string {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		location, err = time.LoadLocation("UTC")
		if err != nil {
			fmt.Printf("failed to load time zone: %v", err)
			return ""
		}
		now := time.Now().In(location).Add(9 * time.Hour)
		return now.Format(YYYYMMDDhhmmss)

	}

	now := time.Now().In(location)
	return now.Format(YYYYMMDDhhmmss)
}

func GetCurrentTimeForFileUpload() string {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		location, err = time.LoadLocation("UTC")
		if err != nil {
			fmt.Printf("failed to load time zone: %v", err)
			return ""
		}
		now := time.Now().In(location).Add(9 * time.Hour)
		return now.Format(YYYYMMDDhhmmssfff)

	}

	now := time.Now().In(location)
	return now.Format(YYYYMMDDhhmmssfff)
}

func GetEpochUnixTime() int64 {
	return time.Now().Unix()
}
