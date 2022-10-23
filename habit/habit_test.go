package habit

import "testing"

func TestNewHabitEmpty(t *testing.T) {
    h := NewHabit("")

    desc := h.Description
    impact := h.Impact

    if desc != "" {
        t.Errorf("description equal %s; want %s", desc, "")
    }

    if impact != Neutral {
        t.Errorf("impact equal %s; want %s", string(impact), string(Neutral))
    }
}

func TestNewHabitPositive(t *testing.T) {
    h := NewHabit("Reading Books+")

    desc := h.Description
    impact := h.Impact

    if desc != "Reading Books" {
        t.Errorf("description equal %s; want %s", desc, "Reading Books")
    }

    if impact != Positive {
        t.Errorf("impact equal %s; want %s", string(impact), string(Positive))
    }
}

func TestNewHabitNegative(t *testing.T) {
    h := NewHabit("Smoking-")

    desc := h.Description
    impact := h.Impact

    if desc != "Smoking" {
        t.Errorf("description equal %s; want %s", desc, "Smoking")
    }

    if impact != Negative {
        t.Errorf("impact equal %s; want %s", string(impact), string(Negative))
    }
}

func TestNewHabitUndefined(t *testing.T) {
    h := NewHabit("Waking Up")

    desc := h.Description
    impact := h.Impact

    if desc != "Waking Up" {
        t.Errorf("description equal %s; want %s", desc, "Waking Up")
    }

    if impact != Neutral {
        t.Errorf("impact equal %s; want %s", string(impact), string(Neutral))
    }
}

func TestNewHabitNeutral(t *testing.T) {
    h := NewHabit("Waking Up|")

    desc := h.Description
    impact := h.Impact

    if desc != "Waking Up" {
        t.Errorf("description equal %s; want %s", desc, "Waking Up")
    }

    if impact != Neutral {
        t.Errorf("impact equal %s; want %s", string(impact), string(Neutral))
    }
}
