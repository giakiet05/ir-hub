#pragma once

#include <cstddef>

namespace deskstation {
namespace core {

constexpr size_t kOutboxCapacity = 16;
constexpr size_t kOutboxMessageLength = 256;

struct OutboxMessage {
  char line[kOutboxMessageLength];
};

class Outbox {
public:
  bool push(const char* line);
  bool pop(OutboxMessage& outMessage);
  bool isEmpty() const;
  bool isFull() const;
  size_t size() const;
  void clear();

private:
  OutboxMessage buffer_[kOutboxCapacity];
  size_t head_ = 0;
  size_t tail_ = 0;
  size_t count_ = 0;
};

} // namespace core
} // namespace deskstation
