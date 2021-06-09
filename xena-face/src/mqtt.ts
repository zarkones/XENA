const mqtt = require('mqtt')

export default class MqttListener {

  constructor () {
    const client  = mqtt.connect('wss://127.0.0.1:880')

    client.on('connect', function () {
      console.log('connected.')

      client.subscribe('xena/network/presence', function (e) {
        if (e)
          return MqttListener.mqtt_error(e)

        client.publish('xena/network/presence', JSON.stringify({
          id: process.env.APP_UUID,
        }))
      })

      client.subscribe(`xena/network/entity/${process.env.APP_UUID}`, function (e) {
        if (e)
          return MqttListener.mqtt_error(e)
      })

      client.subscribe(`xena/network/entity/${process.env.XENA_KERNEL_UUID}`, function (e) {
        if (e)
          return MqttListener.mqtt_error(e)
      })
    })
     
    client.on('message', function (topic: string, message: Buffer) {
      console.log(message.toString())
      client.end()
    })
  }

  private static mqtt_error (error) {
    console.log({
      id: process.env.XENA_KERNEL_UUID,
      message: 'An error occured on connect.',
      error,
    })
  }
}