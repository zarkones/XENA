import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Dialogs extends BaseSchema {
  protected tableName = 'dialogs'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .primary()
        .notNullable()
      
      table.string('input', 4096)
        .notNullable()

      table.string('output', 40000)
        .nullable()

      /**
       * Uses timestamptz for PostgreSQL and DATETIME2 for MSSQL
       */
      table.timestamp('created_at', { useTz: true })
      table.timestamp('updated_at', { useTz: true })
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
