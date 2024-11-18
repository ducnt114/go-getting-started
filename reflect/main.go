package main

import (
	"fmt"
	"reflect"
)

// PrintStructAttributesAndValues takes a slice of structs and prints the attribute names and their values.
func PrintStructAttributesAndValues(slice interface{}) {
	// Use reflection to ensure the input is a slice
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		fmt.Println("Expected a slice of structs")
		return
	}

	// Ensure the slice contains structs
	if val.Len() == 0 {
		fmt.Println("Empty slice provided")
		return
	}

	firstElem := val.Index(0)
	if firstElem.Kind() != reflect.Struct {
		fmt.Println("Expected a slice of structs")
		return
	}

	// Print attribute names from the first struct
	fmt.Println("Attributes:")
	for i := 0; i < firstElem.NumField(); i++ {
		fmt.Printf("%s\t", firstElem.Type().Field(i).Name)
	}
	fmt.Println()

	// Print values for each struct in the slice
	fmt.Println("Values:")
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		for j := 0; j < elem.NumField(); j++ {
			fmt.Printf("%v\t", elem.Field(j).Interface())
		}
		fmt.Println()
	}
}

func main() {
	// Example usage
	type ExampleStruct struct {
		Name  string
		Age   int
		Email string
	}
	list := []ExampleStruct{
		{
			"Alice", 30, "test",
		},
		{
			"Bob", 60, "bob-test",
		},
	}
	ExportCsvToConsole(list)
}

func ExportCsvToConsole(arg interface{}) {
	val := reflect.ValueOf(arg)
	for i := 0; i < val.NumField(); i++ {
		fmt.Print(val.Type().Field(i).Name)
		fmt.Print(",")
	}
	fmt.Println()
	for i := 0; i < val.NumField(); i++ {
		fmt.Print(val.Field(i).Interface())
		fmt.Print(",")
	}
	fmt.Println()
}
