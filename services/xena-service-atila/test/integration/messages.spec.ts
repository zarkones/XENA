import test from 'japa'
import supertest from 'supertest'

import { BASE_URL } from '../config'
import { v4 as uuidv4 } from 'uuid'

test.group('Integration - Messages', () => {
  test('Assure that no messages exist at first.', async (assert) => {
    const { body: maybeMessages } = await supertest(BASE_URL)
      .get('/messages')
      .expect(204)
    
    assert.deepEqual(maybeMessages, {})
  })

  test('Assure that message can be exchanged.', async (assert) => {
    // Insert a client A.
    const { body: client } = await supertest(BASE_URL)
      .post('/clients')
      .send({
        id: uuidv4(),
        publicKey: 'fakepublickey',
        status: 'ALIVE',
      })
      .expect(200)

    console.log(client)

    assert.isString(client.id)
    assert.isString(client.status)

    const { body: message } = await supertest(BASE_URL)
      .post('/messages')
      .send({
        
      })
  })
})
