import { useMenu, MenuItem } from "../contexts/menu-context"

const Menu = () => {
  const {
    currentMenu,
    setCurrentMenu
  } = useMenu()

  return (
    <div
      className="flex w-full h-[200px] px-20 items-center justify-center">
      <ul className="flex items-center text-white gap-4 bangers">
        {
          [MenuItem.Matchfinder, MenuItem.Statistics]
            .map((t, i) => (
              <li
                key={i}
                onClick={() => setCurrentMenu(t)}
                className={`
                  cursor-pointer transition-all w-[260px] text-center
                  ${currentMenu?.text === t.text ?
                    'shadow green-grad-text text-5xl' :
                    'text-4xl white-clip-text'}
              `}>
                {t.text}
              </li>
            ))
        }
      </ul>
      <h1 className={`shadow text-center opacity-80 mx-8`}>
        <p className="text-4xl font-bold tetris-text tetris-shadow">
          Tetris Battle Royale
        </p>
        <p className="text-lg text-white opacity-60 border-t border-zinc-600 pt-1 mt-3 mx-36 truncate">
          Public Lobby
        </p>
      </h1>
      <ul className="flex items-center gap-4 text-white bangers">
        {
          [MenuItem.PlayerProfile, MenuItem.PlayerSettings]
            .map((t, i) => (
              <li
                key={i}
                onClick={() => setCurrentMenu(t)}
                className={`text-4xl
            ${currentMenu?.text === t.text ?
                    'shadow text-5xl green-grad-text' :
                    'white-clip-text'}
            cursor-pointer transition-all w-[260px] text-center`}>
                {t.text}
              </li>
            ))
        }
      </ul>
    </div>
  )
}

export default Menu