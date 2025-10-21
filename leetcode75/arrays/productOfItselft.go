func productExceptSelf(nums []int) []int {

    leftPass:= []int{}

    acum:= 1
    for i,_:=range nums{
        if i==0{
            leftPass= append(leftPass, 1)
            continue
        }
        acum= acum*(nums[i-1])
        leftPass= append(leftPass, acum)
    }


length:=len(nums)
acum=1
for i:=len(nums)-1;i>=0; i--{
        if i==len(nums)-1{
           continue 
        }
        acum= acum*(nums[i+1])
        fmt.Println("posicion de i", length-i-1)
        leftPass[i]*= acum
        
}
return leftPass
}



/*
Recorremos el array dos veces (no usamos division)


La primera pasada entonces miramos los valores a la IZQUIERDA, es decir vamos poniendo el acumulado a la izquierda.
en el caso de 1,2,3,4

entonces seria 1,1, 2, 6.
porque entonces el primer valor queda igualito, el segundo es valores a la izquierda de 2? entonces es solo 1
y asi su cesivamente


Luego de vuelta seria lo mismo
24, 12, 4, 1.
 */
