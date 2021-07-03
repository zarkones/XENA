import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Builds extends BaseSchema {
  protected tableName = 'builds'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()
        .primary()

      table.uuid('build_profile_id')
        .notNullable()
        .references('id')
        .inTable('build_profiles')
      
      table.binary('data')
        .notNullable()

      /**
       * Uses timestampz for PostgreSQL and DATETIME2 for MSSQL
       */
      table.timestamp('created_at', { useTz: true })
      table.timestamp('updated_at', { useTz: true })
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
