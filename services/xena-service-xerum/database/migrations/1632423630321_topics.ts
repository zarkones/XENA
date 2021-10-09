import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Topics extends BaseSchema {
  protected tableName = 'topics'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()
      
      table.string('title', 128)
        .notNullable()
        .unique()
      
      table.string('description', 4096)
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
