import { useState } from "react"
import { motion } from 'framer-motion'

/*
 * DEFAULT
 */
const MODEL = Object.freeze({
  username: '',
  password: '',
})

const SignInForm = (
  {
    onSubmit = model => { },
    onSignUp = () => { },
  }) => {
  /*
   * STATES
   */
  const [model, setModel] = useState(MODEL)
  const [errors, setErrors] = useState({})

  // #region METHODS

  const isValid = () => {
    const { username, password } = model
    const no_username = !username || username.trim().length <= 0
    const no_password = !password || password.trim().length <= 0
    const short_username = username.trim().length < 4
    const short_password = password.trim().length < 6

    const valid = !no_username && !no_password && !short_username && !short_password

    if (!valid) {
      setErrors({
        username: no_username ? 'You have not entered your username.' :
          short_username && 'Username is too short (at least 4 characters).',
        password: no_password ? 'No password was entered.' :
          short_password && 'Password is too short (at least 6 characters).'
      })
    } else {
      setErrors({})
    }

    return valid
  }

  // #endregion

  const Actions = () => (
    <>
      <button
        type='submit'
        className={`border-2 py-4 rounded mt-8 transition-all
          text-lg hover:bg-[#19a186] hover:text-black`}>
        Authenticate
      </button>
      <button
        type="button"
        onClick={() => onSignUp()}
        className='opacity-60 hover:opacity-100 josefin text-md mt-2'>
        Don't have an account?
        <span className='text-[#19a186] px-1'>
          Sign up here
        </span>
      </button>
    </>
  )

  // #endregion

  return (
    <motion.form
      initial={{ opacity: 0, x: -window.innerWidth / 2, scale: 0 }}
      animate={{ opacity: 1, x: 0, scale: 1 }}
      exit={{ opacity: 0, x: -window.innerWidth / 2, scale: 0 }}
      transition={{ type: 'spring', duration: 1.5 }}
      onSubmit={e => {
        e.preventDefault()
        if (isValid()) {
          onSubmit(model)
        }
      }}
      className='flex flex-col gap-2'>
      <label className='flex flex-col'>
        Username
        <p className='text-sm text-red-400'>
          {errors?.username}
        </p>
      </label>
      <input
        type='text'
        value={model.username}
        onChange={e => {
          setModel({ ...model, username: e.target.value })
        }}
        placeholder='Username'
        className='border-4 border-[#19a186]'
      />
      <label>
        Password
        <p className='text-sm text-red-400'>
          {errors?.password}
        </p>
      </label>
      <input
        type='password'
        value={model.password}
        onChange={e => {
          setModel({ ...model, password: e.target.value })
        }}
        placeholder='Password'
        className='border-4 border-[#19a186]'
      />
      <Actions />
    </motion.form>
  )
}

export default SignInForm