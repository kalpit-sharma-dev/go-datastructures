package arrays
 import "sort"


func threeSum(nums []int) [][]int {
	var values [][]int
		if len(nums)<3{
			return values
		}
	   
		sort.Ints(nums)
		targetSum:=0
		for i:=0;i<len(nums)-2;i++{
			j:=i+1
			k:=len(nums)-1
			
			if i > 0 && nums[i] == nums[i-1] {
				continue//To prevent the repeat
			}
	
			for j<k{
				cSum:=0
				cSum=cSum+nums[i] 
				cSum+=nums[j]
				cSum+=nums[k]
			
				if cSum==targetSum{
					out := []int{nums[i],nums[j],nums[k]}
					 
					for j < k && nums[j] == nums[j+1]{
						j++
					}
					for j < k && nums[k] == nums[k-1]{
						k--
					}
					values=append(values,out)
					j++
					k--
				}else if cSum>targetSum{
						k--                
				}else{
					j++
				}
			   
			}
		}
		return values
		
	}