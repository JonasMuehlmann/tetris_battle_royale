import { useState } from 'react'
import Request from '../helpers/http'

const styles = {
  container: 'w-full h-full flex flex-col items-center justify-center',
  stack: 'flex flex-col gap-2',
}

const LogIn = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [errors, setErrors] = useState({})

  const onSubmit = () => {
    const no_username = !username || username.trim().length <= 0
    const no_password = !password || password.trim().length <= 0
    const short_username = username.trim().length < 4
    const short_password = password.trim().length < 6

    if (!no_username && !no_password && !short_username && !short_password) {
      /* TODO: LOGIN API WITH REQUEST CLASS */
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
      <div className={styles.stack}>
        <h2>
          Tetris Battle Royale
        </h2>
      </div>
      <form className={styles.stack}>
        <label>
          Username
          <span className='text-sm text-red-800'>
            {errors.username}
          </span>
        </label>
        <input
          type='text'
          value={username}
          onChange={e => setUsername(e.target.value)}
        />
        <label>
          Password
          <span className='text-sm text-red-800'>
            {errors.password}
          </span>
        </label>
        <input
          type='password'
          value={password}
          onChange={e => setUsername(e.target.value)}
        />
      </form>
    </div>
  )
}

export default LogIn