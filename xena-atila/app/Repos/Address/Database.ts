import Address from 'App/Models/Address'

type Get = {
  id: string
  banned: boolean
}

type GetMultiple = {
  page?: number
  banned: boolean
}

type Insert = {
  x: number
  y: number
  z: number
  w: number
  banned: boolean
}

class Database {
  public get = ({ id, banned }: Get) => Address.query()
    .select('*')
    .where('id', id)
    .if(banned, builder => builder.where('banned', banned as boolean))
    .first()
    .then(client => client?.serialize())

  public getMultiple = ({ page, banned }: GetMultiple) => Address.query()
    .select('*')
    .if(banned, builder => builder.where('banned', banned as boolean))
    .if(page, builder => builder.offset(page as number * 10))
    .if(page, builder => builder.limit(10))
    .exec()
    .then(clients => clients.map(c => c.serialize()))
  
  public insert = ({ x, y, z, w, banned }: Insert) => Address.create({
    x,
    y,
    z,
    w,
    banned,
  }).then(client => client.serialize())
}

export default new Database()