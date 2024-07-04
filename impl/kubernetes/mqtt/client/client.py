import os
import random
import socket
import time

import paho.mqtt.client as mqtt

HOST = os.environ.get("mqtt-host") or "34.101.95.240"
PORT  = os.environ.get("mqtt-port") or 1883
KEEPALIVE = os.environ.get("mqtt-keepalive") or 60


client = mqtt.Client()
client.connect(HOST, PORT, KEEPALIVE)

while True:
    t = time.localtime()
    localtime = time.strftime("[%D - %T]", t)
    hostname = socket.gethostname()
    random_temperature = random.randint(30, 42)
    message = f"{localtime} temperature from {hostname}: {random_temperature}"
    
    client.publish("iot-topic", message)
    print("Message sent: ", message)
    time.sleep(5)