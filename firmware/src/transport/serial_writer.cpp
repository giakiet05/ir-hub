#include "transport/serial_writer.hpp"

namespace deskstation {
namespace transport {

SerialWriter::SerialWriter(Stream& stream) : stream_(stream) {}

bool SerialWriter::writeLine(const char* line) {
  if (line == nullptr) {
    return false;
  }

  const size_t written = stream_.println(line);
  return written > 0;
}

bool SerialWriter::flush() {
  stream_.flush();
  return true;
}

} // namespace transport
} // namespace deskstation
