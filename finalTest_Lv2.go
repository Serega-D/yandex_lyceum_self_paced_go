package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
    "unicode/utf8"
)

func currentDayOfTheWeek() string {
    t := time.Now()
    daysRus := []string{
        "Четверг",
        "Пятница",
        "Суббота",
        "Воскресенье",
        "Понедельник",
        "Вторник",
        "Среда",   
    }
    return fmt.Sprint(daysRus[t.Weekday()])
}

func dayOrNight() string {
    t := time.Now()
    hour := t.Hour()
    if hour >= 10 && hour <= 22{
        return "День"
    }
    return "Ночь"
}

func nextFriday() int {
    t := time.Now()
    return int(1 - t.Weekday() + 7) % 7 
}

func CheckCurrentDayOfTheWeek(answer string) bool {
    return strings.ToLower(answer) == strings.ToLower(currentDayOfTheWeek())
}

func CheckNowDayOrNight(answer string) (bool, error) {
    if utf8.RuneCountInString(answer) != 4 {
        return false, errors.New("исправь свой ответ, а лучше ложись поспать")
    }
    if strings.ToLower(answer) == strings.ToLower(dayOrNight()) {
        return true, nil
    }
    return false, nil
}
