import paho.mqtt.client as mqtt
import json
import time,datetime
payload=""
def on_connect(client, userdata, flags, rc):
    print(f"Recieving data from Raspberry-pi client...")
    # subscribe, which need to put into on_connect
    # if reconnect after losing the connection with the broker, it will continue to subscribe to the raspberry/topic topic
    client.subscribe("IOTDATA-Topic")
# the callback function, it will be triggered when receiving messages
def on_message(client, userdata, msg):
    print(msg.payload.decode("utf-8"))
    payload=msg.payload.decode("utf-8")
    client1 = mqtt.Client() #create new instance
    client1.on_connect=on_connect_to_client #attach function to callback
    client1.connect(broker_address_client, keepalive=60, bind_address="") #connect to broker
    print('sending data to the Client...')
    if payload is not None:
      print(payload)
    else:
      print('Failed to get reading. Try again!')
  	# the four parameters are topic, sending content, QoS and whether retaining the message respectively
    client1.publish('IOTDATA-Topic',payload=payload, qos=0, retain=False)
    time.sleep(5)
def on_connect_to_client(client, userdata, flags, rc):
    if rc == 0:
        print("Connected success")
    else:
        print("Connected fail with code")
broker_address="192.168.0.107"
broker_address_client="192.168.0.108"
client = mqtt.Client()
client.on_connect = on_connect
client.on_message = on_message
# set the will message, when the Raspberry Pi is powered off, or the network is interrupted abnormally, it will send the will message to other clients
client.will_set('raspberry/status', b'{"status": "Off"}')
# create connection, the three parameters are broker address, broker port number, and keep-alive time respectively
client.connect(broker_address, keepalive=60)
# set the network loop blocking, it will not actively end the program before calling disconnect() or the program crash
client.loop_forever()
