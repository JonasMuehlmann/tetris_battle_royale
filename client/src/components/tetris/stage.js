import Cell from './cell'
import { TETROMINOS } from './tetrominos'

const Stage = ({ stage, gameOver }) => (
  <div
    style={{
      ...gameOver && { filter: 'grayscale(1)' },
      display: 'grid',
      gridTemplateRows: `repeat(
        ${stage.length},
        calc(22vw / ${stage[0].length})
      )`,
      gridTemplateColumns: `repeat(
        ${stage[0].length},
        1fr
      )`,
      gap: '1px',
    }}>
    {
      stage?.map(row => row.map((cell, x) =>
        <Cell
          key={x}
          type={cell[0]}
          style={{
            width: `calc(22vw / ${stage[0].length})`,
            background: `${cell[0] === 0 ?
              `rgba(50, 50, 50, 0.2)` :
              `rgba(${TETROMINOS[cell[0]].color}, 0.8)`}`,
            borderWidth: `${cell[0] === 0 ? '0px' : '4px'}`,
            borderBottomColor: `rgba(${TETROMINOS[cell[0]].color}, 0.1)`,
            borderRightColor: `rgba(${TETROMINOS[cell[0]].color}, 1)`,
            borderTopColor: `rgba(${TETROMINOS[cell[0]].color}, 1)`,
            borderLeftColor: `rgba(${TETROMINOS[cell[0]].color}, 0.3)`,
          }}
        />)
      )
    }
  </div>
)

export default Stage