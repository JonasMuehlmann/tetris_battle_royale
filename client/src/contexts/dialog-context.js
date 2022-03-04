import React, { useContext, useState } from "react";
import Loader from "../components/loader";

export class DialogType {
  static Warning = new DialogType("warning")
  static Info = new DialogType("info")

  constructor(type) {
    this.type = type
  }
}

export const DialogContext = React.createContext()

export const DialogProvider = ({ children }) => {
  const [isDialogVisible, setIsDialogVisible] = useState(false)
  const [model, setModel] = useState({})

  const showDialog = ({
    title = 'Dialog Title',
    content = 'dialog content',
  }) => {
    setModel({ title, content })
    setIsDialogVisible(true)
  }

  const hideDialog = () => {
    setIsDialogVisible(false)
    setModel({})
  }

  return (
    <DialogContext.Provider value={{
      showDialog,
      hideDialog,
      component: {
        isDialogVisible,
        model,
      },
    }}>
      <Loader />
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