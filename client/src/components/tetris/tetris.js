import Stage from "./stage";
import { AnimatePresence, motion } from 'framer-motion'
import { useEffect, useRef } from 'react'
import { useTetris } from "../../contexts/tetris-context";
import ScoreBoard from "./score_board"
const Tetris = ({
  onGameOver = () => { }
}) => {
  const {
    onKeyDown,
    onKeyUp,
    score,
    gameStarted,
    gameOver,
    stage,
    timerCount
  } = useTetris();

  const wrapperRef = useRef(null)

  useEffect(() => {
    wrapperRef?.current?.focus()

    if (gameOver) onGameOver()
  }, [wrapperRef, timerCount, gameOver])

  return (
    <div
      ref={wrapperRef}
      tabIndex='-1'
      onKeyDown={onKeyDown}
      onKeyUp={onKeyUp}
      className="flex justify-center items-center gap-2 focus:outline-0 w-screen h-screen relative">
      <ScoreBoard className="absolute inset-x-0 top-0 " score={score} />
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