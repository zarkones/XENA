import { validNumber, validString } from './Validators'

export default class System {
  constructor (
    public readonly id: string,
    public readonly name: string,
    public readonly arch: string | null,
    public readonly count: number,
  ) {
    this.id = validString(id, 'BAD_SYSTEM_ID', 'NON_EMPTY')
    this.name = validString(name, 'BAD_SYSTEM_NAME', 'NON_EMPTY')
    this.arch = arch ? validString(arch, 'BAD_SYSTEM_ARCH', 'NON_EMPTY') : null
    this.count = validNumber(count, 'BAD_SYSTEM_COUNT')
  }

  public static fromJSON = (json) => new System(
    json.id,
    json.name,
    json.arch,
    json.count,
  )
}