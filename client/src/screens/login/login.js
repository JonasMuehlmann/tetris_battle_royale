import { useContext, useState } from 'react'
import { Screen, ScreenContext } from '../../contexts/screen-context'
import SignInForm from './sign_in_form'
import SignUpForm from './sign_up_form'

const MODE = {
  SIGN_IN: 1,
  SIGN_UP: 2,
}

const styles = {
  container: 'w-full h-full flex flex-col items-center justify-center z-20 text-white relative',
  stack: 'flex flex-col gap-2',
}

const LogIn = () => {
  const { navigate } = useContext(ScreenContext)
  const [mode, setMode] = useState(MODE.SIGN_IN)

  const onSignIn = async model => {
    /* TODO: LOGIN API WITH REQUEST CLASS */
    try {
      // const result = await fetch(`isLogin/jaykim`)
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className={styles.container}>
      <p className='absolute top-10 right-10'>
        IN-DEVELOPMENT
      </p>
      <div className='flex flex-col items-center mb-16'>
        <h2 className='text-9xl font-bold tetris-text tetris-shadow'>
          Tetris
        </h2>
        <h2 className='text-5xl green-grad-text rounded pt-2 pb-5'>
          Battle Royale
        </h2>
        <p className='text-sm josefin'>
          Massively Multiplayer Classic Tetris
        </p>
      </div>
      {
        mode === MODE.SIGN_IN ? (
          <SignInForm
            onSubmit={onSignIn}
            onSignUp={() => setMode(MODE.SIGN_UP)}
          />
        ) : (
          <SignUpForm
            onSubmit={onSignIn}
            onSignIn={() => setMode(MODE.SIGN_IN)}
          />
        )
      }
    </div>
  )
}

export default LogIn