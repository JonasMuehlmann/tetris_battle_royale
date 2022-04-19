// 4X4
export const TETROMINOS = {
  0: { shape: [[0]], color: '0, 0, 0' },
  I: {
    shape: [
      [0, 0, 0, 0],
      ['I', 'I', 'I', 'I'],
      [0, 0, 0, 0],
      [0, 0, 0, 0],
    ],
    color: '80, 227, 230'
  },
  J: {
    shape: [
      [0, 0, 0, 0],
      [0, 0, 0, 'J'],
      ['J', 'J', 'J', 'J'],
      [0, 0, 0, 0],
    ],
    color: '36, 95, 223'
  },
  L: {
    shape: [
      [0, 0, 0, 0],
      ['L', 0, 0, 0],
      ['L', 'L', 'L', 'L'],
      [0, 0, 0, 0],
    ],
    type: 'L',
    color: '223, 173, 36'
  },
  O: {
    shape: [
      [0, 0, 0, 0],
      [0, 'O', 'O', 0],
      [0, 'O', 'O', 0],
      [0, 0, 0, 0],
    ],
    color: '223, 173, 36'
  },
  S: {
    shape: [
      [0, 'S', 'S', 0],
      ['S', 'S', 0, 0],
      [0, 0, 0, 0],
      [0, 0, 0, 0],
    ],
    color: '48, 211, 56'
  },
  T: {
    shape: [
      [0, 0, 0, 0],
      [0, 'T', 0, 0],
      ['T', 'T', 'T', 0],
      [0, 0, 0, 0],
    ],
    color: '132, 61, 198'
  },
  Z: {
    shape: [
      [0, 0, 0, 0],
      ['Z', 'Z', 0, 0],
      [0, 'Z', 'Z', 0],
      [0, 0, 0, 0],
    ],
    color: '227, 78, 78'
  },
}

// #TODO NOT SURE IF NEEDED
export const convertToTetromino = data => {
  [...data].forEach(row => {
  })
}

export const randomTetromino = () => {
  const types = 'IJLOSTZ';
  const randTetro = types[Math.floor(Math.random() * types.length)]

  return TETROMINOS[randTetro]
}