package main


import (
    "fmt"
    "github.com/TarasJan/atomic-habit-scoreboard/habit"
)

func main() {
    h := habit.NewHabit("Reading books +")
    fmt.Println("Hello Scoreboard")
    fmt.Println("My new habit is:")
    fmt.Println(h.Description())
    fmt.Println("Habit impact: " + string(h.Impact()))
}
