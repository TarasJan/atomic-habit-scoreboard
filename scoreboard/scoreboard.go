package scoreboard

import (
    "strings"
    "fmt"
    "encoding/json"
    "os"
    . "github.com/TarasJan/atomic-habit-scoreboard/habit"
)

type Scoreboard []*Habit


func (s *Scoreboard) Add(h *Habit) *Scoreboard {
    ns := append(*s, h)
    return &ns 
}

func (s *Scoreboard) Remove(desc string) *Scoreboard {
    ns := new(Scoreboard)
    for i, habit := range *s {
        if habit.Description != desc && fmt.Sprintf("%d", i + 1) != desc {
            *ns = append(*ns, habit)
        }
    }
    return ns
}

func (s Scoreboard) String() string {
    buffer := make([]string, 0)
    for i, habit := range s {
        buffer = append(buffer, fmt.Sprintf("%d) %s", i+1, habit.String())) 
    }

    return strings.Join(buffer, "\n")
}

func (s *Scoreboard) Export(filename string) error {
    data, err := json.Marshal(s)

    if err != nil {
        return err
    }
    
    err = os.WriteFile(filename, []byte(data), 0644)
    if err != nil {
        return err
    }
    return nil
}

func (s *Scoreboard) Import(filename string) (*Scoreboard, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return new(Scoreboard), err
    }
    var board *Scoreboard

    if err = json.Unmarshal(data, &board); err != nil {
        return board, err
    } else {
        return board, nil
    }
}
