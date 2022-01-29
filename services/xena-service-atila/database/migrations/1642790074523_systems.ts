import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Systems extends BaseSchema {
  protected tableName = 'systems'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()
        .primary()
      
      table.string('name', 64)
        .notNullable()
      
      table.string('arch', 32)
        .nullable()
      
      table.integer('count')
        .defaultTo(0)
        .notNullable()

      table.timestamps(true)
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
