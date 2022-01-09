import logging

from env import MODULES, LOGGING_LEVEL
from multiprocessing import Process
from importlib import import_module

class Main:
  def __init__ (self) -> None:
    self.logging_init()

    # Starting our processes.
    modules: dict = MODULES
    for name in modules:
      module: dict = modules[name]
      try:
        Process(
          target = self.module_init,
          args = (module, )
        ).start()
      except Exception as e:
        logging.warning('Could not start the module: ' + name)
        logging.exception(e)

  # Handles our processes.
  def module_init (self, module: dict) -> None:
    try:
      logging.debug('Module imported: ' + module['name'])
      import_module(module['path'])
      logging.debug('Module exited: ' + module['name'])
    except Exception as e:
      logging.warning('Could not import the module: ' + module['name'])
      logging.exception(e)

  # Our logging configuration.
  def logging_init (self) -> None:
    logging.basicConfig(
      format = '%(asctime)s: %(message)s',
      level = LOGGING_LEVEL,
      datefmt = '%H:%M:%S'
    )

if __name__ == '__main__':
  main: Main = Main()