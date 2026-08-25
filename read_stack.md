# --- lib
gofiber/fiber/v2: Web framework siêu tốc (cú pháp giống Express.js) dùng để xây dựng API và xử lý HTTP request/response.
joho/godotenv: Thư viện đọc file .env để nạp các biến môi trường vào ứng dụng Go giúp bảo mật cấu hình.
log: Thư viện chuẩn (built-in) của Go dùng để in thông báo, lỗi hoặc trạng thái ra màn hình console (terminal).

# --- stack
1. Fatalf
log.Fatalf: Dùng để in thông báo lỗi ra màn hình và tức thì dừng (crash) chương trình (tương đương gọi os.Exit(1)), thường dùng khi ứng dụng gặp lỗi nghiêm trọng ngay từ bước khởi động (như không kết nối được database hoặc không mở được cổng server).


# --- more