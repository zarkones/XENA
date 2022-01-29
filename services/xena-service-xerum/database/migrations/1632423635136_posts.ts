import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Posts extends BaseSchema {
  protected tableName = 'posts'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()

      table.uuid('topic_id')
        .notNullable()

      table.uuid('author_id')
        .notNullable()

      table.string('title', 128)
        .notNullable()
      
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
