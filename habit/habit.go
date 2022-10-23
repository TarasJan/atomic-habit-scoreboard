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
    Description string `json:"description"`
    Impact Impact `json:"impact"`
}

func (h *Habit) String() string {
    return h.Description + string(h.Impact)
}

func NewHabit(habitDesc string) *Habit {
    var impact Impact = Neutral
    habitDescLength := len(habitDesc)
    if habitDescLength == 0 {
        return &Habit{Description: habitDesc, Impact: impact}
    }

    desc := habitDesc
    finalDescRune := rune(habitDesc[habitDescLength - 1])
    switch(finalDescRune) {
        case '+', '-', '|':
            impact = Impact(finalDescRune)
            desc = desc[:habitDescLength - 1]
    }
    return &Habit{Description: desc, Impact: impact}
}
