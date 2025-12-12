package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxSize = 5

func printQueue (queue []string) {
    for i, name := range queue {
        fmt.Printf("%d. %s\n", i+1, name)
    }
}

func printStats (occupied int) {
    fmt.Printf("Осталось свободных мест: %d\n", maxSize-occupied)
    fmt.Printf("Всего человек в очереди: %d\n", occupied)
}


func main () {
    queue := []string{"-", "-", "-", "-", "-"}
    occupied := 0

    scanner := bufio.NewScanner(os.Stdin)

    for {
        if !scanner.Scan() {
            break
        }
    
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }

        parts := strings.Fields(line)
        
        if len(parts) == 1 {
            cmd := parts[0]

            switch cmd {
            case "очередь": 
                printQueue(queue)
                continue

            case "количество": 
                printStats (occupied)
                continue

            case "конец":
                printQueue(queue)
                return

            default: fmt.Println("некорректный ввод")

            }
        }

        if len(parts) != 2 {
            fmt.Println("некорректный ввод")
        }
        name := parts[0]
        placeStr := parts[1]

        place, err := strconv.Atoi(placeStr)
        if err != nil {
            fmt.Println("некорректный ввод")
            continue
        }

        if place < 1 || place > 5 {
            fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
            continue
        }

        if occupied == maxSize {
            fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", place)
            continue
        }

        if queue[place-1] != "-" {
            fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
            continue
        }
        
        queue[place-1] = name 
        occupied++
    }
}
