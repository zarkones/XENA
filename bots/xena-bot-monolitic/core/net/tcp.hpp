#ifndef TCP_HPP
#define TCP_HPP

#include <stdio.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <string>

#include "../../env.hpp"

#if defined(TALK)
#include <iostream>
#endif

class TCP {
  private:
    std::string address = "";
    unsigned int port;

  public:
    TCP (std::string _address, unsigned int _port) : address(_address), port(_port) {}

    std::string stream (const std::string data) {
      int sock = socket(AF_INET, SOCK_STREAM, 0);
      #if defined(TALK)
      if (sock < 0) {
        std::cout << "Socket creation failed." << std::endl;
      }
      #endif

      struct sockaddr_in serv_addr;
      serv_addr.sin_family = AF_INET;
      serv_addr.sin_port = htons(this->port);

      int bin_addr = inet_pton(AF_INET, this->address.c_str(), &serv_addr.sin_addr);
      #if defined(TALK)
      if (bin_addr <= 0) {
        std::cout << "Failed to transform IP address into binary format." << std::endl;
      }
      #endif

      int connection = connect(sock, (struct sockaddr *) &serv_addr, sizeof(serv_addr));
      if (connection < 0) {
        return O("CONNECTION_ERROR");
        #if defined(TALK)
        std::cout << "Connection refused." << std::endl;
        #endif
      }

      // Send data to the destination.
      send(sock, data.c_str(), strlen(data.c_str()), 0);

      // Receive data.
      char buffer [ENV_NETWORK_READ_MAX] = {0};
      int response = read(sock, buffer, ENV_NETWORK_READ_MAX);

      return buffer;
    }

  std::string port_str () {
    return std::to_string(this->port);
  }
};

#endif // TCP_HPP