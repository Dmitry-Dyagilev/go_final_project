package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const format = "20060102"

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(format, nowStr)
		if err != nil {
			http.Error(w, "неверный формат даты", http.StatusBadRequest)
			return
		}
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	req, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(req))
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(format, dstart)
	if err != nil {
		return "", err
	}

	interval, err := validateRepeat(repeat)
	if err != nil {
		return "", err
	}
	switch {
	case repeat == "y":
		for {
			date = date.AddDate(interval, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	case strings.HasPrefix(repeat, "d "):
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
	}
	return date.Format(format), nil
}

// проверка что будующая дата больше текущей
func afterNow(date, now time.Time) bool {
	return date.Format(format) > now.Format(format)
}

// проверка валидности формата повторения для d, y
func validateRepeat(repeat string) (int, error) {
	if repeat == "" {
		return 0, fmt.Errorf("поле не может быть пустым")
	}
	switch {
	case repeat == "y":
		return 1, nil
	case strings.HasPrefix(repeat, "d "):
		return validateDaily(repeat)
	//дальше возможно появятся функции для проверки w и m
	default:
		return 0, fmt.Errorf("неподдерживаемый формат: %s", repeat)
	}
}

// проверка валидности формата повторения (числа дней)
func validateDaily(repeat string) (int, error) {
	parts := strings.Split(repeat, " ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("ожидается d <число>")
	}

	num, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("символ не является числом: %w", err)
	}
	if num < 1 || num > 400 {
		return 0, fmt.Errorf("число должно быть от 1 до 400")
	}
	return num, nil
}
