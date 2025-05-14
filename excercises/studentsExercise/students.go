package studentsexercise

// Crear, actualizar y eliminar estudiantes.
// Buscar estudiantes por nombre.
// Calcular el promedio de sus calificaciones.
type Student struct{
	name string
	age int
	grade int
}


func NewStudent(name string, age, grade int) *Student{

	student:=Student{
		name: name,
		age: age,
		grade: grade,
	}
	
	return &student
}
