import { BaseModel } from '@ioc:Adonis/Lucid/Orm'

BaseModel.namingStrategy.serializedName = (_model, attributeName: string): string => {
  return attributeName
}