import React, { useContext, useEffect, useRef, useState } from "react";
import axios from 'axios'
import { useWS } from "./websocket-context";

export const QueueContext = React.createContext()

/**
 * QUEUE CONTEXT PROVIDER
 * PROVIDES METHODES AND STATES RELATED TO CURRENT QUEUE
 */
export const QueueProvider = ({ children, user }) => {
  // INITIALIZE ONLY IF USER
  if (user === undefined || user === null) {
    return (<>{children}</>)
  }

  const {
    lastJsonMessage
  } = useWS()

  const queueTimer = useRef()

  const [currentUser, setCurrentUser] = useState(user)
  const [currentMatch, setCurrentMatch] = useState(null)

  const [isInQueue, setIsInQueue] = useState(false)
  const [queueType, setQueueType] = useState(null)
  const [elapsed, setElapsed] = useState(0)

  async function requestQueue(request) {
    if (currentUser === undefined || currentUser === null) return

    setQueueType(request)

    setIsInQueue(true)

    await axios.post('/join', {
      userID: currentUser.id
    }).catch(err => {
      console.error(err)
      setIsInQueue(false)
    })
  }

  function cancelQueue() {
    if (!isInQueue
      || currentUser === undefined || currentUser === null) return

    await axios.post('/leave', {
      userID: currentUser.id
    }).catch(err => {
      console.error(err)
    })

    setQueueType(null)
    setElapsed(0)
    setIsInQueue(false)
  }

  useEffect(() => {
    // MATCH FOUND
    if (lastJsonMessage?.matchID) {
      setCurrentMatch({ ...lastJsonMessage })
      clearInterval(queueTimer.current)
    }

    // QUEUE JOINED
    if (isInQueue) {
      queueTimer.current = setInterval(() => {
        setElapsed(elapsed + 1)
      }, 1000)
    }

    return () => {
      if (queueTimer.current) clearInterval(queueTimer.current)
    }
  }, [
    lastJsonMessage,
    isInQueue,
    elapsed,
    currentMatch,
    setCurrentMatch,
    setCurrentUser,
    setElapsed,
  ])

  return (
    <QueueContext.Provider value={{
      isInQueue,
      elapsed,
      queueType,
      currentMatch,
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

/**
 * DELEIVERS 'VALUE'-OBJECT OF THE PROVIDER
 */
export const useQueue = () => useContext(QueueContext)