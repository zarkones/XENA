import { DateTime } from 'luxon'
import { BaseModel, column } from '@ioc:Adonis/Lucid/Orm'

export default class Address extends BaseModel {
  @column({ isPrimary: true })
  public x: number

  @column({ isPrimary: true })
  public y: number

  @column({ isPrimary: true })
  public z: number

  @column({ isPrimary: true })
  public w: number

  @column()
  public status: 'OK' | 'BANNED' | 'UNKNOWN'

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
