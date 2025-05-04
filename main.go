package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Старт разработки приложения Погоды на Go")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)
}
