import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class BuildProfiles extends BaseSchema {
  protected tableName = 'build_profiles'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()
        .primary()

      table.string('name', 255)
        .notNullable()

      table.string('description', 4096)
        .nullable()

      table.string('git_url', 2000)
        .notNullable()
      
      table.json('config')
        .notNullable()

      table.enum('status', [ 'ENABLED', 'DISABLED', 'DELETED' ])
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
