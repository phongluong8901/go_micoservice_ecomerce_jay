# --- lib
gofiber/fiber/v2: Web framework siêu tốc (cú pháp giống Express.js) dùng để xây dựng API và xử lý HTTP request/response.
joho/godotenv: Thư viện đọc file .env để nạp các biến môi trường vào ứng dụng Go giúp bảo mật cấu hình.
log: Thư viện chuẩn (built-in) của Go dùng để in thông báo, lỗi hoặc trạng thái ra màn hình console (terminal).

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