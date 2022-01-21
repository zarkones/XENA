import { DateTime } from 'luxon'
import { BaseModel, column, hasOne, HasOne } from '@ioc:Adonis/Lucid/Orm'
import System from 'App/Models/System'

export default class Client extends BaseModel {
  @column({ isPrimary: true })
  public id: string

  @column()
  public ip: string

  @column()
  public osId: string

  @hasOne(() => System, {
    foreignKey: 'id',
    localKey: 'osId',
  })
  public system: HasOne<typeof System>

  @column()
  public publicKey: string
  
  @column()
  public status: 'ALIVE' | 'DEAD' | 'BANNED'

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
