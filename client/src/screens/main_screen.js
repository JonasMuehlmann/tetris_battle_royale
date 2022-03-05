import { AnimatePresence, motion } from "framer-motion";
import { useContext } from "react"
import { withDialogContext } from "../contexts/dialog-context";
import { Screen, ScreenContext } from "../contexts/screen-context"
import ErrorScreen from "./404";
import LogIn from "./login";
import Menu from "./menu";

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
              transition={{ duration: 3, type: 'spring' }}
              key={currentScreen.name} >
              <LogIn />
            </motion.div>
          )
        }
        {
          currentScreen?.name === 'menu' &&
          (
            <motion.div
              className="w-screen h-screen"
              initial={{ opacity: 0, x: window.innerWidth }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: window.innerWidth }}
              transition={{ duration: 3, type: 'spring' }}
              key={currentScreen.name}>
              <Menu />
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