package habit

type Impact rune

const (
    Positive Impact = '+'
    Neutral = '|'
    Negative = '-'
)

func (i Impact) String() string {
    switch(i) {
        case Positive:
            return "Positive"
        case Neutral:
            return "Neutral"
        case Negative:
            return "Negative"
        default:
            return "Unknown"
    }
}

type Habit struct {
    description string
    impact Impact
}

func (h *Habit) Description() string {
    return h.description
}

func (h *Habit) Impact() Impact {
    return h.impact
}

func (h *Habit) String() string {
    return h.Description() + string(h.Impact())
}

func NewHabit(habitDesc string) *Habit {
    var impact Impact = Neutral
    habitDescLength := len(habitDesc)
    if habitDescLength == 0 {
        return &Habit{description: habitDesc, impact: impact}
    }

    desc := habitDesc
    finalDescRune := rune(habitDesc[habitDescLength - 1])
    switch(finalDescRune) {
        case '+', '-', '|':
            impact = Impact(finalDescRune)
            desc = desc[:habitDescLength - 1]
    }
    return &Habit{description: desc, impact: impact}
}
