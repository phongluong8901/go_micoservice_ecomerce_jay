# --- lib
gofiber/fiber/v2: Web framework siêu tốc (cú pháp giống Express.js) dùng để xây dựng API và xử lý HTTP request/response.
joho/godotenv: Thư viện đọc file .env để nạp các biến môi trường vào ứng dụng Go giúp bảo mật cấu hình.
log: Thư viện chuẩn (built-in) của Go dùng để in thông báo, lỗi hoặc trạng thái ra màn hình console (terminal).
"crypto/rand": Thư viện cung cấp bộ tạo số ngẫu nhiên an toàn về mặt mật mã (dùng để sinh dữ liệu ngẫu nhiên chống đoán trước).
"strconv": Thư viện hỗ trợ chuyển đổi kiểu dữ liệu (ở đây dùng để chuyển chuỗi số thành số nguyên int).

# --- stack
1. Fatalf
log.Fatalf: Dùng để in thông báo lỗi ra màn hình và tức thì dừng (crash) chương trình (tương đương gọi os.Exit(1)), thường dùng khi ứng dụng gặp lỗi nghiêm trọng ngay từ bước khởi động (như không kết nối được database hoặc không mở được cổng server).

2. ID uint `json:"id"`
Đổi tên trường khi làm việc với JSON: Giúp biến viết hoa trong Go (ID) tự động chuyển thành chữ thường (id) khi trả về response cho client hoặc khi nhận dữ liệu từ request JSON.

3. (*domain.User, error)
Ý nghĩa thực tế: Thay vì bắt Go sao chép toàn bộ một cục dữ liệu lớn (struct) đem đi trả về, hàm chỉ trả về một con số nhỏ gọn (địa chỉ ô nhớ). Nơi nào gọi hàm sẽ cầm "địa chỉ" này để đi đến ô nhớ đó lấy giá trị ra dùng.

4. (*domain.User) va (domain.User)
findUserByEmail, GetProfile: Đây là các hàm đi tìm kiếm dữ liệu trong Database.
Nếu tìm thấy, nó trả về thông tin user. Nhưng nếu không tìm thấy, nó cần phải trả về nil, nil (không có user nào cả và không có lỗi). Chỉ có kiểu con trỏ mới cho phép trả về giá trị nil này.

GetVerificationCode(e domain.User) hoặc CreateCart(..., u domain.User):
struct domain.User được truyền vào với vai trò là dữ liệu đầu vào (input parameter) hoặc đối tượng đã chắc chắn tồn tại để xử lý nghiệp vụ tiếp theo (ví dụ: đã đăng nhập, đã xác thực ID).
Không cần kiểm tra nil ở đây vì hệ thống đã có sẵn thông tin user đó rồi, truyền thẳng giá trị trực tiếp cho nhanh gọn và an toàn (tránh việc hàm bên trong vô tình làm thay đổi dữ liệu gốc của user ở bên ngoài).

5. interface{} và []interface{}
interface{} (Kiểu giao diện rỗng / Any)
Bản chất: Nó là một kiểu dữ liệu tổng quát có thể chứa bất kỳ loại dữ liệu nào trong Go (tương tự kiểu any trong TypeScript hoặc Object trong Java/C#).
Dùng để làm gì: Khi bạn viết một hàm mà chưa biết trước tham số truyền vào là kiểu gì (string, int, struct,...).

[]interface{} (Slice chứa các phần tử bất kỳ)
Bản chất: Nó là một mảng động (slice) mà mỗi phần tử bên trong mảng đó có kiểu là interface{}.
Dùng để làm gì: Khi bạn muốn trả về hoặc lưu trữ một danh sách các phần tử có kiểu dữ liệu khác nhau hoặc chưa xác định (ví dụ: một mảng vừa chứa số, vừa chứa chữ, vừa chứa object).

# --- more

1. Go ORM
GO ORM (Object-Relational Mapping) là một kỹ thuật và tập hợp các thư viện giúp lập trình viên ánh xạ các bảng trong cơ sở dữ liệu quan hệ (như MySQL, PostgreSQL, SQLite) thành các cấu trúc dữ liệu (struct) trong ngôn ngữ lập trình Go (Golang).

Thay vì phải viết các câu lệnh SQL thủ công (Raw SQL) dài dòng, ORM cho phép bạn thao tác với cơ sở dữ liệu hoàn toàn thông qua các đoạn code Go tự nhiên (như thêm, sửa, xóa, truy vấn dữ liệu thông qua các phương thức đối tượng).

Các tính năng chính của GO ORM
Ánh xạ mô hình (Model Mapping): Tự động chuyển đổi các Go Struct thành các bảng cơ sở dữ liệu và ngược lại.
Truy vấn an toàn và linh hoạt: Cung cấp các phương thức dạng chuỗi (chaining methods) để viết câu lệnh điều kiện, sắp xếp, phân trang mà không lo bị lỗi cú pháp SQL.
Quản lý mối quan hệ (Associations): Dễ dàng xử lý các mối quan hệ giữa các bảng như One-to-One, One-to-Many, và Many-to-Many.
Tự động migrate (Auto Migration): Tự động tạo bảng hoặc cập nhật cấu trúc bảng (thêm/sửa cột) trực tiếp từ code Go.
Hỗ trợ Transaction: Quản lý các giao dịch cơ sở dữ liệu an toàn để đảm bảo tính toàn vẹn dữ liệu.

Các thư viện ORM phổ biến nhất trong Go
Gorm (gorm.io):
Ưu điểm: Cú pháp cực kỳ thân thiện, hỗ trợ đầy đủ tính năng, cộng đồng lớn, có sẵn nhiều plugin hỗ trợ (cache, soft delete, v.v.).

2. Twilio/SNS/SES cho SMS
https://www.twilio.com/en-us
Twilio: Chuyên dùng để gửi SMS (và gọi điện, WhatsApp). Đây là dịch vụ bên thứ ba phổ biến nhất thế giới cho việc gửi tin nhắn OTP, thông báo.

AWS SNS (Simple Notification Service): Dịch vụ của Amazon dùng để gửi SMS (và thông báo đẩy push notification, gửi qua email, v.v.). Thường được các hệ thống chạy trên AWS sử dụng để gửi mã OTP tiết kiệm.

AWS SES (Simple Email Service): Dịch vụ của Amazon chuyên dùng để gửi Email (như email kích hoạt tài khoản, newsletter, v.v.), không dùng cho SMS.


3. stripe
https://stripe.com/
https://docs.stripe.com/testing
Stripe (bạn gõ nhầm thành strtip) là một nền tảng cổng thanh toán trực tuyến (Payment Gateway) cực kỳ phổ biến và mạnh mẽ trên thế giới.

Thanh toán trực tuyến: Cho phép khách hàng mua hàng thanh toán bằng thẻ tín dụng/ghi nợ quốc tế (Visa, MasterCard, JCB, American Express) hoặc các ví điện tử (Apple Pay, Google Pay).

Quản lý Marketplace (Sàn thương mại điện tử đa người bán): Thông qua tính năng Stripe Connect, bạn có thể tự động hóa việc nhận tiền từ khách hàng, giữ tiền hoa hồng của sàn, và chuyển phần tiền còn lại trực tiếp vào tài khoản ngân hàng của từng Seller (rất phù hợp với chức năng BecomeSeller và quản lý tài khoản ngân hàng trong code của bạn).

Thu phí định kỳ (Subscriptions): Dùng cho các mô hình kinh doanh dạng gói thuê bao tháng/năm (như Netflix, Spotify hoặc các dịch vụ SaaS).

Bảo mật và chống gian lận: Stripe tích hợp sẵn các công cụ bảo mật cao cấp (như Stripe Radar) để phát hiện giao dịch giả mạo, bảo vệ người bán và tuân thủ các tiêu chuẩn bảo mật tài chính quốc tế (PCI-DSS).


