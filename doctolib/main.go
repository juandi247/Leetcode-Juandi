/* 
Ejercicio 1 — Appointment Scheduler

Descripción:
Dado un rango de trabajo (por ejemplo, 09:00–17:00) y una lista de franjas horarias ya reservadas, devuelve todas las franjas disponibles de 30 minutos dentro del horario laboral.

Entrada:
workStart y workEnd: strings con formato "HH:MM".
booked: lista de strings con formato "HH:MM" que representan horas ya ocupadas.

Salida:
Lista de strings con todas las horas disponibles de 30 minutos.
*/


/* 
Input:
workStart = "09:00"
workEnd = "12:00"
booked = ["09:30", "10:30"]

Output:
["09:00", "10:00", "11:00", "11:30"]
*/ 


/* 
APROACH:

We receive an hour. The range betwen that hour - hour+30. its ocupied. 
So we would need to first. Traverse or create some kind of thing
to have all the hours available in order.
9, 9:30, 10, 10:30. etc. etc

Then we just traverse it again, and have a pointer to our initial array. so we only have a o(n)

we traverse, if an hour is a match. or is NOT a match, then we simply add it to our solution, and thats it.



EXAMPLE.
range: 9 - 11

create array of hours: 9:00 , 9:30, 10, 10:30, 11:00, 11:30, 12


Then we traverse this array. and start our pointer. 

We have our pointer on the booked part as 0. so we point to 9:30.

We start traversing so we start with 9:00. we check if this matches our pointer. if not. we add it to the answer array
We dont update our pointer there.

Then 9:30, matches, we update our pointer, and we do not add it to the response.

10:00, does not match with booked[1], because its 10:30, so we add it to the array. 
And continue. If we see that our pointer is bigger than the length - 1, means that we already llooked at tjhe booked
so we just copy the rest of the array. and thats it.







STEPS: 
A function that returns all the ranges.

We could even calculate that. but im not sure.
First aproach would be. Calculate the current time. if its a 30 time. then we return the next
9 - 12. 

If last time is minutes its :00, then we would add Hour:30
if the minutes are 30, then we add hour +1 and minutes:00
*/


package main
import ("strconv"

"fmt"
)
func main() {
fmt.Println("hola")
	fmt.Println(calculateNextTime("08:30"))
}

/* 
9 - 12

ocupied 
		P
[9:30, 10:30]
*/
func checkAvailableSpots(start string, end string, booked []string) []string{
	availableSpots:= []string{}

	schedule:= buildSchedule(start, end)

	startPointer:=0

	for i,v:=range schedule{
		if startPointer>=len(booked){
			newSlice:= schedule[i:]
			availableSpots = append(availableSpots, newSlice...)
			return availableSpots
		}

		
		if booked[startPointer]==v{
			startPointer++
			continue
		}

		availableSpots=append(availableSpots, v)
	}

	return availableSpots

}



func buildSchedule(start string, end string) []string{
	schedule:= []string{start}
	currTime:=start
					
	for currTime!=end{
		currTime = calculateNextTime(currTime)
		schedule=append(schedule, currTime)
	}

	schedule = append(schedule, end)
	return schedule
}


func calculateNextTime( time string) string{
	hour:= time[0:2]
	min:= time[3:]

	intHour,_:= strconv.Atoi(hour)

	if min=="30"{
		fmt.Println(min)
		fmt.Println(intHour+1)
		return strconv.Itoa(intHour+1)	+ ":00"	
	}
	return hour+":30"
}
