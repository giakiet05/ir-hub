# Project: Desk Station (tên cũ: control-hub)

## 1. Mục tiêu dự án
Tạo một thiết bị phần cứng cắm trực tiếp vào máy tính qua cổng USB (Type-C) để làm một trạm điều khiển/hiển thị đa năng trên bàn làm việc.
- **Tính năng hiện tại:** Nhận tín hiệu hồng ngoại (IR) từ remote để điều khiển media, âm lượng, chuột trên PC.
- **Tính năng mở rộng (tương lai):** Đo nhiệt độ, độ ẩm phòng (DHT22/BME280), hiển thị thông số, và bất cứ module nào có thể cắm thêm vào vi điều khiển.
- **Giao diện:** Một ứng dụng Desktop trên Linux (GNOME) có thể chạy ngầm dưới System Tray (taskbar), hiển thị thông số môi trường trên icon, và có giao diện (Dashboard) để map nút điều khiển IR thay vì hardcode.

## 2. Kiến trúc hệ thống
Dự án được chia làm 2 thành phần chính, giao tiếp với nhau qua cổng Serial (USB Native hoặc UART):

### A. `firmware/` (Mã nguồn Vi điều khiển)
- **Phần cứng:** Dùng **RP2040 Zero** (USB Native). Nguồn lấy trực tiếp từ cáp USB PC.
- **Nền tảng:** PlatformIO (C/C++).
- **Nhiệm vụ:** Đọc tín hiệu từ các linh kiện đang dùng (IR, DHT11, button, LED state), đóng gói theo JSON schema thống nhất và gửi lên PC qua Serial.

### B. `app/` (Ứng dụng Desktop chạy trên Linux)
- **Nền tảng:** Go + framework **Wails** (Backend Go, Frontend Web dùng Vanilla TypeScript).
- **Backend (Go):**
  - Chạy ngầm, liên tục mở cổng Serial ổn định (ưu tiên `/dev/serial/by-id/...`) để nhận data từ MCU.
  - Xử lý giả lập bàn phím và chuột ở mức OS bằng thư viện `uinput` (để vượt qua các hạn chế bảo mật của Wayland/GNOME).
  - Quản lý logic Tray Icon (System Tray) bằng tính năng native của Wails, cho phép app chạy ngầm khi đóng cửa sổ.
- **Frontend (Web/TS):**
  - Cung cấp giao diện đồ họa (Dashboard) mượt mà để hiển thị nhiệt độ, độ ẩm.
  - Cung cấp giao diện (Settings) để người dùng bấm remote, bắt mã HEX, và map vào các phím tắt tùy ý (lưu xuống file JSON config của backend).

## 3. Tình trạng hiện tại (Status)
- Đã cấu trúc lại thư mục thành `app/` và `firmware/`.
- Đã khởi tạo thành công Wails project (template `vanilla-ts`) bên trong thư mục `app/`.
- File code Go cũ thực hiện nhiệm vụ đọc Serial và uinput đang được lưu trữ tạm tại `app/old_main.go` chờ được tích hợp vào logic Backend của Wails (`app/app.go`).

## 4. Kế hoạch tiếp theo (Roadmap)
1. Bứng logic đọc Serial và `uinput` từ `app/old_main.go` vào Backend của Wails (`app.go`).
2. Cấu hình Wails để hỗ trợ System Tray (ẩn/hiện cửa sổ thay vì tắt hẳn).
3. Refactor lại code C++ và Go để chuẩn hóa chuỗi giao tiếp qua Serial thành JSON (chuẩn bị đón cảm biến nhiệt độ).
4. Xây dựng giao diện Frontend map nút.

## 5. Decisions (2026-05-04)

### 5.1 Current active components on breadboard
- IR receiver
- DHT11 module
- 3 single-color LEDs (red, yellow, green)
- 1 button
- Buzzer is removed from current scope

### 5.2 Firmware <-> App communication contract
- Keep USB Serial as transport.
- Use a **unified line-delimited JSON message schema** for all messages.
- Go app parses each message with `json.Unmarshal` into typed Go structs.
- Avoid ad-hoc string formats like `IR_CODE:0x...` for new flows.

### 5.3 Event-driven architecture boundary
- Event bus is implemented in the **Go app only**.
- Firmware is responsible for hardware I/O, state sampling, and publishing/subscribing JSON messages.
- App side routes events by topic/priority and processes them concurrently with goroutines.
