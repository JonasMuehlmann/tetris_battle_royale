import { AnimatePresence, motion } from "framer-motion";
import { useContext } from "react"
import { withDialogContext } from "../contexts/dialog-context";
import { Screen, ScreenContext } from "../contexts/screen-context"
import ErrorScreen from "./404";
import LogInScreen from "./login";
import LobbyScreen from "./lobby";

const MainScreen = (props) => {
  const { currentScreen, navigate } = useContext(ScreenContext)

  const renderCurrentScreen = () => {
    return (
      <AnimatePresence
        exitBeforeEnter>
        {
          currentScreen?.name === 'login' &&
          (
            <motion.div
              className="w-screen h-screen"
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: -window.innerWidth }}
              transition={{ duration: 2, type: 'spring' }}
              key={currentScreen.name} >
              <LogInScreen />
            </motion.div>
          )
        }
        {
          currentScreen?.name === 'menu' &&
          (
            <motion.div
              className="w-screen h-screen"
              initial={{ opacity: 0, y: -window.innerHeight }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, x: window.innerWidth }}
              transition={{ duration: 2, type: 'spring' }}
              key={currentScreen.name}>
              <LobbyScreen />
            </motion.div>
          )
        }
      </AnimatePresence>
    )
  }

  return (
    currentScreen ?
      (
        renderCurrentScreen()
      ) :
      (
        <ErrorScreen
          onNavigate={() => navigate(Screen.LogIn)}
        />
      )
  );
}

export default withDialogContext(MainScreen)