export const validString = (value: string, message: string, state?: 'NON_EMPTY') => {
  switch (state) {
    case 'NON_EMPTY':
      if (typeof value !== 'string')
        throw Error(`${message}:${value}`)

      return value

    default:
      throw Error(`${message}:${value}`)
  }
}

export const validNumber = (value: number, message: string, unsigned?: boolean) => {
  if (typeof value !== 'number')
    throw Error(`${message}:${value}`)
    
  if (unsigned && value < 0)
    throw Error(`${message}:${value}`)

  return value
}

export const validEnum = <T> (value: T, list: string[], message: string) => {
  if (typeof value !== 'string')
    throw Error(`${message}:${value}`)

  if (!list.length)
    throw Error(`${message}:${value}`)

  if (!list.includes(value))
    throw Error(`${message}:${value}`)

  return value
}