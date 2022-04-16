import { useState } from 'react'
import { AnimatePresence, motion } from 'framer-motion'
/*
 * COMPONENTS 
 */
import SignInForm from './sign_in_form'
import SignUpForm from './sign_up_form'
import GlowingText from '../../components/glowing_text/glowing_text'
import { AuthContext } from '../../contexts/auth-context'
/*
 * CONSTANTS
 */
const MODE = Object.freeze({
  SIGN_IN: 1,
  SIGN_UP: 2,
})

const LogInScreen = () => {
  /*
   * STATES
   */
  const [mode, setMode] = useState(MODE.SIGN_IN)

  // #endregion

  return (
    <AuthContext.Consumer>
      {
        ({ signIn, signUp }) => (
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
                    onSubmit={signIn}
                    onSignUp={() => setMode(MODE.SIGN_UP)}
                  />
                ) : (
                  <SignUpForm
                    key={mode}
                    onSubmit={signUp}
                    onSignIn={() => setMode(MODE.SIGN_IN)}
                  />
                )
              }
            </AnimatePresence>
          </div>
        )
      }
    </AuthContext.Consumer>
  )
}

export default LogInScreen