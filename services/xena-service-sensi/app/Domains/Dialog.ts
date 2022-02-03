import { validString } from './Validators'

export default class Dialog {
  constructor (
    public readonly id: string,
    public readonly input: string,
    public readonly output: string | null,
  ) {
    id = validString(id, 'BAD_DIALOG_ID', 'NON_EMPTY')
    input = validString(input, 'BAD_DIALOG_INPUT', 'NON_EMPTY')
    output = output ? validString(output, 'BAD_DIALOG_OUTPUT', 'NON_EMPTY') : null
  }

  public static fromJSON = (json: any) => new Dialog(
    json.id,
    json.input,
    json.output,
  )
}