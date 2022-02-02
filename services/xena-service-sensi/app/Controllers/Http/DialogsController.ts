import * as Validator from 'App/Validators'
import * as Service from 'App/Services'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class DialogsController {
  public getMultiple = async ({ response }: HttpContextContract) => {
    const maybeDialog = await Repo.Dialog.getMultiple()
    if (!maybeDialog)
      return response.noContent()

    const dialog = maybeDialog.map(d => Domain.Dialog.fromJSON(d))

    return response.ok(dialog)
  }

  public ask = async ({ request, response }: HttpContextContract) => {
    const { prompt } = await request.validate(Validator.Dialog.Ask)
    
    const answer = await Service.OpenAI.ask(
      'I am XENA, a highly intelligent question answering system. If you ask me a question that is rooted in truth, I will give you the answer. ' +
      'If you ask me a question that is nonsense, trickery, or has no clear answer, I will tell you that I am unable to respond.\n\n' +
      'Q: What is malware?\nA: Malware is any software intentionally designed to cause disruption to a computer, server, client, or computer network.\n\n' +
      'Q: What is XENA?\nA: I am XENA, managed remote administration system for botnet creation, development and its protection.\n\n' +
      `Q: ${prompt.trim()}\n`
    )

    const output = answer?.choices.map(c => c.text).join(' \n').replace('A: ', '')

    if (!answer || !output) {
      await Repo.Dialog.insert(Domain.Dialog.fromJSON({ id: uuidv4(), input: prompt, output: 'Unable to respond...' }))
      return response.internalServerError({ success: false, message: 'An error has occured.' })
    }

    await Repo.Dialog.insert(Domain.Dialog.fromJSON({ id: uuidv4(), input: prompt, output }))

    return response.ok(answer)
  }
}
