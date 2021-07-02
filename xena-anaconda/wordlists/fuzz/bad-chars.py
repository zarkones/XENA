def Bad_Chars() -> set[str]:
  return {
    '00',
    '%00',
    'x00',
    '\00',
    '\0',
    'u\\00',
    '\'',
    '"',
    '`'
  }