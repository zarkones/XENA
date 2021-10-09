import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class CommentsController {
  public insert = async ({ request, response }: HttpContextContract) => {
    const { authorId, postId, description } = await request.validate(Validator.Comments.Insert)

    const post = await Repo.Posts.get({ id: postId })
    if (!post)
      return response.notFound({ success: false, message: 'Post not found.' })

    const comment = await Repo.Comments.insert({ id: uuidv4(), authorId, postId, description })

    return response.ok(comment)
  }

  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { postId, limit, offset } = await request.validate(Validator.Comments.GetMultiple)

    const comments = await Repo.Comments.getMultiple({ postId, limit, offset })
    if (!comments.length)
      return response.noContent()

    return response.ok(comments)
  }
}
