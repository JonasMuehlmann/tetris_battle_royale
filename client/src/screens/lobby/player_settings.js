import { Screen, useScreens } from "../../contexts/screen-context"
import { AiFillSound } from "react-icons/ai"
import { useKeybinds } from "../../contexts/keybinds-context"
import { useEffect, useState } from "react"

const OPTIONS = [
  {
    text: 'Log Out',
    description: 'Click to navigate to LogIn View',
  },
]

const keybindsSample = [
  {
    control: 'Move Left',
    keybind: '*',
  },
  {
    control: 'Move Right',
    keybind: '*',
  },
  {
    control: 'Roatate Clockwise',
    keybind: '*',
  },
  {
    control: 'Roatate Counter-Clockwise',
    keybind: '*',
  },
  {
    control: 'Drop: Soft',
    keybind: '*',
  },
  {
    control: 'Drop: Hard',
    keybind: '*',
  },
  {
    control: 'Hold piece',
    keybind: '*',
  },
]

const PlayerSettings = () => {
  const {
    navigate
  } = useScreens()

  const {
    updateKeybind,
    keybinds
  } = useKeybinds()

  const [iconBackground, setIconBackground] = useState("#107896")
  const [selectedKeybind, setSelectedKeybind] = useState(null)

  const toggleSound = (e) => {
    setIconBackground(prev => prev === "#107896" ? "#9a2617" : "#107896")
  }

  const onKeyDown = (e) => {
    console.log(selectedKeybind)
    if (!selectedKeybind) return;


    updateKeybind(selectedKeybind[0], {
      label: selectedKeybind[1].label,
      key: e.keyCode,
      keyName: e.key,
    })
  }

  useEffect(() => {
    document.addEventListener('keydown', onKeyDown)
  }, [])

  const KeybindTable = () => (
    <table className="absolute top-0 left-0">
      <thead className="yellow text-3xl">
        <tr>
          <th className="pr-52 pb-4">Control</th>
          <th className="pb-4">Keybind</th>
        </tr>
      </thead>
      <tbody className="text-white">
        {Object.entries(keybinds)?.map((entry, i) => (
          <tr
            key={i}
            className="border border-transparent border-b-zinc-600">
            <td className="py-2">
              {entry[1].label}
            </td>
            <td
              onClick={e => setSelectedKeybind(entry)}
              className={`py-2 text-center transition-all focus:outline-0
                ${entry[0] === selectedKeybind?.[0] ? 'opacity-80' : 'opacity-20'}`}>
              <button className="h-8 rounded border-solid hover:border-2 hover:border-yellow-200">
                {entry[1].keyName}
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table >
  )

  return (
    <div
      className="w-full h-full relative">
      <AiFillSound
        fill={iconBackground}
        onClick={toggleSound}
        className="absolute h-14 w-14 top-0 right-0 text-white opacity-30 hover:opacity-100"
      />

      <ul className="flex flex-col gap-8 text-white bangers absolute bottom-0 right-0 transition-all hover:right-4">
        {
          OPTIONS.map((t, i) => (
            <li
              key={i}
              onClick={() => navigate(Screen.LogIn)}
              className={`text-left
                cursor-pointer transition-all w-[270px]
                opacity-30 hover:opacity-100`}>
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

      <KeybindTable />
    </div>
  )
}

const KeybindMessage = () => (
  <div>
    <p>Press the desired key</p>
  </div>
)

export default PlayerSettings