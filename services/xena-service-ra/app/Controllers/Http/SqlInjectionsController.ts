import * as Validator from 'App/Validators'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { execSync } from 'child_process'
import { quote } from 'shell-quote'

export default class SqlInjectionsController {
  public scan = async ({ request, response }: HttpContextContract) => {
    const { url } = await request.validate(Validator.SqlInjections.Scan)

    const command = quote(['sqlmap', '--batch', '-u', url])

    const result = execSync(command).toString('utf-8')
      .split('caused by this program')[1]
      .split('\n')
      .filter(line => line)
    
    return response.ok(result)
  }
}
