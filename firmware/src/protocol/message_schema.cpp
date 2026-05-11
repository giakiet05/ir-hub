#include "protocol/message_schema.hpp"

#include <cstdio>
#include <cstring>

namespace deskstation {
namespace protocol {

bool encodeIrEventJson(const IrEventPayload& payload,
                       char* outBuffer,
                       size_t outBufferSize) {
  if (outBuffer == nullptr || outBufferSize == 0) {
    return false;
  }

  const int written = std::snprintf(
      outBuffer,
      outBufferSize,
      "{\"raw_code\":\"0x%08lX\",\"address\":%u,\"command\":%u,\"is_repeat\":%s}",
      static_cast<unsigned long>(payload.rawCode),
      payload.address,
      payload.command,
      payload.isRepeat ? "true" : "false");

  return written > 0 && static_cast<size_t>(written) < outBufferSize;
}

} // namespace protocol
} // namespace deskstation
