from cv2 import VideoCapture

import logging

class Camera:
  # Camera device.
  take: VideoCapture = VideoCapture(0)

  # Takes a picture from a camera.
  def shoot(self):
    logging.debug('[+] Taking a camera picture.')
    buffer = self.take.read()
    return buffer