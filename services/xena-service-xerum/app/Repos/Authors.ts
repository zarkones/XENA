import Author from 'App/Models/Author'

type Insert = {
  id: string
  name: string
  publicKey: string
}

type Get = {
  id: string
}

export default new class {
  public insert = (payload: Insert) => Author.create(payload)

  public get = ({ id }: Get) => Author.query()
    .where('id', id)
    .first()
  
  public getByName = (name: string) => Author.query()
    .where('name', name)
    .first()
}