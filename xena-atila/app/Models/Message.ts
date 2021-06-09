import { DateTime } from 'luxon'
import { BaseModel, column } from '@ioc:Adonis/Lucid/Orm'

export default class Message extends BaseModel {
  @column({ isPrimary: true })
  public id: string
  
  @column()
  public from: string | null
  
  @column()
  public to: string | null
  
  @column()
  public subject: string
  
  @column()
  public content: string
  
  @column()
  public status: 'SEEN' | 'SENT'

  @column()
  public replyTo: string | null

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
