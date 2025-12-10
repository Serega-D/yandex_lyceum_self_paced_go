package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

type Company struct {
	stuff []Worker
}

func NewCompany() *Company {
	return &Company{
		stuff: make([]Worker, 0),
	}
}

var priorityPosition = map[string]int{
	"директор":       5,
	"зам. директора": 4,
	"начальник цеха": 3,
	"мастер":         2,
	"рабочий":        1,
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if name == "" || position == "" || salary <= 0 || experience < 0 {
		return errors.New("wrong input data")
	}
	validPos := strings.TrimSpace(strings.ToLower(position))

	if _, ok := priorityPosition[validPos]; !ok {
		return errors.New("not valid position")
	}

	initialLen := len(c.stuff)

	c.stuff = append(c.stuff, Worker{name, validPos, salary, experience})

	if initialLen >= len(c.stuff) {
		return errors.New("Error: worker not added")
	}

	return nil
}

func TotalIncome(w Worker) uint {
	return w.Salary * w.Experience
}

func (c Company) SortWorkers() ([]string, error) {
	slices.SortFunc(c.stuff, func(a, b Worker) int {
		aTotal := TotalIncome(a)
		bTotal := TotalIncome(b)

		if aTotal != bTotal {
			if aTotal > bTotal {
				return -1
			}
			return 1
		}
		if priorityPosition[a.Position] != priorityPosition[b.Position] {
			if priorityPosition[a.Position] > priorityPosition[b.Position] {
				return -1
			}
			return 1
		}
		if a.Name != b.Name {
			if a.Name > b.Name {
				return -1
			}
			return 1
		}
		return 0
	})
	var output []string
	var tmp string
	for _, w := range c.stuff {
		tmp = fmt.Sprintf("%s — %d — %s", w.Name, TotalIncome(w), w.Position)
		output = append(output, tmp)
	}
	return output, nil
}

func main() {
	workers := []Worker{
		{Name: "Михаил", Position: "директор", Salary: 200, Experience: 5},
		{Name: "Игорь", Position: "зам. директора", Salary: 180, Experience: 3},
		{Name: "Николай", Position: "начальник цеха", Salary: 120, Experience: 2},
		{Name: "Андрей", Position: "мастер", Salary: 90, Experience: 10},
		{Name: "Виктор", Position: "рабочий", Salary: 80, Experience: 3},
	}
	var c Company
	c.stuff = workers
	fmt.Println(c.SortWorkers())

	// expected: []string{
	// 	"Михаил — 12000 — директор",
	// 	"Андрей — 10800 — мастер",
	// 	"Игорь — 6480 — зам. директора",
	// 	"Николай — 2880 — начальник цеха",
	// 	"Виктор — 2880 — рабочий",
	// },
}

// Напишите программу, в которой тип Company реализует такой интерфейс:

// type CompanyInterface interface{
//     AddWorkerInfo(name, position string, salary, experience uint) error
//     SortWorkers() ([]string, error)
// }

// AddWorkerInfo — метод добавления информации о новых сотрудниках, где name — имя сотрудника,
// position — должность, salary — зарплата, experience — стаж работы (месяцев).

// SortWorkers — метод, который сортирует список сотрудников по доходу за время работы на предприятии (по убыванию),
//  должности (от высокой до низкой) и возвращает слайс формата: *имя* — *доход* — *должность*.
//  Допустимые должности в порядке убывания: «директор», «зам. директора», «начальник цеха», «мастер», «рабочий».
// Примечания

// Пример отсортированного вывода:

// Михаил — 12000 — директор
// Андрей — 11800 — мастер
// Игорь — 11000 — зам. директора
