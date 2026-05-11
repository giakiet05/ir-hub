#include "components/ir_receiver_component.hpp"

#include <Arduino.h>
#include <IRremote.hpp>

#include "core/outbox.hpp"
#include "protocol/message_schema.hpp"

namespace deskstation {
namespace components {

namespace {
constexpr uint32_t kIrDuplicateWindowMs = 90;
constexpr uint32_t kIrRepeatMinIntervalMs = 90;
}

IRReceiverComponent::IRReceiverComponent(uint8_t receivePin, core::Outbox& outbox)
    : receivePin_(receivePin), outbox_(outbox) {}

const char* IRReceiverComponent::name() const { return "ir_receiver"; }

void IRReceiverComponent::begin() { IrReceiver.begin(receivePin_, DISABLE_LED_FEEDBACK); }

void IRReceiverComponent::tick(uint32_t nowMs) {
  if (!IrReceiver.decode()) {
    return;
  }

  const IRData& data = IrReceiver.decodedIRData;
  const uint16_t address = data.address;
  const uint16_t command = data.command;
  const uint32_t rawCode =
      static_cast<uint32_t>(static_cast<uint64_t>(data.decodedRawData) & 0xFFFFFFFFULL);

  bool isRepeat = false;
#ifdef IRDATA_FLAGS_IS_REPEAT
  isRepeat = (data.flags & IRDATA_FLAGS_IS_REPEAT) != 0;
#endif

  if (hasLastEvent_ && address == lastAddress_ && command == lastCommand_ &&
      rawCode == lastRawCode_) {
    const uint32_t elapsedMs = nowMs - lastEmitMs_;
    if (!isRepeat && elapsedMs < kIrDuplicateWindowMs) {
      IrReceiver.resume();
      return;
    }
    if (isRepeat && elapsedMs < kIrRepeatMinIntervalMs) {
      IrReceiver.resume();
      return;
    }
  }

  protocol::IrEventPayload payload{
      rawCode,
      address,
      command,
      isRepeat,
  };

  char line[protocol::kMaxJsonLineLength] = {};
  if (protocol::encodeIrEventJson(payload, line, sizeof(line))) {
    outbox_.push(line);
    hasLastEvent_ = true;
    lastAddress_ = address;
    lastCommand_ = command;
    lastRawCode_ = rawCode;
    lastEmitMs_ = nowMs;
  }

  IrReceiver.resume();
}

} // namespace components
} // namespace deskstation
