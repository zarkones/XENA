class Validators:
  @staticmethod
  def valid_string(data, exception: str):
    if isinstance(data, str):
      return data
    raise TypeError(exception)