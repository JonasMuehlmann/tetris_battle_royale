import React, { Component, useContext, useEffect, useState } from "react";

const QueueContext = React.createContext()

const QueueProvider = ({ children }) => {
  const [isInQueue, setIsInQueue] = useState()
  const [elapsed, setElapsed] = useState(1)

  useEffect(() => {
    if (isInQueue) {
      const unset = setInterval(() => {
        setElapsed(elapsed + 1)
      }, 1000)

      return () => {
        clearInterval(unset)
      }
    } else {
      setElapsed(0)
    }
  }, [isInQueue, elapsed])

  return (
    <QueueContext.Provider value={{
      isInQueue,
      setIsInQueue,
      elapsed,
    }}>
      {children}
    </QueueContext.Provider>
  )
}

export const withQueue = Component => ({ ...props }) => (
  <QueueProvider>
    <Component {...props} />
  </QueueProvider>
)

export const useQueue = () => useContext(QueueContext)