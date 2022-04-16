import React, { useContext, useState } from 'react'
import { checkCollision, createStage } from '../components/tetris/helpers'
import { useTimer } from '../hooks/useTimer'
import { useKeybinds } from './keybinds-context'
import { usePlayer } from '../hooks/usePlayer'
import { useStage } from '../hooks/useStage'
import { useStatus } from '../hooks/useStatus'
import { useWS } from './websocket-context'
import { useAuth } from './auth-context'
import { useQueue } from './queue-context'

const TetrisContext = React.createContext()

export const TetrisProvider = ({ children }) => {
  const {
    user
  } = useAuth()

  const {
    currentMatch
  } = useQueue()

  const {
    sendJsonMessage,
    lastJsonMessage,
  } = useWS()

  const {
    keybinds
  } = useKeybinds()

  const [player, updatePlayerPos, resetPlayer, playerRotate] = usePlayer()
  const [stage, setStage, rowsCleared] = useStage(player, resetPlayer)
  const [score, setScore, rows, setRows, level, setLevel] = useStatus(rowsCleared)

  const [timerCount, setTimerCount] = useState(3)
  const [dropTime, setDropTime] = useState(1000)
  const [gameStarted, setGameStarted] = useState(false)
  const [gameOver, setGameOver] = useState(false)

  const onKeyUp = ({ keyCode }) => {
    if (!gameOver) {
      if (keyCode === keybinds.drop.key) {
        setDropTime(1000 / (level + 1))
      }
    }
  }

  async function _sendRequest({ type, direction }) {
    if (
      user === undefined || user === null ||
      currentMatch === null
    ) return

    let message = {
      type,
      userID: user.id,
      matchID: currentMatch.id
    }

    if (type === 'MoveBlock' || type === 'RotateBlock') {
      message = { ...message, direction }
    }

    sendJsonMessage(message)
  }

  /*
   * Control function to be exposed
   */
  const onKeyDown = ({ keyCode }) => {
    if (!gameOver) {
      if (keyCode === keybinds.left.key) {
        // _sendRequest({ type: 'MoveBlock', direction: 'left' })
        actions.moveBlock(-1)
      } else if (keyCode === keybinds.right.key) {
        // _sendRequest({ type: 'MoveBlock', direction: 'right' })
        actions.moveBlock(1)
      } else if (keyCode === keybinds.drop.key) {
        // _sendRequest({ type: 'MoveBlock', direction: 'down' })
        actions.softDrop()
      } else if (keyCode === keybinds.rotate.key) {
        // _sendRequest({ type: 'RotateBlock', direction: 'right' })
        playerRotate(stage, 1)
      }
    }
  }

  const actions =
  {
    /*
     * Initializes/Resets game
     */
    start: () => {
      setGameStarted(true)
      setStage(createStage())
      setDropTime(1000)
      resetPlayer()
      setScore(0)
      setLevel(0)
      setRows(0)
      setGameOver(false)
    },
    /*
     * Triggers on every drop time
     */
    drop: () => {
      if (rows > (level + 1) * 10) {
        setLevel(prev => prev + 1)
        setDropTime(1000 / (level + 1) + 200)
      }

      if (!checkCollision(player, stage, { x: 0, y: 1 })) {
        updatePlayerPos({ x: 0, y: 1, collided: false })
      } else {
        if (player.pos.y < 1) {
          console.log('GAME OVER!')
          setGameOver(true)
          setDropTime(null)
        }
        updatePlayerPos({ x: 0, y: 0, collided: true })
      }
    },
    /*
     * Instantly triggers drop
     */
    softDrop: () => {
      setDropTime(null)
      actions.drop()
    },
    /*
     * Horizontal movements
     */
    moveBlock: dir => {
      if (!checkCollision(player, stage, { x: dir, y: 0 })) {
        updatePlayerPos({ x: dir, y: 0 });
      }
    },
  }

  useTimer(() => {
    if (gameStarted) actions.drop()
    if (!gameStarted && timerCount <= 1) actions.start()
    if (timerCount > 1) setTimerCount(timerCount - 1)
  }, dropTime)

  return (
    <TetrisContext.Provider
      value={{
        onKeyUp,
        onKeyDown,
        score,
        gameStarted,
        gameOver,
        stage,
        timerCount
      }}>
      {children}
    </TetrisContext.Provider>
  )
}

export const withTetris = Component => (props) => (
  <TetrisProvider>
    <Component {...props} />
  </TetrisProvider>
)

export const useTetris = () => useContext(TetrisContext)