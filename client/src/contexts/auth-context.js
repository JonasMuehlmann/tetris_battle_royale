import axios from 'axios'
import React, { useContext, useState } from 'react'
import { useDialog, DialogType } from './dialog-context'
import { useScreens } from './screen-context'

const ENDPOINT_BASE = '/user'
const ENDPOINT_SEGMENT_LOGIN = ENDPOINT_BASE.concat('/login')
const ENDPOINT_SEGMENT_REGISTER = ENDPOINT_BASE.concat('/register')

export const AuthContext = React.createContext()

export const AuthProvider = ({ children }) => {
  const {
    navigate,
  } = useScreens()

  const {
    showDialog,
    hideDialog,
  } = useDialog()

  const [user, setUser] = useState(null)
  const [errors, setErrors] = useState({})

  function isValid(model) {
    const { username, password, passwordReenter } = model
    const no_username = !username || username.trim().length <= 0
    const no_password = !password || password.trim().length <= 0
    const short_username = username.trim().length < 4
    const short_password = password.trim().length < 6
    let no_passwordReenter, not_matching_password

    let valid = !no_username && !no_password && !short_username && !short_password

    if (passwordReenter !== undefined) {
      no_passwordReenter = !passwordReenter || passwordReenter.trim().length <= 0
      not_matching_password = password !== passwordReenter
      valid &= no_passwordReenter && not_matching_password
    }

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

  async function signIn(model, bypass = false) {
    const { username, password } = model
    showDialog(DialogType.Authenticate)

    // DEVELOPER OPTION
    if (bypass) {
      setTimeout(() => { hideDialog(); navigate(Screen.Menu) }, 1500)
      return
    }

    try {
      const response = await axios.post(
        ENDPOINT_SEGMENT_LOGIN,
        {
          username,
          password,
        })

      if (response.status === 200) {
        const { sessionID, userID, username } = response.data
        setUser({
          sessionId: sessionID,
          id: userID,
          username
        })
      }
    } catch (error) {
      console.error(error.message)
    } finally {
      hideDialog()
    }
  }

  async function signUp(model) {
    const { username, password } = model

    try {
      showDialog(DialogType.Authenticate)
      const response = await axios.post(
        ENDPOINT_SEGMENT_REGISTER,
        { username, password }
      )

      if (response.status === 200) {
        const { sessionID, userID, username } = response.data
        setUser({
          sessionId: sessionID,
          id: userID,
          username
        })
      }
    } catch (error) {
      console.info(error)
    } finally {
      hideDialog()
    }
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        errors,
        signIn,
        signUp,
        isValid,
      }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)