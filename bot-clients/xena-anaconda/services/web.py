import logging

from requests import get

# Used for parsing html.
from bs4 import BeautifulSoup
from bs4 import ResultSet

class Web:
  # Returns found links from a web page.
  def get_hyper_links(self, url: str = '', text: str = '', scope: bool = False) -> list:
    logging.debug('Crawling ' + url)

    links: list = []

    # Grab values that may be links.
    for tag in ['a', 'link', 'meta', 'script', 'form']:
      links += self.page_grab_tag(
        url = url,
        tag = tag,
        text = text,
        props = [
          'href',
          'src',
          'content',
          'action'
        ]
      )
    
    # Remove duplicates.
    links = list(set(links))

    # Filter out non-links.
    tmp_links: list = []
    for link in links:
      if link.startswith('http'):
        if scope == True:
          if link.startswith(url):
            tmp_links.append(link)
        else:
            tmp_links.append(link)
      if link.startswith('/'): 
        tmp_links.append(link)
    links = tmp_links

    return links
  
  def unwrap(self, string: str) -> str:
    return string.split('>', 1)[1].split('<', 1)[0]

  # Returns a list of content.
  # If 'props' argument is left unset, the function will return a list of tags.
  # Otherwise, it will fetch values of thos properites and return them as a list.
  def page_grab_tag(self, url: str = '', tag: str = '', props: list = [], text: str = '', value: str = None) -> list:
    logging.debug('[+] Grabbing all hyper links from: ' + url)

    # Tags to be returned.
    tags: list = []

    # Get the web page.
    http_response = None
    soup: BeautifulSoup = None
    if url != '' and text == '':
      http_response = get(url)
      soup = BeautifulSoup(http_response.text, 'html.parser')
    if text != '':
      soup = BeautifulSoup(text, 'html.parser')

    if value != None:
      found_by_tag: ResultSet = soup.findAll(
        tag,
        {
          props[0]: value
        }
      )
      for t in found_by_tag:
        tags.append(str(t))
      #tags.append(str(found_by_tag))

      return tags

    # Find every tag specific tag.
    found_by_tag: ResultSet = soup.findAll(tag)

    # If no properites are specified, return all found by a tag.
    if props is []:
      # Build up the tag dict.
      for t in found_by_tag:
        tags.append(str(t))

      return tags

    # If 'props' argument is specified,
    # we enumerate all tags for those attributes and return their values.
    tags_with_props: list = []

    for t in found_by_tag:
      for prop in props:
        if t.get(prop) is not None:
          tags_with_props.append(t.get(prop))

    return tags_with_props