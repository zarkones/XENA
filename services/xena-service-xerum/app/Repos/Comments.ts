import Comment from 'App/Models/Comment'

type GetMultiple = {
  postId: string
  limit: number
  offset?: number
}

type Insert = {
  id: string
  authorId: string
  postId: string
  description: string
}

export default new class {
  public insert = (payload: Insert) => Comment.create(payload)

  public getMultiple = ({ postId, limit, offset }: GetMultiple) => Comment.query()
    .if(postId, builder => builder.where('post_id', postId!))
    .if(limit, builder => builder.limit(limit!))
    .if(offset, builder => builder.offset(offset!))
    .exec()
}