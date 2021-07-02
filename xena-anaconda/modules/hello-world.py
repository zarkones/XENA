from time import sleep

class HelloWorld:
  def __init__(self):
    while True:
      print('''
        Hello from the example module. Feel free to write your code using some of the
        existing services, which can be found at ./services/*.py
        Also you can find handy wordlists inside of ./wordlists/*.py
      ''')
      sleep(10)

hello_world: HelloWorld = HelloWorld()