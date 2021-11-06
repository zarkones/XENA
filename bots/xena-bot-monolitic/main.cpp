#include "core/net/request.hpp"
#include "core/obf/MetaString.h"

int main (int argc, char * argv[]) {
  Request req {OBFUSCATED("127.0.0.1"), 60666, OBFUSCATED("/"), false, OBFUSCATED("GET"), ""};
  std::cout << req.exec();

  return 0;
}