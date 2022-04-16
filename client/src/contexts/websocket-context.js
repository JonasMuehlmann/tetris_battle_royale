import React, { useState, useEffect, useContext, useRef } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';

const WS_BASE = 'ws://'
const WS_URL = WS_BASE.concat('localhost:8080/ws')

const WebSocketContext = React.createContext();

export const WebSocketProvider = ({ children, user }) => {
  const [URL, setURL] = useState(WS_URL)
  const [history, setHistory] = useState([])

  const eventNotice = useRef()
  const matchStartNotice = useRef()
  const scoreboard = useRef()
  
  const tetrominoPreview = useRef()
  const tetrominoState = useRef()
  const tetrominoLockIn = useRef()
  const tetrominoSpawn = useRef()
  const clearRowIndex = useRef()
  const scoreGain = useRef()
  const eliminatedPlayerID = useRef()

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
    if (lastJsonMessage !== null) {
      setHistory(prev => prev.concat(lastJsonMessage))

      switch (lastJsonMessage.type) {
        case "MatchStartNotice":
          const { matchID, opponents } = lastJsonMessage
          matchStartNotice.current = { matchID, opponents }
          break
        case "StartTetrominoPreview":
          const { newPreview } = lastJsonMessage
          tetrominoPreview.current = [...newPreview]
          break
        case "UpdatedTetrominoState":
          const { tetrominoPosition, rotationChange } = lastJsonMessage.newState
          tetrominoState.current = { tetrominoPosition, rotationChange }
          break
        case "TetrominoLockinNotice":
          const { lockIn } = lastJsonMessage
          tetrominoLockIn.current = lockIn
          break
        case "RowClearNotice":
          const { rowNum } = lastJsonMessage
          clearRowIndex.current = rowNum
          break
        case "TetrominoSpawnNotice":
          const { newTetromino, enqueuedTetromino } = lastJsonMessage
          tetrominoSpawn.current = { newTetromino, enqueuedTetromino }
          break
        case "ScoreGain":
          const { score } = lastJsonMessage
          scoreGain.current = score
          break
        case "EliminationNotice":
          const { eliminatedPlayer } = lastJsonMessage
          eliminatedPlayerID.current = eliminatedPlayer
          break
        case "EventNotice":
          const { event } = lastJsonMessage
          eventNotice.current = event
          break
        case "StartTetrominoPreview":
          const { endOfMatchData } = lastJsonMessage
          scoreboard.current = endOfMatchData?.scoreboard
        default:
          break
      }
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
        connectionStatus,
        matchStartNotice,
        tetrominoPreview,
        /* FILTERED MESSAGES */
        eventNotice: eventNotice.current,
        matchStartNotice: matchStartNotice.current,
        scoreboard: scoreboard.current,
        tetrominoPreview: tetrominoPreview.current,
        tetrominoState: tetrominoState.current,
        tetrominoLockIn: tetrominoLockIn.current,
        tetrominoSpawn: tetrominoSpawn.current,
        clearRowIndex: clearRowIndex.current,
        scoreGain: scoreGain.current,
        eliminatedPlayerID: eliminatedPlayerID.current,
      }}>
      {children}
    </WebSocketContext.Provider>
  )
}

export const useWS = () => useContext(WebSocketContext)