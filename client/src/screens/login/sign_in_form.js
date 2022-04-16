import { useState } from "react"
import { motion } from 'framer-motion'
import { useAuth } from "../../contexts/auth-context"

/*
 * DEFAULT
 */
const MODEL = Object.freeze({
  username: '',
  password: '',
})

const SignInForm = ({ onSignUp }) => {
  const {
    errors,
    signIn,
    isValid,
  } = useAuth()

  /*
   * STATES
   */
  const [model, setModel] = useState({ ...MODEL })

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
        onClick={() => onSignUp?.()}
        className='opacity-60 hover:opacity-100 josefin text-md mt-2'>
        Don't have an account?
        <span className='text-[#19a186] px-1'>
          Sign up here
        </span>
      </button>
    </>
  )

  return (
    <motion.form
      initial={{ opacity: 0, x: -window.innerWidth / 2, scale: 0 }}
      animate={{ opacity: 1, x: 0, scale: 1 }}
      exit={{ opacity: 0, x: -window.innerWidth / 2, scale: 0 }}
      transition={{ type: 'spring', duration: 1.5 }}
      onSubmit={e => {
        e.preventDefault()
        if (isValid(model)) {
          // #REMOVE TRUE TO BYPASS (DEV)
          signIn(model, false)
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