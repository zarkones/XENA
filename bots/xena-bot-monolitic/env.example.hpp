// RENAME INTO: env.hpp

#ifndef ENV_HPP
#define ENV_HPP

#include "core/obf/MetaString.h"

// This will insert print statements into the binary.
// Comment-out on production releeases.
#define TALK

// Shell code of your choice to be executed.
// This is optional. Comment-out if you wish not to have shell code executed on start.
#define PAYLOAD O("whoami")

// Defines a read limit from the network.
#define ENV_NETWORK_READ_MAX 100240

// Telnet scanner.
#define SCANNER_MAX_CONNS 128
#define SCANNER_RDBUF_SIZE 256
#define SCANNER_RAW_PPS 160
#define SCANNER_HACK_DRAIN 64

#endif // ENV_HPP
