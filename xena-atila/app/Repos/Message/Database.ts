import Message from 'App/Models/Message'

import { ClientId, MessageId } from '../Types'

type MessageStatus = 'SEEN' | 'SENT'

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

class Database {
  public get = ({ id, status }: Get) => Message.query()
    .select('*')
    .if(id, builder => builder.where('id', id as MessageId))
    .if(status, builder => builder.where('status', status as MessageStatus))
    .first()
    .then(client => client?.serialize())
  
  public getMultiple = ({ replyTo, clientId, page, status }: GetMultiple) => Message.query()
    .select('*')
    .if(clientId, builder => builder.where('to', clientId as ClientId))
    .if(status, builder => builder.where('status', status as MessageStatus))
    .if(page, builder => builder.offset(page as number * 10))
    .if(page, builder => builder.limit(10))
    .if(!replyTo, builder => builder.whereNull('reply_to'))
    .if(replyTo, builder => builder.where('reply_to', replyTo as MessageId))
    // .orWhereNull('to') Doesn't work with: .if(!replyTo, builder => builder.whereNull('reply_to'))
    .exec()
    .then(clients => clients.map(client => client.serialize()))
  
  public insert = (payload: Insert) => Message.create(payload).then(client => client.serialize())

  public ack = (id: MessageId) => Message.query()
    .where('id', id)
    .update({
      status: 'SEEN'
    })
}

export default new Database()