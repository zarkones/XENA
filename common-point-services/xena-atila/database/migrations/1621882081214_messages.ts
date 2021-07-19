import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Messages extends BaseSchema {
  protected tableName = 'messages'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()
        .primary()
      
      // ID of the originating client.
      table.string('from', 255)
        .nullable()

      // ID od the recipient client.
      table.string('to', 255)
        .nullable()

      // Routing subject.
      table.string('subject', 255)
        .notNullable()

      // Base encoded.
      table.string('content', 65536)
        .notNullable()

      table.uuid('reply_to')
        .nullable()
        .references('id')
        .inTable('messages')

      // Client's state.
      table.enum('status', ['SEEN', 'SENT'])
        .notNullable()

      table.index(['to', 'status'])

      table.timestamps(true)
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
