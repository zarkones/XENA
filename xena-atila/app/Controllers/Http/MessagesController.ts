import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class MessagesController {
  public get = async ({ request, response }: HttpContextContract) => {
    const { id, status } = await request.validate(Validator.Message.Get)

    const maybeMessage = await Repo.Message.get({ id, status })
    if (!maybeMessage)
      return response.noContent()
    
    const message = Domain.Message.fromJSON(maybeMessage)

    return response.ok(message)
  }

  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { page, status, clientId, withReplies } = await request.validate(Validator.Message.GetMultiple)

    const maybeMessages = await Repo.Message.getMultiple({ page, status, clientId, noReplies: !withReplies })
    if (!maybeMessages.length)
      return response.noContent()

    const messages = !withReplies
      ? maybeMessages.map(maybeMessage => Domain.Message.fromJSON(maybeMessage))
      : await Promise.all(maybeMessages.map(async maybeMessage => {
          const message = Domain.Message.fromJSON(maybeMessage)
          const replies = await Repo.Message.getMultiple({ replyTo: message.id })
            .then(messages => messages.map(m => Domain.Message.fromJSON(m)))
          return { ...message, replies }
        }))

    return response.ok(messages)
  }

  public insert = async ({ request, response }: HttpContextContract) => {
    const { from, to, toMultiple, subject, content, replyTo } = await request.validate(Validator.Message.Insert)

    const message = toMultiple
      // Insert a message for each of the recipients.
      ? await Promise.all(toMultiple.map(to => Repo.Message.insert(Domain.Message.fromJSON({ from, to, subject, content, status: 'SENT', replyTo }))
        .then(message => Domain.Message.fromJSON(message))))
      // Insert a message for a single recipient.
      : await Repo.Message.insert(Domain.Message.fromJSON({ from, to, subject, content, status: 'SENT', replyTo }))
        .then(message => Domain.Message.fromJSON(message))

    return response.ok(message)
  }

  public ack = async ({ request, response }: HttpContextContract) => {
    const { id } = await request.validate(Validator.Message.Update)

    const maybeMesssage = await Repo.Message.get({ id })
    if (!maybeMesssage)
      return response.notFound({ success: false, message: 'Message not found.' })

    const message = Domain.Message.fromJSON(maybeMesssage)

    const ack = await Repo.Message.ack(message.id)

    return ack
  }

  public delete = async ({}: HttpContextContract) => {
    // todo
  }
}
