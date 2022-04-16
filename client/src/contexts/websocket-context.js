import React, { useState, useCallback, useEffect, useContext } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';

const WebSocketContext = React.createContext();

export const WebSocketProvider = ({ children, user }) => {
  if (user === undefined || user === null) {
    return (
      <>
        {children}
      </>
    )
  }

  const [URL, setURL] = useState('ws://localhost:8080')
  const [history, setHistory] = useState([])

  const {
    sendMessage, // NEEDS TO BE TESTED
    sendJsonMessage, // NEEDS TO BE TESTED
    lastMessage, // NEEDS TO BE TESTED
    lastJsonMessage, // NEEDS TO BE TESTED
    readyState
  } = useWebSocket(URL, {
    onOpen: () => console.log(`WS-con opened with ${user?.id}`),
    shouldReconnect: closeEvent => true,
    retryOnError: true,
  })

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