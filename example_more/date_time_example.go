package main

import (
	"time"
	"fmt"
)

func main()  {
	now := time.Now()
	fmt.Println("Now:", now)
	fmt.Println("Day:", now.Day())
	fmt.Println("Month:", now.Month())
	fmt.Println("Year:", now.Year())
	fmt.Println("Year Day:", now.YearDay())
	fmt.Println("Time:", now.Hour(), now.Minute(), now.Second())
	fmt.Println("Nanosecond", now.Nanosecond())
	fmt.Println("Local", now.Local())
	fmt.Println("Location", now.Location())
	fmt.Println("UTC", now.UTC())
	fmt.Printf("%02d/%02d/%04d\n", now.Day(), now.Month(), now.Year())

	f := now.Format("20161212")
	fmt.Println(now, "=>", f)
}
