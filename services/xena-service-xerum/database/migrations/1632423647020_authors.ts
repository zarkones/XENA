import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Authors extends BaseSchema {
  protected tableName = 'authors'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()

      table.string('name', 128)
        .notNullable()
        .unique()

      table.string('public_key', 10024)
        .notNullable()
        .unique()
      
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
