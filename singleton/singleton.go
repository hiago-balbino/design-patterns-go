package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Used to read the `capitals` file to simulate data from database.
func readFileData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	data := map[string]int{}
	for scanner.Scan() {
		key := scanner.Text()
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		data[key] = value
	}

	return data, nil
}

var once sync.Once
var instance Database

// Database is an interface to handle the storage data.
type Database interface {
	// GetPopulationFrom returns the population of a given city.
	GetPopulationFrom(city string) int
}

// GetSingletonDatabase is a function to lazily initialize a database dependency.
// Initialization is done once since using the `sync.Once.Do` function.
func GetSingletonDatabase() Database {
	once.Do(func() {
		db := singletonDatabase{}
		capitals, err := readFileData("capitals.txt")
		if err == nil {
			db.capitals = capitals
		}
		instance = db
	})

	return instance
}

type singletonDatabase struct {
	capitals map[string]int
}

// GetPopulationFrom returns the population of a given city.
func (db singletonDatabase) GetPopulationFrom(city string) int {
	return db.capitals[city]
}

func main() {
	db := GetSingletonDatabase()
	population := db.GetPopulationFrom("Seoul")
	fmt.Println("Population = ", population)
}
