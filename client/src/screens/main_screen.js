import { motion } from "framer-motion";
import { Screen, useScreens, withScreenContext } from "../contexts/screen-context"
import ErrorScreen from "./404";
import LogInScreen from "./login";
import LobbyScreen from "./lobby";
import FriendsBox from "../components/friends_box";
import TetrisScreen from "./tetris";
import { KeybindsContextProvider } from "../contexts/keybinds-context";
import { WebSocketProvider } from "../contexts/websocket-context";
import { AuthContext, AuthProvider } from "../contexts/auth-context";
import { QueueProvider } from "../contexts/queue-context";

const MainScreen = () => {
  const {
    currentScreen,
    navigate,
  } = useScreens()

  const renderCurrentScreen = () => {
    switch (currentScreen?.name) {
      case 'login':
        return (
          <motion.div
            className="w-screen h-screen"
            exit={{ opacity: 0, scale: .35 }}
            transition={{ duration: 1.5, type: 'spring' }}
            key={currentScreen.name} >
            <LogInScreen />
          </motion.div>
        )
      case 'menu':
        return (
          <motion.div
            className="w-screen h-screen"
            initial={{ opacity: 0, y: -window.innerHeight }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -window.innerHeight }}
            transition={{ duration: 1.5, type: 'spring', delay: .25 }}
            key={currentScreen.name}>
            <LobbyScreen />
            <FriendsBox />
          </motion.div>
        )
      case 'tetris':
        return (
          <motion.div
            className="w-screen h-screen"
            initial={{ opacity: 0, scale: 0 }}
            animate={{ opacity: 1, scale: 1 }}
            exit={{ opacity: 0, scale: 0 }}
            transition={{ duration: 1.5, type: 'spring' }}
            key={currentScreen.name}>
            <TetrisScreen />
          </motion.div>
        )
      default:
        return <></>
    }
  }

  return (
    <AuthProvider>
      <AuthContext.Consumer>
        {
          ({ user }) => (
            <WebSocketProvider user={user}>
              <QueueProvider user={user}>
                <KeybindsContextProvider>
                  <div className="z-20">
                    {
                      currentScreen ?
                        (
                          renderCurrentScreen()
                        ) :
                        (
                          <ErrorScreen
                            onNavigate={() => navigate(Screen.LogIn)}
                          />
                        )
                    }
                  </div>
                </KeybindsContextProvider>
              </QueueProvider>
            </WebSocketProvider>
          )
        }
      </AuthContext.Consumer>
    </AuthProvider>
  )
}

export default withScreenContext(MainScreen)