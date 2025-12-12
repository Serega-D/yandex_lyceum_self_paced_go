package main

import (
	"fmt"
	"math"
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	p.Rating = float64(p.Goals) + float64(p.Assists)/2.0
	if p.Misses > 0 {
		p.Rating /= float64(p.Misses)
	}
	p.Rating = math.Round(p.Rating*10) / 10
}

func NewPlayer(name string, goals, misses, assists int) Player {
	newP := Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: 0.0}
	newP.calculateRating()
	return newP
}

func goalsSort(players []Player) []Player {
	sortedPlayers := make([]Player, len(players))
	copy(sortedPlayers, players)
	slices.SortFunc(sortedPlayers, func(a, b Player) int {
		if a.Goals != b.Goals {
			if a.Goals > b.Goals {
				return -1
			}
			return 1
		}
		if a.Name < b.Name {
			return -1
		}
		return 0
	})
	return sortedPlayers
}

func ratingSort(players []Player) []Player {
	sortedPlayers := make([]Player, len(players))
	copy(sortedPlayers, players)
	slices.SortFunc(sortedPlayers, func(a, b Player) int {
		if a.Rating != b.Rating {
			if int(a.Rating) > int(b.Rating) {
				return -1
			}
			return 1
		}
		if a.Name < b.Name {
			return -1
		}
		return 0
	})
	return sortedPlayers
}

func gmSort(players []Player) []Player {
	sortedPlayers := make([]Player, len(players))
	copy(sortedPlayers, players)
	slices.SortFunc(sortedPlayers, func(a, b Player) int {
		aGM, bGM := float64(a.Goals), float64(b.Goals)
		if a.Misses > 0 {
			aGM = float64(a.Goals) / float64(a.Misses)
		}
		if b.Misses > 0 {
			bGM = float64(b.Goals) / float64(b.Misses)
		}
		if aGM != bGM {
			if aGM > bGM {
				return -1
			}
			return 1
		}
		if a.Name < b.Name {
			return -1
		}
		return 0
	})
	return sortedPlayers
}

func main() {

	p := NewPlayer("Player2", 15, 7, 2)
	fmt.Println(p.Rating)

	// name:           "No misses",
	// 	player:         NewPlayer("Player1", 10, 0, 4),
	// 	expectedRating: 12.0,
	// },
	// {
	// 	name:           "Normal case",
	// 	player:         NewPlayer("Player2", 8, 2, 4),
	// 	expectedRating: 5.0,
	// },
	// {
	// 	name:           "Zero assists",
	// 	player:         NewPlayer("Player3", 5, 1, 0),
	// 	expectedRating: 5.0,
	// },
	// {
	// ewPlayer("Player1", 10, 5, 3),
	// 		NewPlayer("Player2", 15, 7, 2),
	// 		NewPlayer("Player3", 8, 2, 5),

}

//     Хранить данные о каждом футболисте: Имя, Голы, Промахи и Помощь.
//     Рассчитывать рейтинг для каждого игрока по формуле: (Голы + Помощь / 2) / Промахи
// 	(если количество промахов равно нулю, то Голы + Помощь / 2).
//     Сортировка по:
//         Убыванию количества голов (функция goalsSort(players []Player) []Player)
//         Убыванию рейтинга (функция ratingSort(players []Player) []Player)
//         Убыванию отношения голов к промахам (функция gmSort(players []Player) []Player)

// 		Если рейтинг одинаковый, то во всех функциях необходимо сортировать по имени (Name) в алфавитном порядке.
//     Также нужно реализовать такую структуру:

// и вспомогательный метод calculateRating() для расчёта поля Rating на основе входных данных.
// Конструктор NewPlayer(name string, goals, misses, assists int) Player создаёт нового игрока и
// вычисляет его рейтинг с помощью calculateRating().
