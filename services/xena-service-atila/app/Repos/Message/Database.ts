import Message from 'App/Models/Message'

import { ClientId, MessageId } from '../Types'

type MessageStatus = 'SEEN' | 'SENT' | 'DELETED'

type Get = {
  id?: MessageId,
  status?: MessageStatus,
}

type GetMultiple = {
  clientId?: ClientId
  page?: number
  status?: MessageStatus
  replyTo?: MessageId
  noReplies?: boolean
}

type Insert = {
  id: MessageId
  from: string | null
  to: string | null
  subject: string
  content: string
  status: MessageStatus
  replyTo: MessageId | null
}

export default new class Database {
  public get = ({ id, status }: Get) => Message.query()
    .select('*')
    .whereNot('status', 'DELETED')
    .if(id, builder => builder.where('id', id!))
    .if(status, builder => builder.where('status', status!))
    .first()
    .then(client => client?.serialize())
  
  public getMultiple = ({ replyTo, clientId, page, status }: GetMultiple) => Message.query()
    .select('*')
    .whereNot('status', 'DELETED')
    .if(clientId, builder => builder.where('to', clientId!))
    .if(status, builder => builder.where('status', status!))
    .if(page, builder => builder.offset(page! * 10))
    .if(page, builder => builder.limit(10))
    .if(!replyTo, builder => builder.whereNull('reply_to'))
    .if(replyTo, builder => builder.where('reply_to', replyTo!))
    // .orWhereNull('to') Doesn't work with: .if(!replyTo, builder => builder.whereNull('reply_to'))
    .exec()
    .then(clients => clients.map(client => client.serialize()))
  
  public insert = (payload: Insert) => Message.create(payload).then(client => client.serialize())

  public ack = (id: MessageId) => Message.query()
    .where('id', id)
    .update({
      status: 'SEEN',
    })
    .exec()
  
  public delete = (id: MessageId) => Message.query()
    .where('id', id)
    .update({
      status: 'DELETED',
    })
    .first()
}
