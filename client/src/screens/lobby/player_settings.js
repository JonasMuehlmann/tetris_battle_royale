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
    <ul className="flex flex-col gap-8 text-white bangers">
      {
        OPTIONS.map((t, i) => (
          <li
            key={i}
            onClick={() => navigate(Screen.LogIn)}
            className={`text-left
              cursor-pointer transition-all w-[480px]
              opacity-30 hover:opacity-100 hover:pl-10`}>
            <p className="text-7xl yellow">
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