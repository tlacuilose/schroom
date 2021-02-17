# Schroom: sch(edule of my meeting)room(s)
This is a CLI app to store and read your meeting room schedule
### Usage
- -last
    - Get the room for the last meeting.
- -next
    - Get the room of the meeting happening next.
### Setup

Add shared/schedule.json in proyect folder.
Example json:
```
[{
    "day": "Monday",
    "meetings": [{
        "time": HOUR AS INT ex. 1600,
        "duration": MINUTES AS INT ex. 90,
        "name": "NAME",
        "description": "DESCRIPTION ex, members",
        "program": "PROGRAM ex. Zoom",
        "link": "URL, such as zoom link"
    }, {
        ... more meetings ordered by time
    }]
}, {
    ... more days of the week ordered starts in monday ends in sunday
}]
```
