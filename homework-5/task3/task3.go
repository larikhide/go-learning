package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

//3. Самостоятельно изучите пакет encoding/csv. Напишите
//пример, иллюстрирующий его применение.

func main() {
	//cоздать файл для записи чтения,если его нет, то создать
	file, err := os.OpenFile("task3example.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		return
	}
	defer file.Close()

	//инфа, которую нужно записать в csv
	records := [][]string{
		{"time_ms", "Current_A", "voltage_V"},
		{"1", "5", "10"},
		{"2", "7", "8"},
		{"3", "10", "5"},
	}

	//запсать в файл инфу
	w := csv.NewWriter(file)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv: ", err)
	}

	//прочитать из файл инфу
	r := csv.NewReader(file)
	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("теперь ты можешь построить график по координатам : %v\n", record)
}
