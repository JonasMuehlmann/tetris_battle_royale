import { useContext } from "react";
import { Screen, ScreenContext } from "../contexts/screen-context";

const styles = {
  menuItem: 'text-2xl font-semibold text-left w-1/2 pl-14',
  logOutButton: 'text-md opacity-50 text-left'
}

const Menu = () => {
  const { navigate } = useContext(ScreenContext)

  const items = [
    { text: 'Matchfinder', onClick: e => navigate(Screen.Queue) },
    { text: 'Profile', onClick: e => navigate(Screen.Profile) },
    { text: 'Settings', onClick: e => navigate(Screen.Settings) },
  ]

  return (
    <div className="w-full h-full flex p-24">
      <div className="h-full flex flex-col">
        <h1 className="text-4xl">
          Tetris Battle Royale
        </h1>
        <h3 className="text-xl opacity-50">
          Main Menu
        </h3>
        <ul className="flex flex-col">
          {
            items.map((t, i) => (
              <li
                key={i}
                onClick={t.onClick}
                className={styles.menuItem}>
                {t.text}
              </li>
            ))
          }
        </ul>
        <button
          onClick={e => navigate(Screen.LogIn)}
          className={styles.logOutButton}>
          Log out
        </button>
      </div>
    </div>
  )
}

export default Menu;