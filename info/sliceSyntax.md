package slice

/* 

| Syntax                | Meaning           |
| --------------------- | ----------------- |
| `var s []int`         | empty slice       |
| `s := []int{1,2,3}`   | initialized slice |
| `s := make([]int, 3)` | slice of 3 zeros  |
| `s = append(s, x)`    | add element       |

*/




****
🧩 Example Slice
nums := []int{10, 20, 30, 40, 50}


Suppose you want to remove the element at index 2 (30).

✅ Method 1 — Using append() (most common)
index := 2
nums = append(nums[:index], nums[index+1:]...)
fmt.Println(nums)


Output:

[10 20 40 50]


💡 Explanation:

nums[:index] → elements before index (10, 20)

nums[index+1:] → elements after index (40, 50)

The ... (variadic) expands the slice when appending

✅ Method 2 — Remove by value

If you know the value (not the index):

val := 30
for i, v := range nums {
    if v == val {
        nums = append(nums[:i], nums[i+1:]...)
        break
    }
}
fmt.Println(nums)


Output:

[10 20 40 50]

✅ Method 3 — Remove multiple values (filter style)

If you want to remove all 30s, for example:

result := []int{}
for _, v := range nums {
    if v != 30 {
        result = append(result, v)
    }
}
fmt.Println(result)

⚡ Summary Table
Goal	Code
Remove by index	s = append(s[:i], s[i+1:]...)
Remove by value	loop + append
Remove multiple	filter into new slice