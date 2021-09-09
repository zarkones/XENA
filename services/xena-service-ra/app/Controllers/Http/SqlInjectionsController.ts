import * as Validator from 'App/Validators'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class SqlInjectionsController {
  public scan = async ({ request, response }: HttpContextContract) => {
    const { url } = await request.validate(Validator.SqlInjections.Scan)

    const tries = []

    return response.ok(parsedUrl)
  }
}
