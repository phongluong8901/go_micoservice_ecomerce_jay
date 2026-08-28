package helper

import (
	"crypto/rand"
)

func RandomNumbers(length int) (string, error) {
	const numbers = "123456789"

	//Tạo một mảng byte ([]byte) tạm thời có kích thước đúng bằng length để chứa các giá trị ngẫu nhiên thô.
	buffer := make([]byte, length)
	_, err := rand.Read(buffer) //Đọc dữ liệu ngẫu nhiên: Dùng hàm an toàn rand.Read để lấp đầy mảng buffer bằng các byte ngẫu nhiên từ hệ thống
	if err != nil {
		return "", err
	}

	numLength := len(numbers)

	for i := 0; i < length; i++ {
		//int(buffer[i]) % numLength: Lấy giá trị byte ngẫu nhiên chia lấy phần dư cho tổng số ký tự nguồn (9). Thao tác này giúp ánh xạ một byte ngẫu nhiên bất kỳ thành một chỉ mục hợp lệ từ 0 đến 8.
		buffer[i] = numbers[int(buffer[i])%numLength]
	}

	return string(buffer), nil //Chuyển chuỗi số đó thành kiểu số nguyên (int).
}
