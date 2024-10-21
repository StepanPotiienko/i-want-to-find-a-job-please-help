class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        self.erase_val_from_array(val, nums)
        while val in nums: 
            self.erase_val_from_array(val, nums)
        return len(nums)

    def erase_val_from_array(self, val, array):
        for i in array:
            print(i)
            if i == val:
                array.remove(i)
    