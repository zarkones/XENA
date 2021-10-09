import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Clients extends BaseSchema {
  protected tableName = 'clients'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id')
        .notNullable()
        .unique()
        .primary()

      table.string('public_key', 4096)
        .notNullable()

      // Client's state.
      table.enum('status', ['ALIVE', 'DEAD', 'BANNED'])
        .notNullable()

      table.timestamps(true)
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
