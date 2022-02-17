import { useState } from "react"

const SignInForm = (
  {
    onSubmit = model => { },
    onSignUp = () => { },
  }) => {
  const [errors, setErrors] = useState({})
  const [model, setModel] = useState({
    username: '',
    password: '',
    isValid: false,
  })

  const isModelValid = () => {
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

  return (
    <form
      onSubmit={e => {
        e.preventDefault()
        if (isModelValid()) {
          onSubmit(model)
        }
      }}
      className='flex flex-col gap-2'>
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
        onChange={e => {
          if (e.target || e.target.value !== ' ')
            setModel({ ...model, password: e.target.value })
        }}
        placeholder='Password'
        className='border-4 border-[#19a186]'
      />
      <button
        type='submit'
        className={`border-2 py-4 rounded mt-8 transition-all
          text-lg hover:bg-[#19a186] hover:text-black`}>
        Authenticate
      </button>
      <button
        type="button"
        onClick={() => onSignUp()}
        className='opacity-60 hover:opacity-100 josefin text-md mt-1'>
        Don't have an account?
        <span className='text-[#19a186] px-1'>
          Sign up here
        </span>
      </button>
    </form>
  )
}

export default SignInForm