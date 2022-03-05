import React, { Component, useContext, useState } from "react";

export class MenuItem {
  static Matchfinder = new MenuItem(
    'Matchfinder'
  )
  static Statistics = new MenuItem(
    'Stats'
  )
  static PlayerProfile = new MenuItem(
    'Profile'
  )
  static PlayerSettings = new MenuItem(
    'Settings'
  )

  constructor(text) {
    this.text = text
  }
}

const MenuContext = React.createContext()

const MenuContextProvider = ({ children }) => {
  const [currentMenu, setCurrentMenu] = useState(MenuItem.Matchfinder)

  return (
    <MenuContext.Provider value={{
      currentMenu,
      setCurrentMenu,
    }}>
      {children}
    </MenuContext.Provider>
  )
}

export const withMenuContext = Component => ({ ...props }) => (
  <MenuContextProvider>
    <Component {...props} />
  </MenuContextProvider>
)

export const useMenu = () => useContext(MenuContext)