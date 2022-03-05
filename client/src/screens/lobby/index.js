import { useContext, useState } from "react";
import { Screen, ScreenContext, useScreens } from "../../contexts/screen-context";
import { AnimatePresence, motion } from 'framer-motion'
import Menu from "../../components/menu";
import { MenuItem, useMenu, withMenuContext } from "../../contexts/menu-context";
import Matchfinder from "./matchfinder";
import PlayerProfile from "./player_profile";
import PlayerSettings from "./player_settings";
import Statistics from "./statistics";

const LobbyScreen = () => {
  const { navigate } = useScreens()
  const {
    currentMenu,
  } = useMenu()

  const renderCurrentMenu = () => (
    <AnimatePresence exitBeforeEnter>
      {
        currentMenu === MenuItem.Matchfinder &&
        (
          <motion.div
            initial={{ opacity: 0, x: -window.innerWidth }}
            animate={{ opacity: 1, x: 0 }}
            exit={{ opacity: 0, x: -window.innerWidth }}
            transition={{ duration: .75 }}
            key={currentMenu.text}>
            <Matchfinder />
          </motion.div>
        )
      }
      {
        currentMenu === MenuItem.Statistics &&
        (
          <motion.div
            initial={{ opacity: 0, x: -window.innerWidth }}
            animate={{ opacity: 1, x: 0 }}
            exit={{ opacity: 0, x: -window.innerWidth }}
            transition={{ duration: .75 }}
            key={currentMenu.text}>
            <Statistics />
          </motion.div>
        )
      }
      {
        currentMenu === MenuItem.PlayerProfile &&
        (
          <motion.div
            initial={{ opacity: 0, y: window.innerHeight }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: window.innerHeight }}
            transition={{ duration: .75 }}
            key={currentMenu.text}>
            <PlayerProfile />
          </motion.div>
        )
      }
      {
        currentMenu === MenuItem.PlayerSettings &&
        (
          <motion.div
            initial={{ opacity: 0, y: window.innerHeight }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: window.innerHeight }}
            transition={{ duration: .75 }}
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
    </div >
  )
}

export default withMenuContext(LobbyScreen)