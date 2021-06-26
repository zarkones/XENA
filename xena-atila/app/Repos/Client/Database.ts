import Client from 'App/Models/Client'

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
  identificationKey: string
  status: ClientStatus
}

class Database {
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
  
  public insert = ({ id, status }: Insert) => Client.create({
    id,
    status,
  }).then(client => client.serialize())
}

export default new Database()