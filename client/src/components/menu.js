import { useMenu, MenuItem } from "../contexts/menu-context"
import GlowingText from "./glowing_text/glowing_text"

const Menu = () => {
  const {
    currentMenu,
    setCurrentMenu
  } = useMenu()

  return (
    <div
      className="flex w-full h-[220px] 2xl:px-40 px-24 items-center justify-between relative">
      <ul className="flex items-center text-white gap-4 bangers">
        {
          [MenuItem.Matchfinder, MenuItem.Statistics]
            .map((t, i) => (
              <li
                key={i}
                onClick={() => setCurrentMenu(t)}
                className={`
                  cursor-pointer transition-all text-center w-56
                  ${currentMenu?.text === t.text ?
                    'green-grad-text lg:text-5xl text-3xl' :
                    'lg:text-4xl text-2xl white-clip-text'}
              `}>
                {t.text}
              </li>
            ))
        }
      </ul>
      <h1 className={`absolute left-1/2 top-20 -translate-x-1/2 
        text-center opacity-80`}>
        <GlowingText
          glow={false}
          className={`2xl:text-4xl text-3xl font-bold`}>
          Tetris Battle Royale
        </GlowingText>
        <p className="2xl:text-xl text-lg text-white opacity-60 border-t border-zinc-600 pt-2 mt-3 mx-36 truncate">
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
                className={`text-4xl w-56
            ${currentMenu?.text === t.text ?
                    'lg:text-5xl text-3xl green-grad-text' :
                    'white-clip-text'}
            cursor-pointer transition-all text-center`}>
                {t.text}
              </li>
            ))
        }
      </ul>
    </div>
  )
}

export default Menu