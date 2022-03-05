import { AnimatePresence, motion } from 'framer-motion'
import { useContext, useState } from 'react'
import { DialogType, useDialog } from '../../contexts/dialog-context'
import { Screen, ScreenContext } from '../../contexts/screen-context'
import SignInForm from './sign_in_form'
import SignUpForm from './sign_up_form'
import axios from 'axios'

const MODE = {
  SIGN_IN: 1,
  SIGN_UP: 2,
}

const LogInScreen = () => {
  const { navigate } = useContext(ScreenContext)
  const [mode, setMode] = useState(MODE.SIGN_IN)
  const {
    showDialog,
    hideDialog,
  } = useDialog()

  const onSignIn = async model => {
    try {
      showDialog(DialogType.Authenticate)
      /**
       * IN ORDER TO TEST THIS API
       * YOU MUST START THE GATEWAY FIRST (SERVER)
       * OTHERWISE IT WILL AUTOMATICALLY NAVIGATE TO LOBBY
       * 
       * const response = await axios.post(`/user/login`, {
       *    username: model.username,
       *    password: model.password,
       * })
       * hideDialog()
       * 
       */

      /**
       * FOR UI TEST PURPOSE
       * COMMENT OUT IF SERVER IS ON
       */
      setTimeout(() => { hideDialog(); navigate(Screen.Menu) }, 2500)
    } catch (error) {
      console.info(error)
    }
  }

  const onSignUp = async model => {
    try {
      showDialog(DialogType.Authenticate)
      /**
       * IN ORDER TO TEST THIS API
       * YOU MUST START THE GATEWAY FIRST (SERVER)
       * OTHERWISE IT WILL AUTOMATICALLY NAVIGATE TO LOBBY
       * 
       * const result = await axios.post(`/user`, {
       *    username: model.username,
       *    password: model.password,
       * })
       * hideDialog()
       */

      /**
       * FOR UI TEST PURPOSE
       * COMMENT OUT IF SERVER IS ON
       */
      setTimeout(() => { hideDialog() }, 2500)
    } catch (error) {
      console.info(error)
    }
  }

  return (
    <div
      className={`flex flex-col items-center justify-center z-20 
        w-full h-full text-white relative transition-all`}>
      <motion.div
        initial={{ opacity: 0, y: -window.innerHeight / 2 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 3, type: 'spring' }}
        className='flex flex-col items-center mb-16'>
        <h2 className='text-9xl font-bold tetris-text tetris-shadow'>
          Tetris
        </h2>
        <h2 className='text-5xl green-grad-text rounded pt-2 pb-5'>
          Battle Royale
        </h2>
        <p className='text-sm josefin'>
          Massively Multiplayer Classic Tetris
        </p>
      </motion.div>
      <AnimatePresence
        exitBeforeEnter>
        {
          mode === MODE.SIGN_IN ? (
            <SignInForm
              key={mode}
              onSubmit={onSignIn}
              onSignUp={() => setMode(MODE.SIGN_UP)}
            />
          ) : (
            <SignUpForm
              key={mode}
              onSubmit={onSignUp}
              onSignIn={() => setMode(MODE.SIGN_IN)}
            />
          )
        }
      </AnimatePresence>
    </div>
  )
}

export default LogInScreen