package main

import "fmt"

type Animal interface {
    MakeSound() string
    GetName() string
    GetInfo() string
}

type animal struct {
    name string
    species string
    age int
    sound string
}

func (a animal) MakeSound() string {return a.sound}
func (a animal) GetName() string {return a.name}
func (a animal) GetInfo() string {
    return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}
func NewAnimal(name, species string, age int, sound string) *animal {
    a := animal{name: name, species: species, age: age, sound: sound}
    return &a
}

func ZooShow(animals []Animal) {
    for _, animal := range animals {
        fmt.Println(animal.GetInfo())
        fmt.Println(animal.MakeSound())
    }
}

type ZooKeeper struct {}

func (z ZooKeeper) Feed(animal Animal) {
    fmt.Printf("Смотритель зоопарка кормит %s. %s!", animal.GetName(), animal.MakeSound())
}

