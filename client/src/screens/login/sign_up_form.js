import { motion } from "framer-motion"
import { useState } from "react"

const SignUpForm = (
  {
    onSubmit = model => { },
    onSignIn = () => { },
  }) => {
  const [errors, setErrors] = useState({})
  const [model, setModel] = useState({
    username: '',
    password: '',
    passwordReenter: '',
    isValid: false,
  })

  const isModelValid = () => {
    const { username, password, passwordReenter } = model
    const no_username = !username || username.trim().length <= 0
    const no_password = !password || password.trim().length <= 0
    const no_passwordReenter = !passwordReenter || passwordReenter.trim().length <= 0
    const short_username = username.trim().length < 4
    const short_password = password.trim().length < 6
    const not_matching_password = password !== passwordReenter

    const valid = !no_username &&
      !no_password &&
      !no_passwordReenter &&
      !not_matching_password &&
      !short_username &&
      !short_password

    if (!valid) {
      setErrors({
        username: no_username ? 'You have not entered your username.' :
          short_username && 'Username is too short (at least 4 characters).',
        password: no_password ? 'No password was entered.' :
          short_password && 'Password is too short (at least 6 characters).',
        passwordReenter: no_passwordReenter ? 'Please enter your password one more time.' :
          not_matching_password && 'Passwords do not match',
      })
    } else {
      setErrors({})
    }

    return valid
  }

  return (
    <motion.form
      initial={{ opacity: 0, x: 200 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 200 }}
      transition={{ type: 'spring', duration: 1 }}
      autoComplete="off"
      onSubmit={e => {
        e.preventDefault()
        if (isModelValid()) {
          onSubmit(model)
        }
      }}
      className='flex flex-col gap-2'>
      <input autoComplete="off" name="hidden" type="text" style={{ display: 'none' }} />
      <label className='flex flex-col'>
        Username
        <span className='text-sm text-red-400'>
          {errors?.username}
        </span>
      </label>
      <input
        type='text'
        value={model.username}
        onChange={e => {
          if (e.target || e.target.value !== ' ') {
            setModel({ ...model, username: e.target.value })
          }
        }}
        placeholder='Username'
        autoComplete="off"
        className='border-4 border-[#19a186]'
      />
      <label>
        Password
        <span className='text-sm text-red-800'>
          {errors?.password}
        </span>
      </label>
      <input
        type='password'
        value={model.password}
        autoComplete="off"
        onChange={e => {
          if (e.target || e.target.value !== ' ')
            setModel({ ...model, password: e.target.value })
        }}
        placeholder='Password'
        className='border-4 border-[#19a186]'
      />
      <label>
        Re-enter Password
        <span className='text-sm text-red-800'>
          {errors?.passwordReenter}
        </span>
      </label>
      <input
        type='password'
        autoComplete="off"
        value={model.passwordReenter}
        onChange={e => {
          if (e.target || e.target.value !== ' ')
            setModel({ ...model, passwordReenter: e.target.value })
        }}
        placeholder='Reenter password.'
        className='border-4 border-[#19a186]'
      />
      <button
        type='submit'
        className={`border-2 py-4 rounded mt-8 transition-all
          text-lg hover:bg-[#19a186] hover:text-black`}>
        Register
      </button>
      <button
        type="button"
        onClick={() => onSignIn()}
        className='opacity-60 hover:opacity-100 josefin text-md mt-1'>
        Have already account?
        <span className='text-[#19a186] px-1'>
          Sign in
        </span>
      </button>
    </motion.form>
  )
}

export default SignUpForm