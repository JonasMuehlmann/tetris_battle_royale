import React, { useState, useEffect, useContext } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';

const WS_BASE = 'ws://'
const WS_URL = WS_BASE.concat('localhost:8080')

const WebSocketContext = React.createContext();

export const WebSocketProvider = ({ children, user }) => {
  const [URL, setURL] = useState(WS_URL)
  const [history, setHistory] = useState([])

  const {
    sendMessage, // NEEDS TO BE TESTED
    sendJsonMessage, // NEEDS TO BE TESTED
    lastMessage, // NEEDS TO BE TESTED
    lastJsonMessage, // NEEDS TO BE TESTED
    readyState
  } = useWebSocket(URL, {
    onOpen: () => {
      console.log(`WS-con opened with ${user?.id}`)
      // SENDING USER ID BACK FOR REGISTER
      sendJsonMessage({ userID: user?.id }, false)
    },
    shouldReconnect: closeEvent => true,
    retryOnError: true,
  }, user !== null)

  useEffect(() => {
    if (user === undefined || user === null) return

    // ON MESSAGE
    if (lastMessage !== null || lastJsonMessage !== null) {
      setHistory(prev => prev.concat(lastMessage || lastJsonMessage))
    }
  }, [lastMessage, setHistory, user])

  const connectionStatus = {
    [ReadyState.CONNECTING]: 'Connecting',
    [ReadyState.OPEN]: 'Open',
    [ReadyState.CLOSING]: 'Closing',
    [ReadyState.CLOSED]: 'Closed',
    [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
  }[readyState];

  // INITIALIZE ONLY IF USER
  if (user === undefined || user === null) {
    return (<>{children}</>)
  }

  return (
    <WebSocketContext.Provider
      value={{
        sendJsonMessage,
        lastJsonMessage,
        connectionStatus
      }}>
      {children}
    </WebSocketContext.Provider>
  )
}

export const useWS = () => useContext(WebSocketContext)