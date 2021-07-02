import socket

import logging

class Sublister:
  # Return a list of reachable subdomains.
  def enumerate(self, domain: str, wordlist: list,  stop_at: int = 0) -> list:
    logging.debug('[A] Enumerating: ' + domain)
    
    # Ports to check.
    ports: list = [
      # Http
      80,
      # Https
      443,
      # Default port for Nuxt applications.
      3000,
      # Often used port for web services in development.
      8080
    ]

    # Our connection entity.
    client: socket = socket.socket(
      socket.AF_INET,
      socket.SOCK_STREAM
    )

    # Keeping track of how many subdomains we went trough.
    index: int = 0
    # List of reachable subdomains.
    alive: list = []

    # Enumerate subdomains.
    for word in wordlist:
      index += 1

      # Full subdomain name.
      subdomain: str = word + '.' + domain

      for port in ports:
        try:
          client.connect((
            subdomain,
            port
          ))

          alive.append(subdomain + ':' + str(port))

          logging.info('[+] Subdomain: ' + subdomain + ' is alive.')
        except Exception as e:
          logging.debug('[E]:')
          logging.debug(e)
          pass
      
      # Break, if we've reached our limit.
      if index >= stop_at and stop_at != 0:
        break

    return alive