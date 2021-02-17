package main

import (
    "encoding/json"
    "errors"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

type Day struct {
    Day string
    Meetings []Meeting
}

type Meeting struct {
    Time int
    Duration int
    Name string
    Description string
    Program string
    Link string
}

func (m Meeting) Print() {
    fmt.Println("Name:", m.Name)
    fmt.Println("Description:", m.Description)
    fmt.Println("Program:", m.Program)
    fmt.Println("Link:")
    fmt.Println(m.Link)
}

func makemt() int {
    t := time.Now()
    return t.Hour() * 100 + t.Minute()
}

func strMt() string {
    t := time.Now()
    return fmt.Sprintf("%d:%d", t.Hour(), t.Minute())
}

func pschedule(schedule []Day) {
    fmt.Println("Your schedule:")
    for _, d := range schedule {
        fmt.Printf("<<<<<<<<<<< %s <<<<<<<<<<<\n", d.Day)
        for _, m := range d.Meetings {
            fmt.Printf(">>>>>>>>>>> %d >>>>>>>>>>>\n", m.Time);
            m.Print()
            fmt.Println();
        }
    }
}

func scheduleFromFile(filePath string) ([]Day, error) {
    // Read schedule file.
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Print(err)
    }

    // Parse into values.
    var schedule []Day
    err = json.Unmarshal(data, &schedule)
    if err != nil {
        return nil, errors.New("Unable to parse schedlule.")
    }

    return schedule, nil
}

func main() {

    // Get the schedule from file.
    filePath := "{path-to-project}/shared/schedule.json"
    schedule, err := scheduleFromFile(filePath)
    if err != nil {
        fmt.Println("Error reading file at: ", filePath)
    }

    // Read flag arguments.
    lastPtr := flag.Bool("last", false, "Get the room for the last meeting.")
    nextPtr := flag.Bool("next", false, "Get the room of the meeting happening next.")
    flag.Parse()

    if !*nextPtr && !*lastPtr {
        // fmt.Printf("Schedule: %+v\n", schedule)
        pschedule(schedule)
        fmt.Println("-------------")
        fmt.Println("Usage:")
        flag.PrintDefaults()
        os.Exit(1)
    }

    mt := makemt()
    wi := int(time.Now().Weekday())
    todayMeetings := schedule[wi - 1].Meetings

    // 11:00 >> 11:30, 
    // Move til now + 10 == last meeting
    // Return next meeting or size

    // 11:00 >> 11:30
    // when reach 11:30 or more then it is no longer next
    // Give 15 min of wiggle room.
    lasti := -1 // If there is no meetings leave it negative
    for i, v := range todayMeetings {
        if mt - 15 >=  v.Time {
            lasti = i
        }
    } // lasti is the last meeting

    fmt.Println("Current time:", strMt())

    if *lastPtr {
        fmt.Println("Last meeting:")
        if (lasti >= 0) {
            last := todayMeetings[lasti]
            fmt.Printf("Started at: %d\n", last.Time)
            last.Print()
        } else {
            fmt.Println("There were no meetings today.")
        }
    }

    if *nextPtr {
        // check if there is a next meeting that day, if there is set it
        nexti := -1
        if (lasti < len(todayMeetings) - 1) {
            nexti = lasti + 1
        }

        fmt.Println("Next meeting:")
        if (nexti >= 0) {
            next := todayMeetings[nexti]
            fmt.Printf("Starts at: %d\n", next.Time)
            next.Print()
        } else {
            fmt.Println("No more meetings next.")
        }
    }
}
