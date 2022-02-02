import Env from '@ioc:Adonis/Core/Env'

export const openAI = {
  apiKey: Env.get('OPEN_AI_KEY'),
  remote: Env.get('OPEN_AI_HOST'),
  temperature: 0.9,
  maxTokens: 200,
  topP: 1,
  frequencyPenalty: 0,
  presencePenalty: 0.6,
  stop: ['\n', ' Q:', ' A:']
}