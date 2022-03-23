import Stage from "./stage";
import { AnimatePresence, motion } from 'framer-motion'
import { useEffect, useRef, useState } from 'react'
import { useTimer } from '../../hooks/useTimer'
import { usePlayer } from '../../hooks/usePlayer'
import { useStage } from '../../hooks/useStage'
import { useStatus } from '../../hooks/useStatus'
import { checkCollision, createStage } from "./helpers";


const Tetris = ({
  onGameOver = () => { }
}) => {
  const wrapperRef = useRef(null)

  const [player, updatePlayerPos, resetPlayer, playerRotate] = usePlayer()
  const [stage, setStage, rowsCleared] = useStage(player, resetPlayer)
  const [score, setScore, rows, setRows, level, setLevel] = useStatus(rowsCleared)

  const [actionTimer, setActionTimer] = useState(null)
  const [timerCount, setTimerCount] = useState(3)
  const [dropTime, setDropTime] = useState(1000)
  const [gameStarted, setGameStarted] = useState(false)
  const [gameOver, setGameOver] = useState(false)

  const startGame = () => {
    setGameStarted(true)
    setStage(createStage())
    setDropTime(1000)
    resetPlayer()
    setScore(0)
    setLevel(0)
    setRows(0)
    setGameOver(false)
  }

  //#region MOVEMENTS

  const drop = () => {
    if (rows > (level + 1) * 10) {
      setLevel(prev => prev + 1)
      setDropTime(1000 / (level + 1) + 200)
    }

    if (!checkCollision(player, stage, { x: 0, y: 1 })) {
      updatePlayerPos({ x: 0, y: 1, collided: false })
    } else {
      if (player.pos.y < 1) {
        console.log('GAME OVER!')
        onGameOver()
        setGameOver(true)
        setDropTime(null)
      }
      updatePlayerPos({ x: 0, y: 0, collided: true })
    }
  }

  const dropBlock = () => {
    setDropTime(null)
    drop()
  }

  const moveBlock = dir => {
    if (!checkCollision(player, stage, { x: dir, y: 0 })) {
      updatePlayerPos({ x: dir, y: 0 });
    }
  }

  const softDrop = ({ keyCode }) => {
    if (!gameOver) {
      if (keyCode === 40) {
        setDropTime(1000 / (level + 1));
      }
    }
  }

  const move = ({ keyCode }) => {
    if (!gameOver) {
      if (keyCode === 37) {
        moveBlock(-1);
      } else if (keyCode === 39) {
        moveBlock(1);
      } else if (keyCode === 40) {
        dropBlock();
      } else if (keyCode === 38) {
        playerRotate(stage, 1);
      }
    }
  }

  //#endregion

  useTimer(() => {
    drop()
    if (timerCount > 1) setTimerCount(timerCount - 1)
    else if (!gameStarted) startGame()
  }, dropTime)

  useEffect(() => {
    wrapperRef?.current?.focus()
  }, [wrapperRef, actionTimer, timerCount])

  return (
    <div
      ref={wrapperRef}
      tabIndex='-1'
      onKeyDown={move}
      onKeyUp={softDrop}
      className="flex justify-center items-center gap-2 focus:outline-0 w-screen h-screen relative">
      <div className="absolute left-1/2 top-20 -translate-x-1/2">
        <p className="text-white font-semibold text-2xl">
          {score}
        </p>
      </div>
      <AnimatePresence exitBeforeEnter>
        {
          gameStarted ?
            (
              <motion.div
                key={gameStarted}
                initial={{ opacity: 0, scale: .25, rotateY: 180 }}
                animate={{ opacity: 1, scale: 1, rotateY: 0 }}
                transition={{ duration: 1.5, type: 'spring' }}
                className="flex gap-2">
                <Stage
                  stage={stage}
                  gameOver={gameOver}>
                </Stage>
              </motion.div>
            ) :
            (
              <motion.div
                key={gameStarted}
                initial={{ opacity: .25, y: 15, scale: 0.5 }}
                animate={{ opacity: 1, y: 0, scale: 1 }}
                exit={{ scale: 2.0, opacity: 0 }}
                transition={{ duration: .5, type: 'spring' }}
                className={`text-center
                  ${timerCount <= 1 && 'green-grad-text'}`}>
                <motion.p
                  className="text-8xl bangers">
                  {
                    timerCount
                  }
                </motion.p>
                <span className="text-4xl truncate w-40 bangers">
                  {
                    timerCount <= 1 &&
                    `Get Ready!`
                  }
                </span>
              </motion.div>
            )}
      </AnimatePresence>
    </div>
  )
}

export default Tetris