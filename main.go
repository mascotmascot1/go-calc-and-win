package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// rnd - генератор псевдослучайных чисел
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// randNum возвращает случайное число в интервале [min, max]
func randNum(min, max int) int {
	return rnd.Intn(max-min+1) + min
}

// input запрашивает и возвращает ввод пользователя в консоли
func input(title string) string {
	fmt.Print(title)
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println(err)
	}
	return s
}

func setEnemyHealth() int {
	return randNum(80, 120)
}

func getLiteAttack() int {
	return randNum(2, 5)
}

func getMidAttack() int {
	return randNum(15, 25)
}

func getHardAttack() int {
	return randNum(30, 40)
}

func compareValues(enemyHealth, userTotalAttack int) bool {
	pointDifference := enemyHealth - userTotalAttack
	if pointDifference < 0 {
		pointDifference = -pointDifference
	}
	return pointDifference <= 10
}

func getUserAttack() int {
	total := 0

	for i := 0; i < 5; i++ {
		inputAttack := input("Введи тип атаки: ")

		var attackValue int
		switch inputAttack {
		case "lite":
			attackValue = getLiteAttack()
		case "mid":
			attackValue = getMidAttack()
		case "hard":
			attackValue = getHardAttack()
		default:
			fmt.Println("Неизвестный тип атаки:", inputAttack)
			i--
			continue
		}
		fmt.Println("Количество очков твоей атаки:", attackValue)
		total += attackValue
	}
	return total
}

func runGame() bool {
	enemyHealth := setEnemyHealth()
	userTotalAttack := getUserAttack()
	fmt.Println("Тобой нанесён урон противнику равный", userTotalAttack)
	fmt.Println("Очки здоровья противника до твоей атаки", enemyHealth)
	if compareValues(enemyHealth, userTotalAttack) {
		fmt.Println("Ура! Победа за тобой!")
	} else {
		fmt.Println("В этот раз не повезло :( Бой проигран.")
	}
	answer := input("Чтобы сыграть ещё раз, введи букву [y] или [Y]: ")
	return strings.ToUpper(answer) == "Y"
}

func main() {
	intro := `РАССЧИТАЙ И ПОБЕДИ!
Загрузка...
	
Твоя цель — за 5 ходов набрать такое количество очков урона противнику,
которое попадет в диапазон +– 10 от значения здоровья противника.
	
Значение здоровья противника генерируется случайным образом
в диапазоне от 80 до 120 очков.
	
В твоём распоряжении три вида атак:
lite — урон от 2 до 5 очков;
mid — урон от 15 до 25 очков;
hard — урон от 30 до 40 очков.
ВПЕРЁД К ПОБЕДЕ!!!
`
	fmt.Println(intro)

	for runGame() {
	}
}
