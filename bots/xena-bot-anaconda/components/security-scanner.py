import re
import requests

TARGET = 'xena.network'

class SecurityScanner:
  def __init__ (self) -> None:
    search_results = self.duckit(TARGET)
  
  def duckit (self, term: str) -> list[str]:
    result = requests.post(
      url = 'https://html.duckduckgo.com/html',
      data = {
        'q': term,
      },
      headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0',
        'Content-Type': 'application/x-www-form-urlencoded',
        'Origin': 'https://html.duckduckgo.com',
        'Connection': 'close',
      },
    )

    return re.findall(r'result__url" href="(.*?)">', result.content.decode())

ss = SecurityScanner()