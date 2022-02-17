import { useState } from 'react'

const styles = {
  container: 'w-full h-full flex flex-col items-center justify-center',
  stack: 'flex flex-col gap-2',
}

const LogIn = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [errors, setErrors] = useState({})

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