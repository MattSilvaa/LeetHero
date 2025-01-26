package bot

var Solutions = map[string]string{
	"two-sum": `class Solution:
   def twoSum(self, nums, target):
       seen = {}
       for i, n in enumerate(nums):
           if target - n in seen:
               return [seen[target - n], i]
           seen[n] = i
       return []`,
	"add-two-numbers": `class Solution:
    def addTwoNumbers(self, l1, l2):
        dummy = ListNode(0)
        curr = dummy
        carry = 0
        while l1 or l2 or carry:
            x = l1.val if l1 else 0
            y = l2.val if l2 else 0
            s = x + y + carry
            carry = s // 10
            curr.next = ListNode(s % 10)
            curr = curr.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return dummy.next`,
}
