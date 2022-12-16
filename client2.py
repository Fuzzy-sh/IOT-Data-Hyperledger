import paho.mqtt.client as mqtt #import the client1
import time,datetime
import json
import Adafruit_DHT

# Set sensor type : Options are DHT11,DHT22 or AM2302
sensor=Adafruit_DHT.DHT11
# Set GPIO sensor is connected to
gpio=14

def on_connect(client, userdata, flags, rc):
    if rc == 0:
        print("Connected success")
    else:
        print(f"Connected fail with code {rc}")

broker_address="192.168.0.107"

client = mqtt.Client() #create new instance
client.on_connect=on_connect #attach function to callback
client.connect(broker_address, keepalive=60, bind_address="") #connect to broker
print('sending data to the Broker...')
# send a message to the raspberry/topic every 1 second, 5 times in a row
for i in range(1):
    # Use read_retry method. This will retry up to 15 times to
    # get a sensor reading (waiting 2 seconds between each retry).
    humidity, temperature = Adafruit_DHT.read_retry(sensor, gpio)
    c_date_time=datetime.datetime.now()
    c_time=c_date_time.strftime("%X")
    c_date=c_date_time.strftime("%x")
    #payload='Timestamp= {0}-{1} Temp={2:0.1f}*C  Humidity={3:0.1f}%'.format(c_date, c_time,temperature, humidity)
    payload='{0}-{1},{2:0.1f}*C,{3:0.1f}%'.format(c_date, c_time,temperature, humidity)

    # Reading the DHT11 is very sensitive to timings and occasionally
    # the Pi might fail to get a valid reading. So check if readings are valid.
    if humidity is not None and temperature is not None:
      print(payload)
    else:
      print('Failed to get reading. Try again!')
  	# the four parameters are topic, sending content, QoS and whether retaining the message respectively
    client.publish('IOTDATA-Topic', payload=json.dumps(payload), qos=0, retain=False)
    time.sleep(5)
client.loop_forever() #start the loop