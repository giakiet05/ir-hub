# IR-Hub

A minimalist Infrared (IR) control hub that allows you to control your computer (media, volume, mouse) using a remote control. It consists of a firmware for RP2040 Zero and a background service written in Go for Linux.

## Features

- Minimalist and high-performance architecture.
- Automatic serial reconnection.
- Hot-reloading of IR key mapping (presets) without restarting the app.
- Native Linux input emulation via uinput (compatible with Wayland/GNOME).
- Local status LED control (Toggle via physical button).
- Systemd integration for background operation.

## Hardware Requirements

- **Microcontroller:** RP2040 Zero.
- **IR Receiver:** Connected to GPIO 29.
- **Status LED:** Built-in WS2812 RGB LED on GPIO 16.
- **Toggle Button:** Connected to GPIO 8 (using internal pull-up, connect to GND).

## Project Structure

- `firmware/`: C++ code for RP2040 Zero (PlatformIO).
- `app/`: Go background service for Linux.

## Setup Instructions

### 1. Firmware Installation

Navigate to the `firmware` directory and use PlatformIO to build and upload the code:

```bash
cd firmware
# Build and upload (requires pio CLI)
pio run -t upload
```

Note: This project uses the Earle Philhower RP2040 core for stability.

### 2. App Build

Navigate to the `app` directory and build the binary:

```bash
cd app
go mod tidy
go build -o ir-hub main.go
```

### 3. Deployment (Recommended)

To run the app reliably in the background, it is recommended to create a dedicated deployment folder (e.g., `~/ir-hub-deploy`). This keeps the executable and configuration separate from the source code.

```bash
# Create deployment directory
mkdir -p ~/ir-hub-deploy/presets

# Copy the built binary
cp app/ir-hub ~/ir-hub-deploy/

# Copy your preset files
cp app/presets/*.json ~/ir-hub-deploy/presets/

# Create and configure your .env file
cp app/.env.example ~/ir-hub-deploy/.env
# Now edit ~/ir-hub-deploy/.env with your actual serial path and settings
```

The deployment folder should look like this:
```text
~/ir-hub-deploy/
├── ir-hub (executable)
├── .env (configuration)
└── presets/
    └── your_preset.json
```

### 4. Permissions

Ensure your user has permission to access the serial port (usually `uucp` or `dialout` group):

```bash
# Check which group owns your serial port (e.g., /dev/ttyACM0)
ls -l /dev/ttyACM0
# Add your user to that group (example for uucp)
sudo usermod -aG uucp $USER
# LOGOUT and LOGIN again for changes to take effect
```

### 5. Systemd Service Setup

Create a service file at `/etc/systemd/system/ir-hub.service`: (remember to use your home dir path in the content below)

```ini
[Unit]
Description=IR-Hub Background Service
After=network.target

[Service]
User=your-username
Group=your-group
# Use the absolute path to your deployment folder
WorkingDirectory=/home/your-username/ir-hub-deploy
ExecStart=/home/your-username/ir-hub-deploy/ir-hub

Restart=always
RestartSec=3
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```

Enable and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable ir-hub
sudo systemctl start ir-hub
```

## Usage

- **Check Logs:** `journalctl -u ir-hub -f`
- **Presets:** Edit the JSON files in the `presets/` directory. The app will automatically reload them upon saving.
- **Local Control:** Press the button on GPIO 8 to toggle the status LED on the RP2040 Zero.

## License

MIT
