package main

import (
	"fmt"
	"os"
	"strings"
)

var data = []map[string]interface{}{}

var user1 = map[string]interface{}{
	"account balance": 100,
	"email":           "as",
	"subscription":    "Gold",
}

var user2 = map[string]interface{}{
	"account balance": 200,
	"email":           "qs",
	// "subscription": ""
}

func check(channelAmount int, bal int) bool {
	if bal < channelAmount {
		return true
	}
	return false
}

var input int

func enter() {
	fmt.Print("Enter the option: ")
	fmt.Scanln(&input)
}

func main() {
	data = append(data, user1, user2)
	fmt.Print("Enter email :")
	var id string
	fmt.Scanln(&id)

	enter()

	for _, p := range data {
		//for k := range p {
		if p["email"] == id {
			//Main code
			balance := p["account balance"].(int) //interface type to int type
			var subs string
			if p["subscription"] != nil {
				subs = fmt.Sprintf("%v", p["subscription"]) //type conversion from interface to string
			} else {
				subs = "Not subscribed"
			}

			for input <= 5 {

				switch input {
				case 1:
					fmt.Printf("Account balance is %d Rs.\n", balance)
					fmt.Printf("Current subscription : %s \n", subs)
					enter()

				case 2:
				repeat:
					fmt.Print("Enter the amount to recharge: ")
					var recharge int
					fmt.Scanln(&recharge)
					if recharge > 0 {
						balance = balance + recharge
						fmt.Printf("Recharge completed successfully. Current balance is %d \n", balance)
					} else {
						fmt.Println("Enter Valid Recharge Amount.")
						goto repeat
					}
					enter()

				case 3:
					fmt.Println("Available packs for subscription:\nSilver pack:{Zee, Sony, Star Plus: 50 Rs.}\nGold Pack:{Zee, Sony, Star Plus, Discovery, NatGeo: 100 Rs.}")
					fmt.Println("Available channels for subscription\n{ Zee: 10 Rs, Sony: 15 Rs, Star Plus: 20 Rs, Discovery: 10 Rs, NatGeo: 20 Rs.}")
					enter()

				case 4:
				loop:
					fmt.Print("Enter the Pack you wish to subscribe: (Silver: ‘S’, Gold: ‘G’): ")
					var pack string
					fmt.Scanln(&pack)

					var price int

					switch pack {
					case "G":
						subs = "Gold"
						price = 100
					case "S":
						subs = "Silver"
						price = 50
					default:
						fmt.Println("Enter Valid Pack 'S' or 'D'")
						goto loop
					}

					fmt.Print("Enter the months: ")
					var months int
					fmt.Scanln(&months)

					fmt.Printf("\nYou have successfully subscribed the following pack - %s", subs)
					fmt.Printf("\nMonthly price: %d Rs.", price)
					fmt.Printf("\nNo of months: %d", months)
					subsAmount := price * months
					fmt.Printf("\nSubscription Amount: %d Rs.", subsAmount)

					if months >= 3 {
						discount := subsAmount * 10 / 100
						subsAmount = subsAmount - discount
						fmt.Printf("\nDiscount applied: %d Rs.", discount)
						fmt.Printf("\nFinal Price after discount: %d Rs", subsAmount)
					}
					balance = balance - subsAmount
					fmt.Println("\nEmail notification sent successfully\nSMS notification sent successfully")
					enter()

				case 5:
				call:
					fmt.Print("Enter channel names to add (separated by commas): ")
					var channel string
					fmt.Scanln(&channel)

					result := strings.Split(channel, ",") // Split the string based on delimiter ","
					var channelPrice int
					for i := range result {
						switch result[i] {
						case "Zee":
							channelPrice = 10
							checkBalance := check(channelPrice, balance)
							if checkBalance == true {
								fmt.Printf("Can not add %s , Low Balance\n", result[i])
								break
							}
							balance = balance - channelPrice

						case "Sony":
							channelPrice = 15
							checkBalance := check(channelPrice, balance)
							if checkBalance == true {
								fmt.Printf("Can not add %s , Low Balance\n", result[i])
								break
							}
							balance = balance - channelPrice

						case "Star Plus":
							channelPrice = 20
							checkBalance := check(channelPrice, balance)
							if checkBalance == true {
								fmt.Printf("Can not add %s , Low Balance\n", result[i])
								break
							}
							balance = balance - channelPrice

						case "Discovery":
							channelPrice = 10
							checkBalance := check(channelPrice, balance)
							if checkBalance == true {
								fmt.Printf("Can not add %s , Low Balance\n", result[i])
								break
							}
							check(channelPrice, balance)
							balance = balance - channelPrice

						case "NatGeo":
							channelPrice = 20
							checkBalance := check(channelPrice, balance)
							if checkBalance == true {
								fmt.Printf("Can not add %s , Low Balance\n", result[i])
								break
							}
							balance = balance - channelPrice

						default:
							fmt.Println("Enter Valid Channel Name.")
							goto call

						}

					}
					fmt.Println("Channels added successfully.")
					fmt.Printf("Account balance: %d Rs. \n", balance)
					enter()

				case 6:
					os.Exit(0)
				}
			}

		}

		//}
	}

}
