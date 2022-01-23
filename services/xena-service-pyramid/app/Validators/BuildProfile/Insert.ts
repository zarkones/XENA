import { schema, rules } from '@ioc:Adonis/Core/Validator'
import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { buildTemplates } from 'App/Domains/BuildProfile'

export default class {
  constructor (protected ctx: HttpContextContract) {
  }

	/*
	 * Define schema to validate the "shape", "type", "formatting" and "integrity" of data.
	 *
	 * For example:
	 * 1. The username must be of data type string. But then also, it should
	 *    not contain special characters or numbers.
	 *    ```
	 *     schema.string({}, [ rules.alpha() ])
	 *    ```
	 *
	 * 2. The email must be of data type string, formatted as a valid
	 *    email. But also, not used by any other user.
	 *    ```
	 *     schema.string({}, [
	 *       rules.email(),
	 *       rules.unique({ table: 'users', column: 'email' }),
	 *     ])
	 *    ```
	 */
  public schema = schema.create({
		name: schema.string({ escape: true }, [ rules.maxLength(255) ]),
		description: schema.string.optional({ escape: true }, [ rules.maxLength(4096) ]),
		gitUrl: schema.string({}, [ rules.maxLength(2000) ]),
		config: schema.object().members({
			template: schema.enum(buildTemplates),
			atilaHost: schema.string.optional(),
			trustedPublicKey: schema.string.optional(),
			dgaSeed: schema.string.optional(),
			dgaAfterDays: schema.string.optional(),
			maxLoopWait: schema.string.optional(),
			minLoopWait: schema.string.optional(),
			gettrProfileName: schema.string.optional(),
		}),
		status: schema.enum([ 'ENABLED', 'DISABLED', 'DELETED' ] as const),
  })

	/**
	 * Custom messages for validation failures. You can make use of dot notation `(.)`
	 * for targeting nested fields and array expressions `(*)` for targeting all
	 * children of an array. For example:
	 *
	 * {
	 *   'profile.username.required': 'Username is required',
	 *   'scores.*.number': 'Define scores as valid numbers'
	 * }
	 *
	 */
  public messages = {}
}
