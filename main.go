package main

import "fmt"

func main() {
	var water, milk, beans, cups, money int
	initializeState(&water, &milk, &beans, &cups, &money)

	var action string
	for action != "exit" {
		getAction(&action)

		switch action {
		case "buy":
			buy(&water, &milk, &beans, &cups, &money)
		case "fill":
			fill(&water, &milk, &beans, &cups)
		case "take":
			take(&money)
		case "remaining":
			showState(water, milk, beans, cups, money)
		}
		fmt.Println()
	}
}

func getAction(s *string) {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	_, err := fmt.Scanln(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

func buy(water, milk, beans, cups, money *int) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var coffee string
	_, err := fmt.Scanln(&coffee)
	if err != nil {
		fmt.Println(err)
	}
	switch coffee {
	case "1":
		notEnough, ok := checkResources("espresso", *water, *milk, *beans, *cups)
		if !ok {
			fmt.Printf("Sorry, not enough %s!\n", notEnough)
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*water -= 250
			*beans -= 16
			*cups -= 1
			*money += 4
		}
	case "2":
		notEnough, ok := checkResources("latte", *water, *milk, *beans, *cups)
		if !ok {
			fmt.Printf("Sorry, not enough %s!\n", notEnough)
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*water -= 350
			*milk -= 75
			*beans -= 20
			*cups -= 1
			*money += 7
		}
	case "3":
		notEnough, ok := checkResources("cappuccino", *water, *milk, *beans, *cups)
		if !ok {
			fmt.Printf("Sorry, not enough %s!\n", notEnough)
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*water -= 200
			*milk -= 100
			*beans -= 12
			*cups -= 1
			*money += 6
		}
	}
}

func checkResources(coffee string, water, milk, beans, cups int) (resource string, ok bool) {
	switch coffee {
	case "espresso":
		if water < 250 {
			resource, ok = "water", false
		} else if beans < 16 {
			resource, ok = "beans", false
		} else if cups < 1 {
			resource, ok = "cups", false
		} else {
			resource, ok = "", true
		}
	case "latte":
		if water < 350 {
			resource, ok = "water", false
		} else if milk < 75 {
			resource, ok = "milk", false
		} else if beans < 20 {
			resource, ok = "beans", false
		} else if cups < 1 {
			resource, ok = "cups", false
		} else {
			resource, ok = "", true
		}
	case "cappuccino":
		if water < 200 {
			resource, ok = "water", false
		} else if milk < 100 {
			resource, ok = "milk", false
		} else if beans < 12 {
			resource, ok = "beans", false
		} else if cups < 1 {
			resource, ok = "cups", false
		} else {
			resource, ok = "", true
		}
	}
	return
}

func fill(water, milk, beans, cups *int) {
	fmt.Println("Write how many ml of water you want to add:")
	*water += getAmount()
	fmt.Println("Write how many ml of milk you want to add:")
	*milk += getAmount()
	fmt.Println("Write how many grams of coffee beans you want to add:")
	*beans += getAmount()
	fmt.Println("Write how many disposable cups you want to add:")
	*cups += getAmount()
}

func getAmount() (amount int) {
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func take(money *int) {
	fmt.Printf("I gave you $%d\n", *money)
	*money = 0
}

func showState(water int, milk int, beans int, cups int, money int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", beans)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n", money)
}

func initializeState(water, milk, beans, cups, money *int) {
	*water = 400
	*milk = 540
	*beans = 120
	*cups = 9
	*money = 550
}
