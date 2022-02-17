import { useContext, useState } from 'react'
import { Screen, ScreenContext } from '../contexts/screen-context'
import Request from '../helpers/http'

const styles = {
  container: 'w-full h-full flex flex-col items-center justify-center z-20 text-white',
  stack: 'flex flex-col gap-2',
}

const LogIn = () => {
  const { navigate } = useContext(ScreenContext)

  /* LOG-IN STATES */
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [errors, setErrors] = useState({})

  const onSubmit = e => {
    e.preventDefault()

    const no_username = !username || username.trim().length <= 0
    const no_password = !password || password.trim().length <= 0
    const short_username = username.trim().length < 4
    const short_password = password.trim().length < 6

    if (!no_username && !no_password && !short_username && !short_password) {
      /* TODO: LOGIN API WITH REQUEST CLASS */
      navigate(Screen.Menu)
    } else {
      setErrors({
        username: no_username ? 'You have not entered your username.' :
          short_username && 'Username is too short (at least 4 characters).',
        password: no_password ? 'No password was entered.' :
          short_password && 'Password is too short (at least 6 characters).'
      })
    }
  }

  return (
    <div className={styles.container}>
      <div className='flex flex-col items-center mb-28'>
        <h2 className='text-9xl font-semibold tetris-text'>
          Tetris
        </h2>
        <h2 className='text-8xl green-grad-text pb-5'>
          Battle Royale
        </h2>
        <p className='text-2xl opacity-30'>
          Massively Multiplayer Classic Tetris
        </p>
      </div>
      <form
        onSubmit={onSubmit}
        className={styles.stack}>
        <label>
          Username
          <span className='text-sm text-red-800'>
            {errors?.username}
          </span>
        </label>
        <input
          type='text'
          value={username}
          onChange={e => setUsername(e.target.value)}
          placeholder='Username'
        />
        <label>
          Password
          <span className='text-sm text-red-800'>
            {errors?.password}
          </span>
        </label>
        <input
          type='password'
          value={password}
          onChange={e => setPassword(e.target.value)}
          placeholder='Password'
        />
        <button
          type='submit'
          className='border py-4 rounded mt-10 transition-all hover:bg-[#19a186] hover:text-black'>
          Authenticate
        </button>
      </form>
    </div>
  )
}

export default LogIn