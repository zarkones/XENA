import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Services extends BaseSchema {
  protected tableName = 'services'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .primary()
        .unique()
        .notNullable()
      
      table.string('address', 15)
        .notNullable()
      
      table.integer('port')
        .notNullable()
      
      table.json('details')
        .nullable()

      table.timestamp('created_at')
      table.timestamp('updated_at')
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
