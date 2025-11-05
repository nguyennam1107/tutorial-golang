# Bài Tập Thực Hành - Basics

## Bài 1: Variables và Types
1. Khai báo và khởi tạo các biến với các kiểu khác nhau
2. Thực hiện type conversion giữa int, float64, string
3. Tạo constants cho PI, tên ứng dụng, version

## Bài 2: Functions
1. Viết hàm tính giai thừa (factorial)
2. Viết hàm kiểm tra số nguyên tố
3. Viết hàm tính số Fibonacci thứ n
4. Viết hàm với variadic parameters để tìm min/max

## Bài 3: Control Flow
1. Viết chương trình in ra bảng cửu chương từ 1-10
2. Implement FizzBuzz từ 1-100
3. Tìm các số hoàn hảo (perfect numbers) < 10000

## Bài 4: Arrays và Slices
1. Tạo slice và thực hiện: append, remove, insert, reverse
2. Merge 2 sorted slices thành 1 sorted slice
3. Tìm phần tử xuất hiện nhiều nhất trong slice

## Bài 5: Maps
1. Tạo map để đếm tần suất ký tự trong chuỗi
2. Group students theo class
3. Implement cache cho Fibonacci sử dụng map

## Bài 6: Structs
1. Tạo struct Person với Name, Age, Email
2. Tạo slice of Person và sort theo Age
3. Implement method String() cho Person

## Bài 7: Methods và Interfaces
1. Tạo interface Shape với methods Area() và Perimeter()
2. Implement Shape cho Circle, Rectangle, Triangle
3. Viết hàm tính tổng diện tích của slice of Shapes

## Bài 8: Error Handling
1. Viết hàm validate email address (trả về error nếu invalid)
2. Tạo custom error type cho validation errors
3. Wrap errors với context sử dụng fmt.Errorf()

## 🎯 Challenges

### Challenge 1: Calculator
Tạo calculator với các operations: +, -, *, /, %, ^
- Support nhiều operations liên tiếp
- Error handling cho division by zero
- Support parentheses

### Challenge 2: Todo List
Implement todo list manager:
- Add/Remove/Update/List todos
- Mark as complete
- Filter by status
- Save/Load từ file

### Challenge 3: Student Management
- Thêm/Xóa/Sửa student
- Tính điểm trung bình
- Xếp hạng students
- Export ra file

Xem file solutions.go để tham khảo đáp án!
