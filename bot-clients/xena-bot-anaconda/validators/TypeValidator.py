import logging

from typing import Union

class TypeValidator:
  @staticmethod
  def valid_string(data, exception: str) -> Union[str, None]:
    if isinstance(data, str):
      return data

    logging.warning(exception)
    
    return None