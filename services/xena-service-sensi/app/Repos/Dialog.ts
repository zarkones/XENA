import Dialog from 'App/Models/Dialog'

type Insert = {
  id: string
  input: string
  output: string | null
}

export default new class {
  public getMultiple = () => Dialog.query()
    .select('*')
    .exec()
    .then(dialog => dialog.map(d => d.serialize()))

  public insert = (payload: Insert) => Dialog.create(payload)
}