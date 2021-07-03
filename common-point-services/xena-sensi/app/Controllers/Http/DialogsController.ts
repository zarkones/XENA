import * as Validator from 'App/Validators'
import * as Service from 'App/Services'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class DialogsController {
  public ask = async ({ request, response }: HttpContextContract) => {
    const { prompt } = await request.validate(Validator.Dialog.Ask)
    
    const answer = await Service.OpenAI.ask(prompt)
    if (!answer)
      return response.internalServerError({ success: false, message: 'An error has occured.' })

    return response.ok(answer)
  }
}
