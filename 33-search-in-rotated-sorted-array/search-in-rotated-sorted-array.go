func search(nums []int, target int) int {
    s := 0
    e := len(nums) -1
    return binary_search(nums , target, s, e)
}

func binary_search(nums []int, target int, s, e int) int {
    if s > e {
        return -1
    }
    mid := s + (e-s) / 2
    if  nums[mid] ==  target {
        return mid
    }
    if  nums[s] <= nums[mid] {
        if target >= nums[s] && target <= nums[mid] {
            return binary_search( nums, target, s , mid-1)
        } else {
            return binary_search( nums, target, mid+1 , e)
        }
    }
    if target > nums[mid] && target <= nums[e] {
        return binary_search( nums, target, mid+1 , e)
    } 
    return binary_search( nums, target, s , mid-1)
}