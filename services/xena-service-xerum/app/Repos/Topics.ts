import Topic from 'App/Models/Topic'

type Get = {
  id: string
}

export default new class {
  public get = ({ id }: Get) => Topic.query()
    .where('id', id)
    .first()

  public getMultiple = () => Topic.all()
}