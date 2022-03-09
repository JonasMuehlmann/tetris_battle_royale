import { Screen, useScreens } from "../../contexts/screen-context"

const OPTIONS = [
  {
    text: 'Log Out',
    description: 'Click to navigate to LogIn View',
  },
]

const PlayerSettings = () => {
  const {
    navigate
  } = useScreens()

  return (
    <ul className={`flex flex-col 
      w-full gap-8 justify-center items-center
      text-center text-white bangers`}>
      {
        OPTIONS.map((t, i) => (
          <li
            key={i}
            onClick={() => navigate(Screen.LogIn)}
            className={`
              cursor-pointer transition-all w-[480px]
              opacity-30 hover:opacity-100 hover:scale-110`}>
            <p className="text-7xl">
              {t.text}
            </p>
            <span className="text-2xl text-gray-200">
              {t.description}
            </span>
          </li>
        ))
      }
    </ul>
  )
}

export default PlayerSettings