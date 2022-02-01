import * as Validator from 'App/Validators'
import * as Service from 'App/Services'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class DialogsController {
  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const maybeDialog = await Repo.Dialog.getMultiple()
    if (!maybeDialog)
      return response.noContent()

    const dialog = maybeDialog.map(d => Domain.Dialog.fromJSON(d))

    return response.ok(dialog)
  }

  public ask = async ({ request, response }: HttpContextContract) => {
    const { prompt } = await request.validate(Validator.Dialog.Ask)
    
    const answer = await Service.OpenAI.ask(prompt)
    if (!answer) {
      await Repo.Dialog.insert(Domain.Dialog.fromJSON({ id: uuidv4(), input: prompt, output: null }))
      return response.internalServerError({ success: false, message: 'An error has occured.' })
    }

    // await Repo.Dialog.insert(Domain.Dialog.fromJSON({ id: uuidv4(), input: prompt, output: answer. }))

    return response.ok(answer)
  }
}
