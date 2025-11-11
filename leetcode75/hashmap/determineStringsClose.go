
package main

/*
AProach: 
No importan realmente las letras.

Lo uqe importa son las frecuencias.

Comparamos primero que si coincidan las letras que tenga. Si uno tiene una letra que no exiset en el otro CHAO. 

Comparar ahora las frecuencias.
Creamos un mapa con las frecuecnias y cuantas veces salen y luego lo comparamos con el otro mapa y listo.

*/


func closeStrings(word1 string, word2 string) bool {
//    base cases
    if len(word1)!= len(word2){
        return false
    }

    if len(word1)==0 && len(word2)==0{
        return true
    }

map1:= make(map[rune]int)
map2:= make(map[rune]int)
// base case they have different letters
for _,v:=range word1{
    map1[v]++
}

for _,v:=range word2{
    _, exist:= map1[v]
    if !exist{
        return false
    }
    map2[v]++
}




counterMap:=make(map[int]int)
for _,val:=range map1{
    counterMap[val]++
}


// luego entonces recorremos el mapa que teniamos de letras, pero lo recoreemos para el valor y preugntamos, este valor existe? si no existe CHAO

for _, val:=range map2{
    // if our count is 4 in map 2, and there is NO count 4 in the map 1 , means that they are not frecuency equal
    counterMapValue, exist:= counterMap[val]
    if !exist || counterMapValue<=0{
        return false
    }
    counterMap[val]--
}


return true
}
