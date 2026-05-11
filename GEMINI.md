# Project: IR-Hub

## 1. Mục tiêu dự án
Tạo một thiết bị phần cứng (RP2040 Zero) cắm vào máy tính để nhận tín hiệu hồng ngoại (IR) và điều khiển máy tính (media, âm lượng, chuột) thông qua ứng dụng Go chạy ngầm. Dự án tập trung vào sự tối giản, ổn định và hiệu năng cao.

## 2. Kiến trúc hệ thống
Gồm 2 phần giao tiếp qua Serial (JSON format):

### A. `firmware/` (RP2040 Zero)
- **Phần cứng:** RP2040 Zero. Chân IR: GPIO 29. LED trạng thái: GPIO 16 (WS2812). Nút bấm toggle LED: GPIO 8.
- **Nền tảng:** PlatformIO (Arduino framework, Earle Philhower core).
- **Nhiệm vụ:** Đọc mã IR, gửi payload JSON (`raw_code`, `address`, `command`, `is_repeat`) lên PC. Quản lý LED trạng thái local.

### B. `app/` (Go Background Service)
- **Nhiệm vụ:**
  - Kết nối Serial với MCU (hỗ trợ Auto-reconnect).
  - Parse mã IR và map với các hành động (Actions) được định nghĩa trong file JSON preset.
  - Giả lập bàn phím/chuột qua `uinput` (tương thích Linux/Wayland).
  - Tự động load lại preset khi file JSON thay đổi (Hot-reload).
- **Cấu hình:** Sử dụng biến môi trường (`.env`):
  - `BY_ID_PATH`: Đường dẫn serial device.
  - `BAUD_RATE`: Tốc độ serial (thường là 115200).
  - `PRESET_DIR`: Thư mục chứa các file JSON mapping.

## 3. Trạng thái hiện tại
- Đã tối giản hóa hoàn toàn: Loại bỏ Wails (GUI), Event Bus rườm rà.
- Hệ thống chạy cực nhẹ, kết nối trực tiếp từ Serial Transport sang Handler.
- Hỗ trợ xử lý Signal (SIGINT, SIGTERM) để tắt app an toàn.

## 4. Hướng dẫn vận hành
- Build app: `cd app && go build -o ir-hub main.go`
- Chạy qua Systemd: Tạo service file trỏ vào file thực thi và thư mục làm việc chứa `.env`.
- Log: Xem qua `journalctl -u ir-hub -f`.
