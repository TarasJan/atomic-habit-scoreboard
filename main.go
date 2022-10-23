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
            case command == "EXIT":
                break
            case strings.HasPrefix(command, "ADD "):
                scoreboard = scoreboard.Add(NewHabit(command[4:]))
                fmt.Println(*scoreboard)
            case strings.HasPrefix(command, "REMOVE "):
                scoreboard = scoreboard.Remove(command[7:])
                fmt.Println(*scoreboard)
            case strings.HasPrefix(command, "EXPORT "):
                err := scoreboard.Export(command[7:])
                if err != nil {
                    fmt.Println("Failed to export to filename: " + command[7:])
                } else {
                    fmt.Println("Successfully written file " + command[7:])
                }
            case strings.HasPrefix(command, "IMPORT "):
                ns, err := scoreboard.Import(command[7:])
                if err != nil {
                    fmt.Println("Failed to import to filename: " + command[7:])
                } else {
                    fmt.Println("Successfully read file " + command[7:])
                }
                scoreboard = ns
                fmt.Println(*scoreboard)

            default:
               continue
        }
    }
}

