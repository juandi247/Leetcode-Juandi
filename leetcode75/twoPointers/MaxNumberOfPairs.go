package main
func maxOperations(nums []int, k int) int {

    numsMap:= make(map[int]int)
    operations:=0

    for _,v:=range nums{
        target:= k- v
        if numsMap[target]>0{
            numsMap[target]--
            operations++
        }else{
            numsMap[v]++
        }
       
    }


    return operations
}


/*

 for _,n:=range nums{
        lookfor:= k - n
        if check[lookfor] > 0{
            check[lookfor]--
            res++
        } else {
            check[n]++
        }
    }
    return res


Recorrer el array
Por cada valor hago un ++ y uso la misma lgoica

 */

/* 
NOT ORDERED ARRAY
Operation -> pick two numbers, if they sum up. remove them


1,2,3,4

I visualize it more as a two pointer with a -target sum.


We go and we can obtain with a map. that means o(n). and save how many times the numebrs appear

Then we can go troguht the array. Watch if the currentPointer - target = sum. if yes, we do a value -1 in both and thats it
we verify always that the value that we found has at least 1 or more, if not its not a valid elmeent. and thats it.

would be done in O(N)


3,1,3,4,3

We found 3. then we need to go and find another 3. we found it, on our map it would be 3 value is 1.



Can we find a two pitners aproach? that would be to o(n)

if we go at the start. and do it again. 


*/
