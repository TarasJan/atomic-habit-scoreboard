package main


import (
    "fmt"
    "bufio"
    "strings"
    "os"
    . "github.com/TarasJan/atomic-habit-scoreboard/habit"
    . "github.com/TarasJan/atomic-habit-scoreboard/scoreboard"
)

func main() {
    command := ""
    scoreboard := new(Scoreboard)
    reader := bufio.NewReader(os.Stdin)

    for command != "EXIT" {
        fmt.Printf(">> ")
        command, _ = reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        switch {
            case strings.HasPrefix(command, "EXIT"):
                break
            case strings.HasPrefix(command, "ADD "):
                scoreboard = scoreboard.Add(NewHabit(command[4:]))
                fmt.Println(*scoreboard)
            case strings.HasPrefix(command, "REMOVE "):
                scoreboard = scoreboard.Remove(command[7:])
                fmt.Println(*scoreboard)
            default:
               continue
        }
    }
}

