package main

import (
  	//"fmt"
	"time"
	//"io"
	"bufio"
	"os"
	"strconv"
)

var years int = 101

type Retail454 struct {
	Retail454year int //* update on WoY = 1 && DoW = 1, +1
	Retail454month int //* update on DoM = 1, %12 + 1
	Retail454weekofyear int //* update on DoW = 1, +1 unless MoY = 1 AND WoY > 50, then =1
	Retail454dayofweek int //* update every day, %7 + 1
	Retail454dayofmonth int //* update every day, +1 unless DoM = 28 or 35 depending on WiM, then =1
	Retail454weekofmonth int //* update on DoW = 1, +1 unless WoM = WiM, then =1
	Retail454weeksinmonth int //* update on DoM = 1, use calcWeeksInMonth
}

func calcRetail454(t time.Time, i int, r Retail454) Retail454 {
	if r.Retail454year == 0 {
		r = Retail454{2000, 1, 1, 1, 1, 1, 4}
		return r
	}

	// calculate day of week
	r.Retail454dayofweek = (i % 7) + 1

	// calculate day of month
	if ((r.Retail454weeksinmonth == 4 && r.Retail454dayofmonth == 28) || (r.Retail454weeksinmonth == 5 && r.Retail454dayofmonth == 35)) {
		if (r.Retail454month == 12){
			r.Retail454month = 1
		} else {
			r.Retail454month += 1
		}
		r.Retail454dayofmonth = 1
	} else {
		r.Retail454dayofmonth += 1
	}

	// set the week of month number
	if r.Retail454dayofweek == 1 {
		if r.Retail454weekofmonth == r.Retail454weeksinmonth {
			r.Retail454weekofmonth = 1
		}else{
			r.Retail454weekofmonth += 1
		}
	}

	// calculate number of weeks in month if it is a new month
	if r.Retail454dayofmonth == 1 {
		r.Retail454weeksinmonth = calcWeeksInMonth(t, r)
	}

	// update week of year
	if r.Retail454dayofweek == 1 {
		if r.Retail454month == 1 {
			if r.Retail454weekofyear >= 50 {
				r.Retail454weekofyear = 1
			} else {
				r.Retail454weekofyear += 1
			}
		} else {
			r.Retail454weekofyear += 1
		}
	}

	// update year if day 1 and week of year is 1
	if r.Retail454dayofweek == 1 && r.Retail454weekofyear == 1 {
		r.Retail454year += 1
	}
	
	return r
}

// calculate number of weeks in 454 month
func calcWeeksInMonth(t time.Time, r Retail454) int {
	if r.Retail454month == 2 || r.Retail454month == 5 || r.Retail454month == 8 || r.Retail454month == 11 {
		return 5
	}else if r.Retail454month == 12{
		if t.Month() == 12{
			return 5
		}
		return 4
	}else{
		return 4
	}
}

func main() {
	r := Retail454{}
	// Open output file for writing
// open output file
    fo, err := os.Create("output.txt")
    if err != nil { panic(err) }

    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    str := "blank"

	// generate dates, starting at Jan 30 (Gregorian), for a number of years
	for i:=0; i <= (365 * years); i++ {
		t := time.Date(2000, time.January, (30 + i), 12, 0, 0, 0, time.UTC)
		r = calcRetail454(t, i, r)
		// debug output 
		// fmt.Printf("Gregorian %d %2d %2d %s %d %2d %2d %2d %2d %2d %2d %s\n", t.Year(), int(t.Month()), t.Day(), "  ||  ",
		//	 			r.Retail454year, r.Retail454weekofyear, r.Retail454dayofweek, r.Retail454month, r.Retail454weekofmonth, r.Retail454weeksinmonth, r.Retail454dayofmonth,  "Retail Date")
		

		// build line to write to file
		str = strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()) + " " + strconv.Itoa(r.Retail454year) + " " + strconv.Itoa(r.Retail454weekofyear) + " " + strconv.Itoa(r.Retail454dayofweek) + "\r"
		
		//Print to file
        	buf = []byte(str)
        	// write a chunk
        	if _, err := w.Write(buf); err != nil {
		        panic(err)
       		}
       		if err = w.Flush(); err != nil { panic(err) }

	}
}
