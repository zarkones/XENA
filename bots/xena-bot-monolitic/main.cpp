#include "core/net/request.hpp"
#include "core/obf/MetaString.h"

int main (int argc, char * argv[]) {
  Request req {
    O("127.0.0.1"),
    60666,
    O("/"),
    false,
    O("GET"),
    ""
  };
  req.exec();

  return 0;
}