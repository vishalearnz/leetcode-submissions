func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0;
    }
    //Create a set and add all the array contents into the set so that even if duplicates come it doesnt alter the count
    numSet := make(map[int]bool, len(nums))
    for _, num := range nums {
        numSet[num] = true
    }
    longestStreak := 0
    //Interate and check if the 
    for num := range numSet {
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1

            for(numSet[currentNum + 1]) {
                currentNum++
                currentStreak++
            }
            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }

        }
    }
    return longestStreak
}