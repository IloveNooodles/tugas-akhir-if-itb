import RPi.GPIO as GPIO
import time
import os

PIN_USED = int(os.environ.get('GPIO_LED_PIN') or 2)
BLINK_INTERVAL = float(os.environ.get('BLINK_INTERVAL') or 0.5)

GPIO.setmode(GPIO.BCM)
GPIO.setwarnings(False)
GPIO.setup(PIN_USED, GPIO.OUT)

while True:
 GPIO.output(PIN_USED, True)
 time.sleep(BLINK_INTERVAL)
 GPIO.output(PIN_USED, False)
 time.sleep(BLINK_INTERVAL)

GPIO.output(PIN_USED, False)
GPIO.cleanup()