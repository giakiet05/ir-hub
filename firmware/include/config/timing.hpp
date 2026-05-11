#pragma once

#include <cstdint>

#ifndef DS_LOOP_DELAY_MS
#define DS_LOOP_DELAY_MS 1
#endif

#ifndef DS_BUTTON_DEBOUNCE_MS
#define DS_BUTTON_DEBOUNCE_MS 30
#endif

#ifndef DS_DHT11_PUBLISH_INTERVAL_MS
#define DS_DHT11_PUBLISH_INTERVAL_MS 5000
#endif

#ifndef DS_HEALTH_PUBLISH_INTERVAL_MS
#define DS_HEALTH_PUBLISH_INTERVAL_MS 30000
#endif

namespace deskstation {
namespace config {

constexpr uint32_t kMainLoopDelayMs = DS_LOOP_DELAY_MS;
constexpr uint32_t kButtonDebounceMs = DS_BUTTON_DEBOUNCE_MS;
constexpr uint32_t kDht11PublishIntervalMs = DS_DHT11_PUBLISH_INTERVAL_MS;
constexpr uint32_t kHealthPublishIntervalMs = DS_HEALTH_PUBLISH_INTERVAL_MS;

} // namespace config
} // namespace deskstation
