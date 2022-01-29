import File from 'App/Models/File'

type Insert = {
  id: string
  originId: string
  originPath: string
  data: string
}

export default new class FileRepo {
  public insert = (payload: Insert) => File.create(payload)
    .then(file => file.serialize())
    .catch(e => console.error(e))
}