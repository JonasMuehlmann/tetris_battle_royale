import { useContext, useState } from "react";
import { Screen, ScreenContext } from "../../contexts/screen-context";


const Menu = () => {
  const { navigate } = useContext(ScreenContext)
  const leftItems = [
    { text: 'Matchfinder', onClick: e => navigate(Screen.LogIn) },
    { text: 'Stats', onClick: e => navigate(Screen.LogIn) },
  ]
  const rightItems = [
    { text: 'Profile', onClick: e => navigate(Screen.LogIn) },
    { text: 'Settings', onClick: e => navigate(Screen.LogIn) },
  ]
  const matchOptions = [
    {
      text: 'Single Play',
      description: 'With A.I. Enemies!',
    },
    {
      text: 'Battle Royale',
      description: 'Queue with your friends!',
    },
  ]
  const [currentMenu, setCurrentMenu] = useState(leftItems[0])

  return (
    <div className="w-full h-full flex flex-col z-20 relative">
      <div className="flex w-full h-[160px] pt-5 px-20 items-center justify-center">
        <ul className="flex text-white gap-4 bangers">
          {
            leftItems.map((t, i) => (
              <li
                key={i}
                onClick={t.onClick}
                className={`text-5xl 
                  ${currentMenu?.text === t.text ? 'shadow green-grad-text' : 'opacity-40 hover:opacity-100'}
                  cursor-pointer transition-all w-[280px] text-center
                  `}>
                {t.text}
              </li>
            ))
          }
        </ul>
        <h1 className={`text-4xl shadow text-center opacity-80 mx-10`}>
          <p className="tetris-text tetris-shadow">
            Tetris Battle Royale
          </p>
          <p className="text-lg text-white opacity-60 border-t border-zinc-600 pt-1 mt-2 mx-36 truncate">
            Public Lobby
          </p>
        </h1>
        <ul className="flex gap-4 text-white bangers">
          {
            rightItems.map((t, i) => (
              <li
                key={i}
                onClick={t.onClick}
                className={`text-5xl
                ${currentMenu?.text === t.text ? 'shadow sway-animation' :
                    'opacity-40 hover:opacity-100'}
                cursor-pointer transition-all w-[280px] text-center`}>
                {t.text}
              </li>
            ))
          }
        </ul>
      </div>
      <div className="flex justify-between px-48 py-16">
        <ul className="flex flex-col gap-8 text-white bangers">
          {
            matchOptions.map((t, i) => (
              <li
                key={i}
                onClick={t.onClick}
                className={`text-left
                  cursor-pointer transition-all w-[480px]
                  opacity-30 hover:opacity-100 hover:pl-10`}>
                <p className="text-7xl tetris-text ">
                  {t.text}
                </p>
                <span className="text-2xl text-gray-200">
                  {t.description}
                </span>
              </li>
            ))
          }
        </ul>
      </div>
    </div>
  )
}

export default Menu;