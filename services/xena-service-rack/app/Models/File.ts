import { DateTime } from 'luxon'
import { BaseModel, column } from '@ioc:Adonis/Lucid/Orm'

export default class File extends BaseModel {
  @column({ isPrimary: true })
  public id: string
  
  @column()
  public originId: string
  
  @column()
  public originPath: string
  
  @column()
  public data: string

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
