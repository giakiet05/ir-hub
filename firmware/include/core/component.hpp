#pragma once

#include <cstdint>

namespace deskstation {
namespace core {

class IComponent {
public:
  virtual ~IComponent() = default;

  virtual const char* name() const = 0;
  virtual void begin() = 0;
  virtual void tick(uint32_t nowMs) = 0;
};

} // namespace core
} // namespace deskstation
