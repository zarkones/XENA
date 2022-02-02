import Config from '@ioc:Adonis/Core/Config'

import axios from 'axios'

export type OpenAIConfig = {
  apiKey: string // Keep protected, please.
  remote: string
  temperature: number
  maxTokens: number
  topP: number
  frequencyPenalty: number
  presencePenalty: number
  stop: string[]
}

type OpenAIChoice = {
  text: string
  index: number
  logprobs: any
  finishReason: string
}

type OpenAIResponse = {
  id: string
  object: string
  created: number
  model: string
  choices: OpenAIChoice[]
}

class OpenAI {
  // Handle with care.
  protected readonly apiKey: string

  public readonly remote: string
  public readonly temperature: number
  public readonly maxTokens: number
  public readonly topP: number
  public readonly frequencyPenalty: number
  public readonly presencePenalty: number
  public readonly stop: string[]

  constructor (config: OpenAIConfig) {
    this.apiKey = config.apiKey

    this.remote = config.remote
    this.temperature = config.temperature
    this.maxTokens = config.maxTokens
    this.topP = config.topP
    this.frequencyPenalty = config.frequencyPenalty
    this.presencePenalty = config.presencePenalty
    this.stop = config.stop
  } 

  public ask = async (prompt: string) => {
    const response = await axios.post(`${this.remote}/v1/engines/davinci/completions`,
      {
        prompt,
        temperature: this.temperature,
        max_tokens: this.maxTokens,
        top_p: this.topP,
        frequency_penalty: this.frequencyPenalty,
        presence_penalty: this.presencePenalty,
        stop: this.stop,
      },
      {
        headers: {
          Host: 'api.openai.com',
          Authorization: `Bearer ${this.apiKey}`,
          'Content-Type': 'application/json',
        },
      },
    )
      .catch(error => console.warn(error))
      .then(response => {
        if (response) {
          return {
            id: response.data.id,
            object: response.data.object,
            created: response.data.created,
            model: response.data.model,
            choices: response.data.choices.map(choice => ({
              text: choice.text,
              index: choice.index,
              logprobs: choice.logprobs,
              finishReason: choice.finish_reason,
            }) as OpenAIChoice),
          } as OpenAIResponse
        }
      })

    return response
  }
}

export default new OpenAI(Config.get('api.openAI'))