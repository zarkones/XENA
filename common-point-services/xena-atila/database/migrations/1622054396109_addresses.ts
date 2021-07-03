import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Addresses extends BaseSchema {
  protected tableName = 'addresses'

  public async up () {
    this.schema.createTable(this.tableName, (table) => {
      table.integer('x', 3)
        .nullable()
      table.integer('y', 3)
        .nullable()
      table.integer('z', 3)
        .nullable()
      table.integer('w', 3)
        .nullable()

      table.primary(['x','y','z','w'])

      table.enum('status', ['OK', 'BANNED', 'UNKNOWN'])
        .notNullable()

      table.timestamps(true)
    })
  }

  public async down () {
    this.schema.dropTable(this.tableName)
  }
}
