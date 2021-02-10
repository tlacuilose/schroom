package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

func main() {
    // Read schedule file.
    filePath := "{pathtosrc}/shared/schedule.json"
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Print(err)
    }

    // Parse into values.
    var schedule []Day
    err = json.Unmarshal(data, &schedule)
    if err != nil {
        fmt.Println("Error reading file at: ", filePath)
    }

    // Read flag arguments.
    nextPtr := flag.Bool("next", false, "Get the room of the meeting happening next.")
    flag.Parse()

    if !*nextPtr {
        fmt.Printf("Schedule: %+v\n", schedule)
        fmt.Println("Usage:")
        flag.PrintDefaults()
        os.Exit(1)
    }

    t := time.Now()
    w := t.Weekday()
    mt := t.Hour() * 100 + t.Minute()
    wi := int(w)
    todayMeetings := schedule[wi - 1].Meetings

    icur := 0
    // if time 1000:
    //     900
    //     if 905 < 900
    for i, v := range todayMeetings {
        if mt + 5 > v.Time {
            icur = i
        }
    }

    next := todayMeetings[icur]

    fmt.Println("Current time:", mt)
    fmt.Println("Got meeting:", next.Name)
    fmt.Println("Description:", next.Description)
    fmt.Println("Program:", next.Program)
    fmt.Println("Link:")
    fmt.Println(next.Link)
}

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

func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}
