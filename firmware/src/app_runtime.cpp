#include <Adafruit_NeoPixel.h>
#include <Arduino.h>
#include <cstdio>

#include "components/ir_receiver_component.hpp"
#include "config/pins.hpp"
#include "config/timing.hpp"
#include "core/component.hpp"
#include "core/outbox.hpp"
#include "protocol/message_schema.hpp"
#include "transport/serial_writer.hpp"

namespace deskstation {
namespace {

core::Outbox gOutbox;
transport::SerialWriter gSerialWriter(Serial);

components::IRReceiverComponent gIrComponent(config::kIrReceivePin, gOutbox);

core::IComponent* gComponents[] = {
    &gIrComponent,
};

// Built-in WS2812 RGB LED for RP2040 Zero
Adafruit_NeoPixel gStatusLed(1, config::kStatusLedPin, NEO_GRB + NEO_KHZ800);

// Local LED State
bool gLedOn = true;

// Button state for debouncing
bool gLastButtonState = HIGH;
uint32_t gLastDebounceTime = 0;
constexpr uint32_t kDebounceDelayMs = 50;

void updateLed() {
  if (gLedOn) {
    gStatusLed.setPixelColor(0, gStatusLed.Color(0, 255, 0)); // Green
  } else {
    gStatusLed.setPixelColor(0, gStatusLed.Color(0, 0, 0)); // Off
  }
  gStatusLed.show();
}

void handleLocalButton(uint32_t nowMs) {
  const bool reading = digitalRead(config::kToggleButtonPin);

  if (reading != gLastButtonState) {
    gLastDebounceTime = nowMs;
  }

  if ((nowMs - gLastDebounceTime) > kDebounceDelayMs) {
    static bool buttonState = HIGH;
    if (reading != buttonState) {
      buttonState = reading;
      // Button pressed (assuming pull-up, so LOW means pressed)
      if (buttonState == LOW) {
        gLedOn = !gLedOn;
        updateLed();
      }
    }
  }

  gLastButtonState = reading;
}

void flushOutbox() {
  core::OutboxMessage message{};
  while (gOutbox.pop(message)) {
    gSerialWriter.writeLine(message.line);
  }
  gSerialWriter.flush();
}

} // namespace

void appSetup() {
  Serial.begin(115200);

  // Setup Button (INPUT_PULLUP)
  pinMode(config::kToggleButtonPin, INPUT_PULLUP);

  // Initialize Status LED
  gStatusLed.begin();
  gStatusLed.setBrightness(20);
  updateLed();

  const uint32_t startWaitMs = millis();
  while (!Serial && (millis() - startWaitMs) < 1500U) {
    delay(1);
  }

  for (core::IComponent* component : gComponents) {
    component->begin();
  }

  flushOutbox();
}

void appLoop() {
  const uint32_t nowMs = millis();

  handleLocalButton(nowMs);

  for (core::IComponent* component : gComponents) {
    component->tick(nowMs);
  }

  flushOutbox();
  delay(config::kMainLoopDelayMs);
}

} // namespace deskstation
