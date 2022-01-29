from services.se import Search

TARGET = 'xena.network'

class SecurityScanner:
  def __init__ (self, target: str) -> None:
    search_results = Search.duck('site:' + target)

    for url in search_results:
      print(url)

ss = SecurityScanner(TARGET)