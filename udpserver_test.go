package main

import (
    "testing"
)

func TestTimeParseToUnix(t *testing.T) {

    const time_str = "[17/06/2015 12:10]"
    const time_unix = "1434543000"

    res := TimeParseToUnix(time_str)

    if res != time_unix {
        t.Error("Expected ", time_unix, "got ", )
    }
}



