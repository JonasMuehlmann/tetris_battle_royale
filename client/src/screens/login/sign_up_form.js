import { motion } from "framer-motion"
import { useState } from "react"
import { useAuth } from "../../contexts/auth-context"

/*
 * DEFAULTS
 */
const MODEL = Object.freeze({
  username: '',
  password: '',
  passwordReenter: '',
})

const SignUpForm = ({ onSignIn }) => {
  const {
    errors,
    signUp,
    isValid,
  } = useAuth()

  /*
   * STATES
   */
  const [model, setModel] = useState({ ...MODEL })

  return (
    <motion.form
      initial={{ opacity: 0, x: window.innerWidth / 2, scale: 0 }}
      animate={{ opacity: 1, x: 0, scale: 1 }}
      exit={{ opacity: 0, x: window.innerWidth / 2, scale: 0 }}
      transition={{ type: 'spring', duration: 1.5 }}
      className='flex flex-col gap-2'
      autoComplete="off"
      onSubmit={e => {
        e.preventDefault()
        if (isValid(model)) {
          signUp(model)
        }
      }}>
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
        autoComplete="off"
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
        autoComplete="off"
        onChange={e => {
          setModel({ ...model, password: e.target.value })
        }}
        placeholder='Password'
        className='border-4 border-[#19a186]'
      />
      <label>
        Re-enter Password
        <p className='text-sm text-red-400'>
          {errors?.passwordReenter}
        </p>
      </label>
      <input
        type='password'
        autoComplete="off"
        value={model.passwordReenter}
        onChange={e => {
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