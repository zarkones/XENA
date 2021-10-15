import { validString, validEnum } from './Validators'
import { v4 as uuidv4 } from 'uuid'

type MessageStatus = 'SEEN' | 'SENT'

export default class Message {
  constructor (
    public readonly id: string,
    public readonly from: string | null,
    public readonly to: string | null,
    public readonly subject: string,
    public readonly content: string,
    public readonly status: MessageStatus,
    public readonly replyTo: string | null,
  ) {
    this.id = validString(id ?? uuidv4(), 'BAD_MESSAGE_ID', 'NON_EMPTY')
    this.from = from ? validString(from, 'BAD_MESSAGE_FROM', 'NON_EMPTY') : null
    this.to = to ? validString(to, 'BAD_MESSAGE_TO', 'NON_EMPTY') : null
    this.subject = validString(subject, 'BAD_MESSAGE_SUBJECT', 'NON_EMPTY')
    this.content = validString(content, 'BAD_MESSAGE_CONTENT', 'NON_EMPTY')
    this.status = validEnum(status, ['SEEN', 'SENT'], 'BAD_MESSAGE_STATUS')
    this.replyTo = replyTo ? validString(replyTo, 'BAD_MESSAGE_REPLY_TO_ID', 'NON_EMPTY') : null
  }

  public static fromJSON = (json) => new Message(
      json.id,
      json.from,
      json.to,
      json.subject,
      json.content,
      json.status,
      json.replyTo,
    )
}