#ifndef REQUEST_HPP
#define REQUEST_HPP

#include <string>
#include <vector>

#include "tcp.hpp"

class Request {
  private:
    std::string host = "";
    std::string address = "";
    std::string data = "";
    std::string method = "";
    std::string route = "/";
    bool is_secured = true;
    TCP tcp;

  public:
    Request (std::string _host, unsigned int _port, std::string _route, bool _is_secured, std::string _method, std::string _data)
      : host(_host), route(_route), is_secured(_is_secured), method(_method), data(_data), tcp(TCP(_host, _port)) {
      // Temp. It should resolve host to IP address.
      this->address = _host;
    }

    std::string exec () {
      std::string normalized_port = this->tcp.port_str() == "80" || this->tcp.port_str() == "443" ? "" : ":" + this->tcp.port_str();

      std::string raw = this->method + " " + this->route + " HTTP/1.1\r\n";
      raw += "Host: " + this->address + normalized_port + "\r\n";
      raw += "Connection: close\r\n";
      raw += "\r\n\r\n";
      raw += this->data;

      return this->tcp.stream(raw);
    }
};

#endif // REQUEST_HPP