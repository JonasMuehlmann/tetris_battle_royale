import React, { useContext, useState } from "react";

/**
 * SCREEN DEFINITIONS
 * MUST BE PROVIDED AS ARGUMENT TO NAVIGATE
 */
export class Screen {
  static LogIn = new Screen("login")
  static Menu = new Screen("menu")
  static Queue = new Screen("queue")
  static Tetris = new Screen("tetris")
  static Result = new Screen("result")
  static Profile = new Screen("profile")
  static Settings = new Screen("settings")

  constructor(name, component = null) {
    this.name = name
    this.component = component
  }
}

const ScreenContext = React.createContext()

/**
 * SCREEN CONTEXT PROVIDER
 * PROVIDES METHODES AND STATES RELATED TO SCREENS
 */
export const ScreenProvider = ({ children }) => {
  const [currentScreen, setCurrentScreen] = useState(Screen.Menu)

  const navigate = screen => {
    if (!(screen instanceof Screen)) return;

    setCurrentScreen(screen)
  }

  return (
    <ScreenContext.Provider value={{ currentScreen, navigate }}>
      {children}
    </ScreenContext.Provider>
  )
}

export const withScreenContext = Component => ({ ...props }) => (
  <ScreenProvider>
    <Component {...props} />
  </ScreenProvider>
)

/**
 * DELIEVERS 'VALUE'-OBJECT OF THE PROVIDER
 */
export const useScreens = () => useContext(ScreenContext)