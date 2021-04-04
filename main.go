// main.go
package main
import (
	"fmt"
)

func main() {
	fmt.Println("HAllo")
	a:=App{}
	a.Initialize(
		"postgres",//os.Getenv("APP_DB_USERNAME"),
		"1234567890",//os.Getenv("APP_DB_PASSWORD"),
		"cdue02")//os.Getenv("APP_DB_NAME"))
	a.Run(":8010")
}
