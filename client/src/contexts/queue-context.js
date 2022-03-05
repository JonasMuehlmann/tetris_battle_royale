import React, { Component, useContext, useEffect, useState } from "react";

const QueueContext = React.createContext()

const QueueProvider = ({ children }) => {
  const [isInQueue, setIsInQueue] = useState(false)
  const [queueType, setQueueType] = useState(null)
  const [elapsed, setElapsed] = useState(0)

  const requestQueue = request => {
    setQueueType(request)
    setIsInQueue(true)
  }

  const cancelQueue = () => {
    setQueueType(null)
    setElapsed(0)
    setIsInQueue(false)
  }

  useEffect(() => {
    if (isInQueue) {
      const queueTimer = setInterval(() => {
        setElapsed(elapsed + 1)
      }, 1000)

      return () => {
        clearInterval(queueTimer)
      }
    } else {
      setElapsed(0)
    }
  }, [isInQueue, queueType, elapsed])

  return (
    <QueueContext.Provider value={{
      isInQueue,
      elapsed,
      queueType,
      requestQueue,
      cancelQueue,
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