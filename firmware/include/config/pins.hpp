#pragma once

#include <cstdint>

#ifndef DS_PIN_IR
#define DS_PIN_IR 29
#endif

#ifndef DS_PIN_STATUS_LED
#define DS_PIN_STATUS_LED 16
#endif

#ifndef DS_PIN_TOGGLE_BUTTON
#define DS_PIN_TOGGLE_BUTTON 8
#endif

namespace deskstation {
namespace config {

constexpr uint8_t kIrReceivePin = DS_PIN_IR;
constexpr uint8_t kStatusLedPin = DS_PIN_STATUS_LED;
constexpr uint8_t kToggleButtonPin = DS_PIN_TOGGLE_BUTTON;

} // namespace config
} // namespace deskstation
