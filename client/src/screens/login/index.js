import axios from 'axios'
import { useState } from 'react'
import { AnimatePresence, motion } from 'framer-motion'
import { DialogType, useDialog } from '../../contexts/dialog-context'
import { Screen, useScreens } from '../../contexts/screen-context'
/*
 * COMPONENTS 
 */
import SignInForm from './sign_in_form'
import SignUpForm from './sign_up_form'
import GlowingText from '../../components/glowing_text/glowing_text'
/*
 * CONSTANTS
 */
const MODE = Object.freeze({
  SIGN_IN: 1,
  SIGN_UP: 2,
})

const LogInScreen = () => {
  /*
   * DEPENDENCIES
   */
  const {
    navigate,
  } = useScreens()

  const {
    showDialog,
    hideDialog,
  } = useDialog()
  /*
   * STATES
   */
  const [mode, setMode] = useState(MODE.SIGN_IN)

  // #region EVENTS

  const onSignIn = async (model, bypass = false) => {
    showDialog(DialogType.Authenticate)
    if (bypass) {
      setTimeout(() => { hideDialog(); navigate(Screen.Menu) }, 2500)
    }

    try {
      const res = await axios.post("/login", {
        username: model.username,
        password: model.password
      })
      console.log(res.data.user)

    } catch (error) {
      console.info(error.message)
    } finally {
      hideDialog()
    }
  }

  const onSignUp = async model => {
    try {
      showDialog(DialogType.Authenticate)
      const res = await axios.post(
        '/register',
        {
          username: model.username,
          password: model.password,
        }
      )
      console.log(res)
    } catch (error) {
      console.info(error)
    } finally {
      hideDialog()
    }
  }

  // #endregion

  return (
    <div
      className={`flex flex-col items-center justify-center z-20 
        w-full h-full text-white relative transition-all`}>
      <motion.div
        initial={{ opacity: 0, y: -window.innerHeight / 2 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 3, type: 'spring' }}
        className='flex flex-col items-center mb-16'>
        <GlowingText
          className={`text-9xl font-bold`}>
          Tetris
        </GlowingText>
        <h2
          className='text-5xl green-grad-text rounded pt-2 pb-5'>
          Battle Royale
        </h2>
        <p
          className='text-sm josefin'>
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