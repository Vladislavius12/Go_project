package main
import (
	"fmt"
	"math/rand"
)

const  count = 3

type Character interface {
	Attack(character Character)
	Defend()
	ability1(character Character)
	ability2(character Character)
	Health() int
	StartHealth(int)  //Начальное количество жизней
	StartDefense(int) //Начальное количество защиты

}

type Warrior struct {
	Healths int //Здоровье
	Strength int //Сила
	Defense int //Защита
	Confidence int //Уверенность в себе
}

type Shaman struct {
	Healths int
	Strength int
	Defense int
	Ether int //Эфир
}

type King struct {
	Healths int
	Strength int
	Defense int
	Fear int //Колличество страха перед противником
}

func (war *Warrior) Attack(character Character) { //Функция атаки Воина
	war.Confidence += 5
	damage := war.Strength + (0 + rand.Intn(15-0+1)) - war.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)
}
func (war *Warrior) Defend() { //Функция защиты Воина
	war.Confidence += 5
	war.Defense = war.Defense + (0 + rand.Intn(15-0+1))
}
func (war *Warrior) ability1(character Character) { //Первая Абилка Воина
	if war.Confidence < 10 {
		fmt.Println("You're too insecure about yourself")
		return
	}
	fmt.Println("You pounce on the opponent shouting «For the King» ")
	war.Strength += 1
	war.Confidence -= 5
	damage := war.Strength + 5 + (0 + rand.Intn(15-0+1)) - war.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)
}
func (war *Warrior) ability2 (character Character) {
	if war.Confidence < 20 {
		fmt.Println("I can't shoot on a sober head!")
		return
	}
	fmt.Println("You take a revolver out of your pocket, your confidence grows. Now no one can stop you")
	war.Strength += 5
	war.Confidence -= 15
	damage := war.Strength + 15 + (0 + rand.Intn(15-0+1)) - war.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)

}
func (war *Warrior) Health() int {
	return war.Healths
}
func (war *Warrior) StartHealth(points int) {
	war.Healths = war.Healths - points
}
func (war *Warrior) StartDefense(point int) {
	war.Defense = war.Defense - point
}

func (sh *Shaman) Attack(character Character) {
	sh.Ether += 5
	damage := sh.Strength + 7 + (0 + rand.Intn(15-0+1)) - sh.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)

}
func (sh *Shaman) Defend() {
	sh.Ether += 15
	sh.Defense = sh.Defense + (0 + rand.Intn(15-0+1))
}
func (sh *Shaman) ability1(character Character) {
	if sh.Ether < 10 {
		fmt.Println("The subscriber is outside the network coverage area. I guess I'll have to wait")
		return
	}
	fmt.Println("The old cane is still good for something. You strike with the force of a nuclear warhead")
	sh.Strength += 20
	sh.Ether -= 10
	damage := sh.Strength + (0 + rand.Intn(5-0+1)) - sh.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)

}
func (sh *Shaman) ability2(character Character) {
	if sh.Ether < 20 {
		fmt.Println("I think the battery is dead. We'll have to look for an outlet")
		return
	}
	fmt.Println("I hit two times: one on the head, the second on the lid of the coffin!")
	sh.Strength -= 20
	sh.Ether += 20
	damage := sh.Strength + (0 + rand.Intn(15-0+1)) - sh.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)

}
func (sh *Shaman) Health() int {
	return sh.Healths
}
func (sh *Shaman) StartHealth(points int) {
	sh.Healths = sh.Healths - points
}
func (sh *Shaman) StartDefense(point int) {
	sh.Defense = sh.Defense - point
}

func (kg *King) Attack(character Character) {
	kg.Fear += 12
	damage := kg.Strength + (0 + rand.Intn(12-0+1)) - kg.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)
}
func (kg *King) Defend() {
	kg.Fear += 5
	kg.Defense += 10 + (0 + rand.Intn(10-0+1))
}
func (kg *King) ability1(character Character) {
	if kg.Fear < 10 {
		fmt.Println("If you want to rule the world, you need to make a clear plan!")
		return
	}
	fmt.Println("It's time to show these weaklings who is the king here")
	kg.Strength += 5
	kg.Fear -= 15
	damage := kg.Strength + (5 + rand.Intn(10-5+1)) - kg.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)
}
func (kg *King) ability2(character Character) {
	if kg.Fear < 20 {
		fmt.Println("Fear has consumed you. You're still too young to die")
		return
	}
	fmt.Println("The one who shoots first laughs well!")
	kg.Fear -= 20
	kg.Strength += 5
	damage := kg.Strength + (0 + rand.Intn(20-0+1)) - kg.Defense
	fmt.Println("Damage received: ", damage)
	character.StartHealth(damage)
	character.StartDefense(damage)
}
func (kg *King) Health() int {
	return kg.Healths
}
func (kg *King) StartHealth(points int) {
	kg.Healths = kg.Healths - points
}
func (kg *King) StartDefense(point int) {
	kg.Defense = kg.Defense - point
}

func main()  {
	//Приветствие
	var choiceNull, choiceFirstCh string
	var i int
	i = 0
	Characters:= [count]Character{}
	for (i < 1) {
		
			fmt.Println("Welcome to the game (Text Saga). You have to choose two characters to start playing. " +
				"Press F to continue the game or Q to exit")
			fmt.Scan(&choiceNull)
			switch choiceNull {
			case "F":
				fmt.Println("You are welcome")
				i += 5
				break
			case "Q":
				return
			default:
				fmt.Println("неизвестная комманда")
				continue
			}
	}
	//Выбор первого персонажа
	i = 0
	for (i == 0) {
		for u:= 0; u < count; u++ {
			fmt.Println("Choose Your hero: W - warrior, S - shaman, K - king")
			fmt.Scan(&choiceFirstCh)
			switch choiceFirstCh {

			case "W":
				Characters[u] = &Warrior{Healths: 100, Strength: 10, Defense: 10, Confidence: 0}
				i += 5
				break
			case "S":
				Characters[u] = &Shaman{Healths: 100, Strength: 5, Defense: 5, Ether: 0}
				i += 5
				break
			case "K":
				Characters[u] = &King{Healths: 100, Strength: 15, Defense: 10, Fear: 0}
				i += 5
				break
			default:
				fmt.Println("неизвестная комманда")
				continue
			}
		}
	}

	var choice,p int //Реализация Боя
	p = 0
	for Characters[0].Health() >= 0 || Characters[2].Health() >= 0 {
		for (p == 0){

		fmt.Println("Choose Your character action: 1 - attack, 2 - defense, 3 - first ability (you need 10), 4 - second ability(you need 20), 9 - exit in Windows")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Characters[0].Attack(Characters[1])
			p += 5
			break
		case 2:
			Characters[0].Defend()
			p += 5
			break
		case 3:
			Characters[0].ability1(Characters[1])
			p += 5
			break
		case 4:
			Characters[0].ability2(Characters[1])
			p += 5
			break
		case 9:
			return
		default:
			fmt.Println("неизвестная комманда")
			continue
		}

		if Characters[1].Health() <= 0 {
			fmt.Println("The first one won")
			return
		}
		}
		p = 0
		for (p == 0) {
			fmt.Println("Choose Opponent character action: 1 - attack, 2 - defense, 3 - first skill(you need 10), 4 - second skill(you need 20), 9 - exit")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				Characters[1].Attack(Characters[0])
				p += 5
				break
			case 2:
				Characters[1].Defend()
				p += 5
				break
			case 3:
				Characters[1].ability1(Characters[0])
				p += 5
				break
			case 4:
				Characters[1].ability2(Characters[0])
				p += 5
				break
			case 9:
				return
			default:
				fmt.Println("неизвестная комманда")
				continue
			}
			if Characters[0].Health() <= 0 {
				fmt.Println("The second one won")
				return
			}
		}
	}
}