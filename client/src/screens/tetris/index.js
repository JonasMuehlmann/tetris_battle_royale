import { useEffect, useState } from "react"
import Tetris from "../../components/tetris/tetris"
import { useScreens, Screen } from '../../contexts/screen-context'
import { withTetris } from "../../contexts/tetris-context"

const TetrisScreen = () => {
  const {
    navigate
  } = useScreens()

  return (
    <div className="flex flex-col justify-center items-center z-50 w-full h-full text-white relative">
      <button
        onClick={() => navigate(Screen.Menu)}
        className="absolute left-28 top-20 text-4xl opacity-20 transition-all font-semibold hover:opacity-100">
        BACK
      </button>
      <Tetris
        onGameOver={() => setTimeout(() => navigate(Screen.Menu), 1000)}
      />
    </div>
  )
}

export default withTetris(TetrisScreen)