package db

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {

	res, err := DB.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func Tasks(limit int) ([]*Task, error) {
	rows, err := DB.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date > date('now') ORDER BY date LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	row := DB.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id)
	var task Task
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *Task) error {
	res, err := DB.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("id", task.ID),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("неверный идентификатор для обновления задачи %s", task.ID)
	}
	return nil
}

func DeleteTask(id string) error {
	res, err := DB.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("задача не найдена")
	}
	return nil
}

func UpdateDate(next string, id string) error {
	res, err := DB.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
		sql.Named("id", id),
		sql.Named("date", next))
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("неверный идентификатор для обновления даты ")
	}
	return nil
}
