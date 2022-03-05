import React, { useState } from "react";
import LogIn from '../screens/login/index'
import Menu from "../screens/menu/index";

export class Screen {
  static LogIn = new Screen("login", LogIn)
  static Menu = new Screen("menu", Menu)
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

export const ScreenContext = React.createContext()

export const ScreenProvider = ({ children }) => {
  const [currentScreen, setCurrentScreen] = useState(Screen.LogIn)

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