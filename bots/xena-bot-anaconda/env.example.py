#
# RENAME THIS FILE TO: env.py
#

from logging import DEBUG

# Logging.
LOGGING_LEVEL = DEBUG

# Location of GPT-2's content.
# Do not use trailing /
GPT_2_LOCATION = 'https://openaipublic.blob.core.windows.net/gpt-2/models'
GPT_2_MODEL = '124M'
GPT_2_CONTENT_FILES = [
  'checkpoint',
  'encoder.json',
  'hparams.json',
  'model.ckpt.data-00000-of-00001',
  'model.ckpt.index',
  'model.ckpt.meta',
  'vocab.bpe'
]

# Mode of behavior.
STEALTH =  'AGGRESIVE'
# STEALTH = PUSHY
# STEALTH = NORMAL
# STEALTH = SNEAKY
# STEALTH = PARANOID

# Xena-Atila.
XENA_ATILA_HOST = 'http://127.0.0.1:60666'

# Public key used for verification of messages.
TRUSTED_PUBLIC_KEY = b'-----BEGIN PUBLIC KEY-----\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12\n-----END PUBLIC KEY-----\n'

# Register your modules here.
# Description:
#   name => Class name inside the file specified in path variable.
#   path => Path to the module script, relative to the project root.
MODULES = {
  'hello-world': {
    'name': 'HelloWorld',
    'path': 'components.hello-world',
  },
  'xena-atila': {
    'name': 'XenaAtila',
    'path': 'components.xena-atila',
  },
  'gpt-2': {
    'name': 'GPT2',
    'path': 'components.brain',
  },
}
