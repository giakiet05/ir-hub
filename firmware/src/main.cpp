#include <Arduino.h>

namespace deskstation {
void appSetup();
void appLoop();
} // namespace deskstation

void setup() { deskstation::appSetup(); }

void loop() { deskstation::appLoop(); }
