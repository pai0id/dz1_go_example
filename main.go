package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var NMAX, NMIN int
	var finPath, foutPath string

	flag.IntVar(&NMAX, "NMAX", 0, "Максимальный размер массива (обязательный)")
	flag.IntVar(&NMIN, "NMIN", 0, "Начальный размер массива")
	flag.StringVar(&finPath, "fin", "", "Путь к файлу с сценарием работы с массивом")
	flag.StringVar(&foutPath, "fout", "", "Путь к файлу для записи действий пользователя")

	flag.Parse()

	if NMAX == 0 {
		fmt.Println("Необходимо задать обязательный флаг -NMAX")
		return
	}

	if NMAX < 0 {
		fmt.Println("NMAX не может быть меньше нуля")
		return
	}

	if finPath != "" && foutPath != "" {
		fmt.Println("Нельзя одновременно использовать -fin и -fout")
		return
	}

	var array []int
	if NMIN > 0 {
		rand.Seed(time.Now().UnixNano())
		array = rand.Perm(NMIN)
		printArray(array)
	} else if NMIN < 0 {
		fmt.Println("NMIN не может быть меньше 0")
		return
	}

	var fout *os.File
	if foutPath != "" {
		var err error
		fout, err = os.Create(foutPath)
		if err != nil {
			fmt.Println("Ошибка при открытии файла для записи:", err)
			return
		}
		defer fout.Close()
	}

	if finPath != "" {
		file, err := os.Open(finPath)
		if err != nil {
			fmt.Println("Ошибка при открытии файла со сценарием:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			command := scanner.Text()
			processCommand(command, &array)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка при чтении файла со сценарием:", err)
			return
		}
	}

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1 - Добавить элемент")
		fmt.Println("2 - Удалить элемент")
		fmt.Println("3 - Добавить 1 к каждому элементу")
		fmt.Println("0 - Выйти")

		var choice int
		fmt.Print("Выберите операцию (0-3): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var elem int
			fmt.Print("Введите элемент для добавления: ")
			fmt.Scan(&elem)
			array = addElement(array, elem)
			writeToFile(fout, fmt.Sprintf("addElement\n%d\n", elem))
			printArray(array)
			if len(array) < NMAX {
				continue
			}
			fallthrough
		case 0:
			fmt.Println("Программа завершена.")
			return
		case 2:
			array = removeElement(array)
			printArray(array)
			writeToFile(fout, "removeElement\n")
		case 3:
			array = addOneToArray(array)
			printArray(array)
			writeToFile(fout, "addOneToArray\n")
		default:
			fmt.Println("Неверный выбор. Попробуйте еще раз.")
		}
	}
}

func printArray(array []int) {
	fmt.Println("Текущий массив:", array)
}

func writeToFile(file *os.File, str string) {
	if file != nil {
		file.WriteString(str)
	}
}

func processCommand(command string, array *[]int) {
	var validAddElement = regexp.MustCompile(`addElement [0-9]+`)

	switch {
	case validAddElement.MatchString(command):
		words := strings.Fields(command)
		num, _ := strconv.Atoi(words[1])

		*array = addElement(*array, num)
		printArray(*array)
	case command == "removeElement":
		*array = removeElement(*array)
		printArray(*array)
	case command == "addOneToArray":
		*array = addOneToArray(*array)
		printArray(*array)
	default:
		fmt.Println("Неизвестная команда:", command)
	}
}
