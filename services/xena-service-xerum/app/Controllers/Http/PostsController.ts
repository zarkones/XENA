import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class PostsController {
  public insert = async ({ request, response }: HttpContextContract) => {
    const { authorId, topicId, title, description } = await request.validate(Validator.Posts.Insert)

    const topic = await Repo.Topics.get({ id: topicId })
    if (!topic)
      return response.notFound({ success: false, message: 'Topic not found.' })

    const post = await Repo.Posts.insert({
      id: uuidv4(),
      authorId,
      topicId,
      title,
      description: description ?? null,
    })

    return response.ok(post)
  }

  public get = async ({ request, response }: HttpContextContract) => {
    const { id } = await request.validate(Validator.Posts.Get)
    
    const post = await Repo.Posts.get({ id })
    if (!post)
      return response.notFound({ success: false, message: 'Post not found.' })

    return response.ok(post)
  }

  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { topicId, limit, offset } = await request.validate(Validator.Posts.GetMultiple)

    const posts = await Repo.Posts.getMultiple({ id: topicId, limit, offset })
    if (!posts.length)
      return response.noContent()

    return response.ok(posts)
  }
}
