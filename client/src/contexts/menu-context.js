import React, { Component, useContext, useState } from "react";

/**
 * MENU-SCREENS IN LOBBY DEFINITIONS
 * MUST BE PROVIDED TO NAVIGATE IN BETWEEN
 */
export class MenuItem {
  static Matchfinder = new MenuItem(
    'Match'
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

/**
 * MENU CONTEXT PROVIDER
 * PROVIDES METHODES AND STATES RELATED TO LOBBY-MENU-SCREEN
 */
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

/**
 * DELEIVERS 'VALUE'-OBJECT OF THE PROVIDER
 */
export const useMenu = () => useContext(MenuContext)