impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut candidate: i32 = 0;
        let mut count: i32 = 0;

        for num in nums.into_iter() {
            if count == 0 {
                candidate = num;
            }
            count += if num == candidate { 1 } else { -1 };
        }

        candidate
    }
}