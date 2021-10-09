import Client from 'App/Models/Client'
import LucidDatabase from '@ioc:Adonis/Lucid/Database'

type ClientStatus = 'ALIVE' | 'DEAD' | 'BANNED'

type Get = {
  id: string
  status?: ClientStatus
}

type GetMultiple = {
  page?: number
  status?: ClientStatus
}

type Insert = {
  id: string
  publicKey: string
  status: ClientStatus
}

export default new class Database {
  public get = ({ id, status }: Get) => Client.query()
    .select('*')
    .where('id', id)
    .whereNot('status', 'BANNED')
    .if(status, builder => builder.where('status', status as ClientStatus))
    .first()
    .then(client => client?.serialize())

  public getMultiple = ({ page, status }: GetMultiple) => Client.query()
    .select('*')
    .whereNot('status', 'BANNED')
    .if(status, builder => builder.where('status', status as ClientStatus))
    .if(page, builder => builder.offset(page as number * 10))
    .if(page, builder => builder.limit(10))
    .exec()
    .then(clients => clients.map(c => c.serialize()))

  public getCount = () => LucidDatabase.rawQuery('select extract(epoch from created_at) * 1000 as timestamp from clients')
    .exec()
    .then(result => result.rows.map(unixTime => unixTime.timestamp))
  
  public insert = (payload: Insert) => Client.create(payload).then(client => client.serialize())
}
