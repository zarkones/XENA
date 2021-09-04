import * as Validator from 'App/Validators'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { execSync } from 'child_process'
import { quote } from 'shell-quote'

export default class MappersController {
  public nmap = async ({ request, response }: HttpContextContract) => {
    const { address } = await request.validate(Validator.Mappers.Nmap)

    const command = quote(['nmap', address])

    const result = execSync(command).toString('utf-8')
    
    return response.ok(result)
  }
}
