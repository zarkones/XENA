import test from 'japa'
import supertest from 'supertest'

import { BASE_URL } from '../config'
import { v4 as uuidv4 } from 'uuid'

test.group('Integration - Clients', () => {
  test('Assure that no clients exist at first.', async (assert) => {
    const { body: maybeClients } = await supertest(BASE_URL)
      .get('/clients')
      .expect(204)
    
    assert.deepEqual(maybeClients, {})
  })

  test('Assure that a client can be inserted.', async (assert) => {
    // Insert a client.
    const { body: maybeInsertedClient } = await supertest(BASE_URL)
      .post('/clients')
      .send({
        id: uuidv4(),
        status: 'ALIVE',
      })
      .expect(200)

    assert.isString(maybeInsertedClient.id)
    assert.isString(maybeInsertedClient.status)

    // Read all clients.
    const { body: maybeClients } = await supertest(BASE_URL)
      .get('/clients')
      .expect(200)

    assert.isArray(maybeClients)
    assert.equal(maybeClients.length, 1)

    assert.isString(maybeClients[0].id)
    assert.isString(maybeClients[0].status)
  })
})
