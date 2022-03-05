import { AnimatePresence, motion } from 'framer-motion'
import { MenuItem, useMenu, withMenuContext } from "../../contexts/menu-context";
import { withQueue } from "../../contexts/queue-context";
import Matchfinder from "./matchfinder";
import Statistics from "./statistics";
import PlayerProfile from "./player_profile";
import PlayerSettings from "./player_settings";
import Menu from "../../components/menu";
import QueueBox from "../../components/queue_box";

const LobbyScreen = () => {
  const {
    currentMenu,
  } = useMenu()

  const Motions = {
    SlideLeft: {
      initial: { opacity: 0, x: -window.innerWidth },
      animate: { opacity: 1, x: 0 },
      exit: { opacity: 0, x: -window.innerWidth },
      transition: { duration: .75 }
    },
    SlideDown: {
      initial: { opacity: 0, y: window.innerHeight },
      animate: { opacity: 1, y: 0 },
      exit: { opacity: 0, y: window.innerHeight },
      transition: { duration: .75 }
    },
  }

  const renderCurrentMenu = () => (
    <AnimatePresence exitBeforeEnter>
      {
        currentMenu === MenuItem.Matchfinder ?
          (
            <motion.div
              {...Motions.SlideLeft}
              key={currentMenu.text}>
              <Matchfinder />
            </motion.div>
          ) :
          currentMenu === MenuItem.Statistics ?
            (
              <motion.div
                {...Motions.SlideLeft}
                key={currentMenu.text}>
                <Statistics />
              </motion.div>
            ) :
            currentMenu === MenuItem.PlayerProfile ?
              (
                <motion.div
                  {...Motions.SlideDown}
                  key={currentMenu.text}>
                  <PlayerProfile />
                </motion.div>
              ) :
              currentMenu === MenuItem.PlayerSettings &&
              (
                <motion.div
                  {...Motions.SlideDown}
                  key={currentMenu.text}>
                  <PlayerSettings />
                </motion.div>
              )
      }
    </AnimatePresence>
  )

  return (
    <div className="w-full h-full flex flex-col z-20 relative">
      <Menu />
      <div className="flex justify-between px-52 py-16">
        {renderCurrentMenu()}
      </div>
      <QueueBox />
    </div >
  )
}

export default withQueue(withMenuContext(LobbyScreen))