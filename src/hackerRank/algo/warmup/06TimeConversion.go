/* Given a time in AM/PM format, convert it to military (24-hour) time.

Note: Midnight is 12:00:00AM on a 12-hour clock, and 00:00:00 on a 24-hour clock. Noon is
12:00:00PM on a 12-hour clock, and 12:00:00 on a 24-hour clock.

Input Format
A single string containing a time in 12-hour clock format (i.e.: hh:mm:ssAM or hh:mm:ssPM),
where 01≤hh≤12.

Output Format

Convert and print the given time in 24-hour format, where 00≤hh≤23.

Sample Input
07:05:45PM
Sample Output
19:05:45*/
package main

import (
	"fmt"
	"strconv"
	"os"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	time12 := string(b)
	fmt.Printf("%v", convertTime(time12))
}

func convertTime(time12 string) string {
	// corner cases
	if time12 == "12:00:00AM" {
		return "00:00:00"
	}

	if time12 == "12:00:00PM" {
		return "00:00:00"
	}

	hhStr := time12[0:2]
	permStr := time12[2:8]
	apm := time12[8:]

	if apm == "AM" {
		if hhStr == "12" {
			return "00" + permStr
		}

		return hhStr + permStr
	}

	if hhStr == "12" {
		return hhStr + permStr
	}

	hh, _ := strconv.Atoi(hhStr)

	hh += 12

	return strconv.Itoa(hh) + permStr
}