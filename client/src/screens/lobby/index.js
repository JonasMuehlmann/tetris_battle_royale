import { AnimatePresence, motion } from 'framer-motion'
import { MenuItem, useMenu, withMenuContext } from "../../contexts/menu-context";
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
      transition: { duration: .5, }
    },
    SlideDown: {
      initial: { opacity: 0, y: window.innerHeight },
      animate: { opacity: 1, y: 0 },
      exit: { opacity: 0, y: window.innerHeight },
      transition: { duration: .5 }
    },
    FromCenter: {
      initial: { opacity: 0, scale: 0 },
      animate: { opacity: 1, scale: 1 },
      exit: { opacity: 0, scale: 1.2 },
      transition: { duration: .5 }
    },
  }

  const renderCurrentMenu = () => {
    switch (currentMenu) {
      case MenuItem.Matchfinder:
        return (
          <motion.div
            {...Motions.SlideLeft}
            className='w-full'
            key={currentMenu.text}>
            <Matchfinder />
          </motion.div>
        )
      case MenuItem.Statistics:
        return (
          <motion.div
            {...Motions.FromCenter}
            className='w-full'
            key={currentMenu.text}>
            <Statistics />
          </motion.div>
        )
      case MenuItem.PlayerProfile:
        return (
          <motion.div
            {...Motions.SlideDown}
            className='w-full'
            key={currentMenu.text}>
            <PlayerProfile />
          </motion.div>
        )
      case MenuItem.PlayerSettings:
        return (
          <motion.div
            {...Motions.SlideDown}
            className='w-full h-full'
            key={currentMenu.text}>
            <PlayerSettings />
          </motion.div>
        )
      default:
        return <></>
    }
  }

  return (
    <div className="w-full h-full flex flex-col z-20 relative">
      <Menu />
      <div className="flex justify-between w-full h-full 2xl:px-52 px-28 py-28">
        <AnimatePresence exitBeforeEnter>
          {renderCurrentMenu()}
        </AnimatePresence>
      </div>
      <QueueBox />
    </div >
  )
}

export default withMenuContext(LobbyScreen)