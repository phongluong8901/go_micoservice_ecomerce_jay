package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdateAt  time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}

//Các thẻ gorm:"..." trong đoạn code trên được sử dụng bởi thư viện GORM (một ORM phổ biến trong Go) để cấu hình cách các trường (field) trong struct ánh xạ xuống các cột (column) trong cơ sở dữ liệu quan hệ (như MySQL, PostgreSQL,...).
// Khóa chính (Primary Key): Xác định trường ID là khóa chính (Primary Key) của bảng.

// Ràng buộc cho Emai: gorm:"index;unique;not null"
// index: Tạo một chỉ mục trên cột email giúp tối ưu hóa tốc độ tìm kiếm theo email.
// unique: Đảm bảo các giá trị email trong cơ sở dữ liệu phải là duy nhất, không được phép trùng lặp.
// not null: Ràng buộc cột này không được phép để trống (không nhận giá trị NULL) khi thêm dữ liệu.

// Giá trị mặc định (Default Values)
// gorm:"default:false" (ở trường Verified): Thiết lập giá trị mặc định là false khi tạo mới một người dùng (chưa xác thực).
// gorm:"default:buyer" (ở trường UserType): Thiết lập giá trị mặc định cho loại tài khoản là "buyer" nếu không được chỉ định cụ thể.
// gorm:"default:current_timestamp" (ở trường CreatedAt và UpdateAt):Tự động gán giá trị mặc định là thời điểm hiện tại (CURRENT_TIMESTAMP) của cơ sở dữ liệu khi bản ghi được khởi tạo.
