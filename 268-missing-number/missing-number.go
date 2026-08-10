func missingNumber(nums []int) int {
    slices.Sort(nums)
    for i := 0; i<len(nums); i++ {
        if nums[i] != i {
            return i
        }
    }
    return len(nums)
    
}