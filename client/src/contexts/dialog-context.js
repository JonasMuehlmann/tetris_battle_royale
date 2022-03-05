import React, { useContext, useState } from "react";
import Loader from "../components/loader";

export class DialogType {
  static Warning = new DialogType("warning")
  static Info = new DialogType("info")
  static Load = new DialogType("load")
  static Authenticate = new DialogType(
    "authenticate",
    "Authenticating..",
    "It won\'t take long.."
  )

  constructor(type, title, content) {
    this.type = type
    this.title = title
    this.content = content
  }
}

export const DialogContext = React.createContext()

export const DialogProvider = ({ children }) => {
  const [isDialogVisible, setIsDialogVisible] = useState(false)
  const [currentType, setCurrentType] = useState(null)

  const showDialog = (type) => {
    setCurrentType(type)
    setIsDialogVisible(true)
  }

  const hideDialog = () => {
    setIsDialogVisible(false)
  }

  return (
    <DialogContext.Provider value={{
      showDialog,
      hideDialog,
      component: {
        isDialogVisible,
        currentType,
      },
    }}>
      {children}
    </DialogContext.Provider>
  )
}

export const withDialogContext = Component => ({ ...props }) => (
  <DialogProvider>
    <Component {...props} />
  </DialogProvider>
)

export const useDialog = () => useContext(DialogContext)