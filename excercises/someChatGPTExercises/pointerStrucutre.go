package somechatgptexercises
//Punteros y Estructuras Anidadas
// Crea una estructura Person con nombre, edad y dirección.
//  Crea una función que actualice la dirección usando un puntero a Person.

type Person struct {
	Name string
	Age int
	Addr string
}

func NewPerson( Age int,Addr,Name string) *Person{
	newPerson:= Person{
		Name: Name,
		Age: Age,
		Addr: Addr,
	}
	return &newPerson

}

func (p *Person) WriteAddres(Addr string){
	p.Addr= Addr
}