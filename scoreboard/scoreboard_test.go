package scoreboard

import (
    "testing"
    "os"
    . "github.com/TarasJan/atomic-habit-scoreboard/habit"
)

func TestAddHabit(t *testing.T) {
    scoreboard := new(Scoreboard)
    h := NewHabit("Reading +")
    scoreboard = scoreboard.Add(h)
    if len(*scoreboard) != 1 {
        t.Errorf("Scoreboard does not have the expected size of %d instead it has the size of %d", 1, len(*scoreboard))
    }
    if (*scoreboard)[0] != h {
        t.Errorf("The object in the scoreboard is not the habit that was added")
    }
}

func TestRemoveByDescription(t *testing.T) {
    scoreboard := new(Scoreboard)
    h := NewHabit("Reading +")
    scoreboard = scoreboard.Add(h)
    scoreboard = scoreboard.Remove("Reading ")
    if len(*scoreboard) != 0 {
        t.Errorf("Scoreboard does not have the expected size of %d instead it has the size of %d", 0, len(*scoreboard))
    }
}

func TestRemoveByIndex(t *testing.T) {
    scoreboard := new(Scoreboard)
    h := NewHabit("Reading +")
    h2 := NewHabit("Exercising +")
    scoreboard = scoreboard.Add(h)
    scoreboard = scoreboard.Add(h2)
    scoreboard = scoreboard.Remove("1")
    if len(*scoreboard) != 1 {
        t.Errorf("Scoreboard does not have the expected size of %d instead it has the size of %d", 0, len(*scoreboard))
    }
    if (*scoreboard)[0] != h2 {
        t.Errorf("The object in the scoreboard is not the habit should have stayed in the scoreboard")
    }
}

func TestExport(t *testing.T) {
    scoreboard := new(Scoreboard)
    filename := "export.json"
    h := NewHabit("Reading +")
    h2 := NewHabit("Exercising +")
    scoreboard = scoreboard.Add(h)
    scoreboard = scoreboard.Add(h2)
    defer os.Remove(filename)
    err := scoreboard.Export(filename) 

    expected := "[{\"description\":\"Reading \",\"impact\":43},{\"description\":\"Exercising \",\"impact\":43}]"
    if err != nil {
        t.Errorf(err.Error())
        t.Errorf("An error occurred during file export")
    } else {
        data, err := os.ReadFile(filename)
        if err != nil {
            t.Errorf("An error occurred during file export")
        }
        if string(data) != expected {
            t.Errorf("The content of the JSON : %s\nExpected: %s", string(data), expected)
        }
    }
}

func TestImport(t *testing.T) {
    scoreboard := new(Scoreboard)
    filename := "export.json"
    data := "[{\"description\":\"Reading \",\"impact\":43},{\"description\":\"Exercising \",\"impact\":43}]"
    err := os.WriteFile(filename, []byte(data), 0644)
    defer os.Remove(filename)
    scoreboard, err = scoreboard.Import(filename) 
    if err != nil {
        t.Errorf(err.Error())
    } 
    if len(*scoreboard) != 2 {
        t.Errorf("Scoreboard does not have the expected size of %d instead it has the size of %d", 2, len(*scoreboard))
    }
    if (*scoreboard)[0].Description != "Reading " {
        t.Errorf("The object in the scoreboard is not the habit that was added")
    }
   if (*scoreboard)[1].Description != "Exercising "{
        t.Errorf("The object in the scoreboard is not the habit that was added")
    }
}

