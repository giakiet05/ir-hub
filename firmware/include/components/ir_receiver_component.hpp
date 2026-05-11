#pragma once

#include <cstdint>

#include "core/component.hpp"

namespace deskstation {
namespace core {
class Outbox;
}
namespace components {

class IRReceiverComponent : public core::IComponent {
public:
  IRReceiverComponent(uint8_t receivePin, core::Outbox& outbox);

  const char* name() const override;
  void begin() override;
  void tick(uint32_t nowMs) override;

private:
  uint8_t receivePin_;
  core::Outbox& outbox_;
  uint32_t sequence_ = 0;
  bool hasLastEvent_ = false;
  uint16_t lastAddress_ = 0;
  uint16_t lastCommand_ = 0;
  uint32_t lastRawCode_ = 0;
  uint32_t lastEmitMs_ = 0;
};

} // namespace components
} // namespace deskstation
