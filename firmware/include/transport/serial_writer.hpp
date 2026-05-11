#pragma once

#include <Arduino.h>

namespace deskstation {
namespace transport {

class SerialWriter {
public:
  explicit SerialWriter(Stream& stream);

  bool writeLine(const char* line);
  bool flush();

private:
  Stream& stream_;
};

} // namespace transport
} // namespace deskstation
