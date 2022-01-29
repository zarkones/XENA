#ifndef REQUEST_HPP
#define REQUEST_HPP

#include <cstring>
#include <string>
#include <vector>

#include "tcp.hpp"

#include "../obf/MetaString.h"

class Request {
  private:
    std::string host = "";
    std::string address = "";
    std::string data = "";
    std::string method = "";
    std::string route = O("/");
    bool is_secured = true;
    TCP tcp;

  public:
    Request (std::string _host, unsigned int _port, std::string _route, bool _is_secured, std::string _method, std::string _data)
      : host(_host), route(_route), is_secured(_is_secured), method(_method), data(_data), tcp(TCP(_host, _port)) {
      // Temp. It should resolve host to IP address.
      this->address = _host;
    }

    std::string exec () {
      std::string normalized_port = this->tcp.port_str() == O("80") || this->tcp.port_str() == O("443") ? "" : O(":") + this->tcp.port_str();

      std::string raw = this->method + O(" ") + this->route + O(" HTTP/1.1\r\n");
      raw += O("Host: ") + this->address + normalized_port + O("\r\n");
      raw += O("Connection: close\r\n");
      raw += O("Content-Type: application/json\r\n");
      raw += O("Content-Length: ");
      raw += std::to_string(strlen(this->data.c_str()));
      raw += O("\r\n\r\n");
      raw += this->data;

      return this->tcp.stream(raw);
    }
};

#endif // REQUEST_HPP