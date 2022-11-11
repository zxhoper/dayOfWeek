package dayOfWeek

// day_of_week modify from RFC3339 Date and Time on the Internet: Timestamps
// https://www.rfc-editor.org/rfc/rfc3339
// The following is a sample Go subroutine loosely based on Zeller's
//
//	Congruence [Zeller] which may be used to obtain the day of the week
//	for dates on or after 0000-03-01:(0000-03-1 is Wednesday)
func DayOfWeek(day int, month int, year int) string {
	var cent int
	dayofweek := []string{
		"Sun",
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
	}
	// dayofweek := []string{
	//      "Sunday",
	//      "Monday",
	//      "Tuesday",
	//      "Wednesday",
	//      "Thursday",
	//      "Friday",
	//      "Saturday",
	// }

	// adjust months so February is the last one
	month -= 2
	if month < 1 {
		month += 12
		year -= 1
	}
	// split by century
	cent = year / 100
	year %= 100

	//(  (26*month-2)/10 +  day + year<365%7=1> +
	//    year/4<leap year> + cent/4   +  5*cent )  %7

	return (dayofweek[((26*month-2)/10+day+year+year/4+cent/4+5*cent)%7])
}
