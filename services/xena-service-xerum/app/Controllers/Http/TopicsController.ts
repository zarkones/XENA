import * as Repo from 'App/Repos'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class TopicsController {
  public getMultiple = async ({ response }: HttpContextContract) => {
    const topics = await Repo.Topics.getMultiple()
    if (!topics.length)
      return response.noContent()

    return response.ok(topics)
  }
}
