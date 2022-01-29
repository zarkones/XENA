from domains.validators import Validators

class Message:
  def __init__(self, id: str, from_id: str, to: str, subject: str, content: str, status: str, reply_to: str,):
    self.id = Validators.valid_string(id, 'BAD_MESSAGE_ID')
    self.from_id = from_id
    # self.from_id = Validators.valid_string(from_id, 'BAD_MESSAGE_FROM_ID')
    self.to = Validators.valid_string(to, 'BAD_MESSAGE_TO')
    self.subject = Validators.valid_string(subject, 'BAD_MESSAGE_SUBJECT')
    self.content = Validators.valid_string(content, 'BAD_MESSAGE_CONTENT')
    self.status = Validators.valid_string(status, 'BAD_MESSAGE_STATUS')
    self.reply_to = reply_to
    # self.reply_to = Validators.valid_string(reply_to, 'BAD_MESSAGE_REPLY_TO')

  @staticmethod
  def from_json(json):
    return Message(
      id = json['id'],
      from_id = json['from'],
      to = json['to'],
      subject = json['subject'],
      content = json['content'],
      status = json['status'],
      reply_to = json['replyTo'],
    )

  def serialize(self):
    return {
      'id': self.id,
      'from': self.from_id,
      'to': self.to,
      'subject': self.subject,
      'content': self.content,
      'status': self.status,
      'replyTo': self.reply_to,
    }