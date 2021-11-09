#
# RENAME THIS FILE TO: env.py
#

from logging import DEBUG

def Env() -> dict:
  return {
    # Logging.
    'LOGGING_LEVEL': DEBUG,

    # Xena-Atila.
    'XENA_ATILA_HOST': 'http://127.0.0.1:60666',

    'MASTER_PUBLIC_KEY': b'-----BEGIN PUBLIC KEY-----\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12\n-----END PUBLIC KEY-----\n',

    # Register your modules here.
    # Description:
    #   name => Class name inside the file specified in path variable.
    #   path => Path to the module script, relative to the project root.
    'MODULES': {
      'hello-world': {
        'name': 'HelloWorld',
        'path': 'modules.hello-world',
      },
      
      'xena-atila': {
        'name': 'XenaAtila',
        'path': 'modules.xena-atila',
      },
    }
  }