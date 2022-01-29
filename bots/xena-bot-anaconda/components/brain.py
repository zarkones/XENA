import logging
import os

from env import GPT_2_LOCATION, GPT_2_MODEL, GPT_2_CONTENT_FILES
from requests import get
from tqdm import tqdm

class Brain:
  def __init__(self, gptLocation: str, content: list[str]) -> None:
    # Files used to build our model.
    self.gptLocation = gptLocation
    self.content = content

  @staticmethod
  def ignite(gptLocation: str = GPT_2_LOCATION, content: list[str] = GPT_2_CONTENT_FILES):
    return Brain(gptLocation, content)

  def create_model(self, model_name: str):
    logging.debug('Downloading a model: ' + model_name)

    subdir = os.path.join('models', model_name)
    if not os.path.exists(subdir):
        os.makedirs(subdir)
    # Needed because of Windows operating system.
    subdir = subdir.replace('\\','/')

    for file in self.content:
      response = get(self.gptLocation + '/' + model_name + '/' + file, stream = True)
      
      with open(os.path.join(subdir, file), 'wb') as f:
        file_size = int(response.headers['content-length'])
        chunk_size = 1000
        with tqdm(ncols = 100, desc = 'Fetching ' + file, total = file_size, unit_scale = True) as pbar:
          # 1k for chunk_size, since Ethernet packet size is around 1500 bytes.
          for chunk in response.iter_content(chunk_size = chunk_size):
            f.write(chunk)
            pbar.update(chunk_size)

Brain.ignite().create_model(GPT_2_MODEL)