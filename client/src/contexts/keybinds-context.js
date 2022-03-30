import React, { useContext, useState } from 'react'

const KeybindsContext = React.createContext()

export const KeybindsContextProvider = ({ children }) => {

  const [keybinds, setKeybinds] = useState({
    left: {
      label: 'Move Left',
      key: 37,
      keyName: "Left",
    },
    right: {
      label: 'Move Right',
      key: 39,
      keyName: "Right",
    },
    rotate: {
      label: "Rotate Clockwise",
      key: 38,
      keyName: "Up",
    },
    drop: {
      label: "Drop: Soft",
      key: 40,
      keyName: "Down",
    },
    dropHard: {
      label: "Drop: Hard",
      key: 32,
      keyName: "Space",
    },
    hold: {
      label: "Hold Piece",
      key: 16,
      keyName: "Right Shift"
    }
  })

  const updateKeybind = (key, value) => {
    setKeybinds({
      ...keybinds,
      [key]: value,
    })
  }

  return (
    <KeybindsContext.Provider
      value={{
        keybinds,
        updateKeybind
      }}>
      {children}
    </KeybindsContext.Provider>
  )
}

export const useKeybinds = () => useContext(KeybindsContext)