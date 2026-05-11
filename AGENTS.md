# IR-Hub Project Guidelines

Dự án này là phiên bản tối giản của Desk Station, tập trung duy nhất vào tính năng điều khiển IR.

## Core Mandates
- **Minimalism:** Không thêm các thư viện hoặc layer trung gian (như Event Bus, GUI) nếu không thực sự cần thiết.
- **Direct Flow:** Dữ liệu đi thẳng từ Serial -> Parser -> Handler -> uinput.
- **Stability:** Luôn xử lý reconnection cho Serial và Graceful Shutdown cho ứng dụng.
- **Configuration:** Mọi thông số môi trường (cổng serial, đường dẫn preset) phải nằm trong `.env`.

## Tech Stack
- **Firmware:** C++, Arduino, Earle Philhower RP2040 Core, IRremote, Adafruit NeoPixel.
- **App:** Go 1.23+, uinput, godotenv, fsnotify.

## Deployment
- Ứng dụng được thiết kế để chạy như một systemd service trên Linux.
- Logs được ghi trực tiếp ra stdout/stderr để `journalctl` thu thập.
