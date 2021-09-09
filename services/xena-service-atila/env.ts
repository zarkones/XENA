/*
|--------------------------------------------------------------------------
| Validating Environment Variables
|--------------------------------------------------------------------------
|
| In this file we define the rules for validating environment variables.
| By performing validation we ensure that your application is running in
| a stable environment with correct configuration values.
|
| This file is read automatically by the framework during the boot lifecycle
| and hence do not rename or move this file to a different location.
|
*/

import Env from '@ioc:Adonis/Core/Env'

export default Env.rules({
	HOST: Env.schema.string({ format: 'host' }),
	PORT: Env.schema.number(),
	APP_KEY: Env.schema.string(),
	APP_NAME: Env.schema.string(),
	NODE_ENV: Env.schema.enum(['development', 'production', 'testing'] as const),

	// Database config.
	DB_CONNECTION: Env.schema.string(),

	// Postgres database. Optional because of DATABASE_URL, so make sure one of those is set.
	PG_HOST: Env.schema.string.optional({ format: 'host' }),
  PG_PORT: Env.schema.number.optional(),
  PG_USER: Env.schema.string.optional(),
  PG_PASSWORD: Env.schema.string.optional(),
  PG_DB_NAME: Env.schema.string.optional(),

	// Heroku.
	DATABASE_URL: Env.schema.string.optional(),

	// Cross-Origin Resource Sharing Policy.
	// Comma separated list of hosts derived from the origin header.
	CORS_POLICY_ALLOWED_ORIGINS: Env.schema.string(),
})
