package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    var date, firstName, surname, secondName string
    var pay1, pay2, pay3 float64
    fmt.Scan(&date, &firstName, &surname, &secondName, &pay1, &pay2, &pay3)
    t, err := time.Parse("02.01.2006", date)
    if err != nil {
        fmt.Println("Error1", err)
        return
    }
    tNow := t.AddDate(0, 0, 15)
    
    totalKopeik := int(math.Round((pay1+pay2+pay3)*100))
    rub := totalKopeik / 100
    kop := totalKopeik % 100
    
    person := surname + " " + firstName + " " + secondName

    const mailTemplate = "Уважаемый, %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы. " + 
    "\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день." + 
    "\nОбщая сумма выплат составит %d руб. %d коп." + "\n\nС уважением,\nГл. бух. Иванов А.Е."

    fmt.Printf(mailTemplate, person, tNow.Format("02.01.2006"), rub, kop)

}
