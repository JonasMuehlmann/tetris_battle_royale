import axios from 'axios'
import React, { useContext, useState } from 'react'
import { useDialog, DialogType } from './dialog-context'
import { useScreens } from './screen-context'

const ENDPOINT_BASE = '/'
const ENDPOINT_SEGMENT_LOGIN = ENDPOINT_BASE.concat('login')
const ENDPOINT_SEGMENT_REGISTER = ENDPOINT_BASE.concat('register')

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

      if (response) {
        console.log(response)
        setUser(response)
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
      const user = await axios.post(
        ENDPOINT_SEGMENT_REGISTER,
        {
          username,
          password,
        }
      )
      console.log(user)
      if (user) {
        setUser(user)
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
        signIn,
        signUp
      }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)