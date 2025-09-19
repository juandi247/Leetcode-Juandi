package main

import "strconv"

func MeetingBoth(skipper, cody, whSkipper, whCody [][2]string, meetingMinutes int) [][2]string {





}


func obtainFreeSpaces(plan, workHours [][2]string, meetingMinutes int ) [][2]string{

	validSpaces:= [][2]string{}

	// 1. compare the starting work hours wuth the first element
	workHours[0], plan[0][0]



	// 2. traverse the whole plan


	// 3. compare the last point with th elast working hour


}




func calculateTime(start, end string) int {

	// 12:15 - 10:30
	horas, err := strconv.Atoi(string(end[0:2]))
	horaAbajo, err := strconv.Atoi(string(start[0:2]))

	if err!=nil{
		return 0
	}

	// 2 horas
	horas= horas- horaAbajo


	minutosArriba, err:= strconv.Atoi(string(start[3:]))
	minutosAbajo, err:= strconv.Atoi(string(end[3:]))


	if minutosArriba> minutosAbajo{
		minutosArriba+=60
		horas-=1
	}

	totalMinutos:= minutosArriba-minutosAbajo
	totalTiempo:= (horas*60)+totalMinutos


	return totalTiempo


}
