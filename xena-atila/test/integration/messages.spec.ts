import test from 'japa'
import supertest from 'supertest'

import { BASE_URL } from '../config'

test.group('Integration - Messages', () => {
  test('Assure that no messages exist at first.', async (assert) => {
    const { body: maybeMessages } = await supertest(BASE_URL)
      .get('/messages')
      .expect(204)
    
    assert.deepEqual(maybeMessages, {})
  })

  /* 
  // To be continued...
  test('Assure that message can be exchanged between clients.', async (assert) => {
    // Insert a client A.
    const clientAId = uuidv4()
    const { body: maybeInsertedClientA } = await supertest(BASE_URL)
      .post('/clients')
      .send({
        id: clientAId,
        status: 'ALIVE',
      })
      .expect(200)

    assert.isString(maybeInsertedClientA.id)
    assert.isString(maybeInsertedClientA.status)
    
    // Insert a client B.
    const clientBId = uuidv4()
    const { body: maybeInsertedClientB } = await supertest(BASE_URL)
      .post('/clients')
      .send({
        id: clientBId,
        status: 'ALIVE',
      })
      .expect(200)

    assert.isString(maybeInsertedClientB.id)
    assert.isString(maybeInsertedClientB.status)
  })
  */
})
