#include "core/outbox.hpp"

#include <cstring>

namespace deskstation {
namespace core {

bool Outbox::push(const char* line) {
  if (line == nullptr || isFull()) {
    return false;
  }

  const size_t lineLength = std::strlen(line);
  if (lineLength >= kOutboxMessageLength) {
    return false;
  }

  std::strcpy(buffer_[tail_].line, line);
  tail_ = (tail_ + 1) % kOutboxCapacity;
  ++count_;
  return true;
}

bool Outbox::pop(OutboxMessage& outMessage) {
  if (isEmpty()) {
    return false;
  }

  outMessage = buffer_[head_];
  head_ = (head_ + 1) % kOutboxCapacity;
  --count_;
  return true;
}

bool Outbox::isEmpty() const { return count_ == 0; }

bool Outbox::isFull() const { return count_ == kOutboxCapacity; }

size_t Outbox::size() const { return count_; }

void Outbox::clear() {
  head_ = 0;
  tail_ = 0;
  count_ = 0;
}

} // namespace core
} // namespace deskstation
