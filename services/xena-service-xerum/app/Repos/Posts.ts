import Post from 'App/Models/Post'

type GetMultiple = {
  id?: string
  topicId?: string
  limit: number
  offset?: number
}

type Get = {
  id: string
}

type Insert = {
  id: string
  topicId: string
  authorId: string
  title: string
  description: string | null
}

export default new class {
  public insert = (payload: Insert) => Post.create(payload)

  public get = ({ id }: Get) => Post.query()
    .where('id', id)
    .preload('author', builder => builder.select('name'))
    .first()

  public getMultiple = ({ id, topicId, limit, offset }: GetMultiple) => Post.query()
    .if(id, builder => builder.where('id', id!))
    .if(topicId, builder => builder.where('topic_id', topicId!))
    .if(limit, builder => builder.limit(limit!))
    .if(offset, builder => builder.offset(offset!))
    .preload('author', builder => builder.select('name'))
    .exec()
}