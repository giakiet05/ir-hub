#pragma once

#include <cstddef>
#include <cstdint>

namespace deskstation {
namespace protocol {

struct IrEventPayload {
  uint32_t rawCode;
  uint16_t address;
  uint16_t command;
  bool isRepeat;
};

constexpr size_t kMaxJsonLineLength = 256;

bool encodeIrEventJson(const IrEventPayload& payload,
                       char* outBuffer,
                       size_t outBufferSize);

} // namespace protocol
} // namespace deskstation
