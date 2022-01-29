import System from 'App/Models/System'

type Get = {
  name: string
}

type Insert = {
  id: string
  name: string
  arch: string | null
}

type Update = {
  id: string
  count: number
}

export default new class Database {
  public get = ({ name }: Get) => System.query()
    .select('*')
    .where('name', name)
    .first()
    .then(system => system?.serialize())
  
  public getMultiple = () => System.query()
    .select('*')
    .exec()
    .then(systems => systems.map(system => system.serialize()))

  public insert = (payload: Insert) => System.create(payload)
    .then(system => system?.serialize())

  public update = ({ id, count }: Update) => System.query()
    .where('id', id)
    .update({
      count,
    })
    .first()
}