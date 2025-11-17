
/* 
just use a "queue", as soon as we find a timestamp that is valid, we dequeue all the stamps on the left becasue we know that they are invalid and return the length of the new array

*/
type RecentCounter struct {
    counter []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        counter:[]int{},
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.counter=append(this.counter, t)
    upperBound:= t-3000
    for i,v:=range this.counter{
        if v>=upperBound{
            this.counter= this.counter[i:]
            break
        }
    }

    return len(this.counter)
}
