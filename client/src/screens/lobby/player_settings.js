import { Screen, useScreens } from "../../contexts/screen-context"

const OPTIONS = [
  {
    text: 'Log Out',
    description: 'Click to navigate to LogIn View',
  },
]

const keybinds = [
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
    control: 'Drop: Sorf',
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

  /*const bgColor = "green"

  toggelSound = (e) => {
    if (bgColor == "green") {
      bgColor = "red"  
    } else {
      bgColor = "green"
    }
  }*/

  return (
    <div className="absolute h-3/6 w-4/6">
      <div className="absolute h-14 w-14 top-0 right-0 bg-red-800"
           //style={{backgroundColor: bgColor}}
           /*onClick={this.toggelSound}*/></div>

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

const KeybindTable = () => (
  <table className="absolute top-0 left-0">
    <thead className="yellow text-3xl">
      <tr>
        <th className="pr-52 pb-4">Control</th>
        <th className="pb-4">Keybind</th>
      </tr>
    </thead>
    <tbody className="text-white">
      {keybinds.map(({control, keybind}) => (
        <tr className="border border-transparent border-b-zinc-600">
          <td className="py-2">
            {control}
          </td>
          <td className="py-2 text-center">
            <button className="w-8 h-8 rounded border-solid hover:border-2 hover:border-yellow-200">
              {keybind}
            </button>
          </td>
        </tr>
      ))}
    </tbody>
  </table>
)

const KeybindMessage = () => (
  <div>
    <p>Press the desired key</p>
  </div>
)

export default PlayerSettings