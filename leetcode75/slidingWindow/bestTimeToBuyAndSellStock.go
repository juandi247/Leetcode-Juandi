/* 
We have an array of stocks, we need to get the maximum profit, between a buy and a sale.
a buy can only ocurr before the sell
if there is no profit retunr 0


7,1,2,3,4,2,6

the answer will be 5, becaue the profit with the most value was 6 (sell)- 1 (buy).

Every index of the array represents a DAY so we can not buy on the index 6, and sell on the index 3.

APRAOCH:
mantain the biggest value and the smallest value between a window.
When we expand the window and find a value that is smaller than our current smallest value. MEANS that we need to upate the start of the window to that smaller value. and search apartir de ahi.

We keep track of the biggest profit everytime we can, so that at the end if the last window doesnt contain that value is already saved

*/j

func bestTimeToSellAndBuy(prices []int) int{

	if len(prices)<=1{
		return 0
	}
	start:=0
	end:=1

	small:=prices[start]
	big:=prices[start]

	profit:=0


	for end<len(prices){
		

		if prices[end]>big{
		profit= max(profit, prices[end]-prices[start]
		big= prices[end]
		}

		//means that we need to start a  new window
		if prices[end]<big{
			start=end
			small=prices[start]
			big=prices[start]
		}

		end++
	}


	return profit
}
