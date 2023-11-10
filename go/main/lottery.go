package main

// Перехрест АВ проверил

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const MaxPoint = 10

var length int = 6

func main() {
	var response string
	var sum int

	for true {
		fmt.Print("Купить лотерейный билет? (Да/Нет) ")
		fmt.Scanf("%s\n", &response)

		if response == "Да" || response == "да" || response == "" {
			initialTicket := get_initial_lottery_ticket()
			lotteryTicket := get_lottery_ticket()

			fmt.Println("\nБилет куплен")
			fmt.Println(initialTicket)

			for position := 1; position < length+1; position++ {
				fmt.Printf("\nСтереть позицию %d? (Да/Нет) ", position)
				fmt.Scanf("%s\n", &response)

				if response == "Да" || response == "да" || response == "" {
					initialTicket[position-1] = strconv.Itoa(lotteryTicket[position-1])
					sum += lotteryTicket[position-1]
					fmt.Println(initialTicket)
				} else {
					sum = 0
					return
				}
			}

			fmt.Printf("\nСумма очков равна: %d\n", sum)

			if sum == 21 {
				sum = 0
				fmt.Println("Поздравляю! Вы выиграли автомобиль!!!")
				return
			} else {
				sum = 0
				fmt.Println("К сожалению вы проиграли :(")
			}
		} else {
			return
		}
	}

}

func get_initial_lottery_ticket() []string {
	initialLotteryTicket := make([]string, length)

	for i := 0; i < length; i++ {
		initialLotteryTicket[i] = "*"
	}
	return initialLotteryTicket
}

func get_lottery_ticket() []int {
	numbers := make([]int, length)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(MaxPoint)
	}
	return numbers
}
