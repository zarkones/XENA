type Query = {
  prefix?: string
  suffix?: string
}

const DefaultQueryOptions: Query = {
  prefix: '',
  suffix: ';',
}

export default class Postgres {
  public static readonly pgSleep = (options: Query = DefaultQueryOptions, sleepAmount: number) =>
    `${options.prefix}SELECT PG_SLEEP(${sleepAmount})${options.suffix}`

  public static readonly informationSchemaTables = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * FROM information_schema.tables${options.suffix}`

  public static readonly informationSchemaSchemata = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * from information_schema.schemata${options.suffix}`

  public static readonly pgCatalogNamespace = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * FROM pg_catalog.pg_namespace${options.suffix}`

  public static readonly pgCatalogTables = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * FROM pg_catalog.pg_tables${options.suffix}`
  
  /**
   * Return user's currently selected database.
   */
  public static readonly userSelectedDatabase = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT current_database()${options.suffix}`
  
  /**
   * Return all tables and their properties.
   */
  public static readonly allDatabases = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * FROM pg_catalog.pg_database${options.suffix}`
  
  /**
   * Given the table name, it will return columns and their properties.
   */
  public static readonly tableColumns = (tableName: string, options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * FROM information_schema.columns WHERE table_name = '${tableName}' ORDER BY ordinal_position${options.suffix}`
  
  /**
   * Returns all visible indexes.
   */
  public static readonly allIndexes = (options: Query = DefaultQueryOptions) =>
    `${options.prefix}SELECT * from pg_catalog.pg_indexes${options.suffix}`
}