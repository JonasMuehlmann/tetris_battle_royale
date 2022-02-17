const Menu = () => {
  return (
    <div className="w-full h-full flex p-24">
      <div className="h-full flex flex-col">
        <h1 className="text-4xl">
          Tetris Battle Royale
        </h1>
        <h3 className="text-xl opacity-50">
          Main Menu
        </h3>
      </div>
      <ul className="flex flex-col">
        <li>
          Matchfinder
        </li>
        <li>
          Profile
        </li>
        <li>
          Settings
        </li>
      </ul>
      <button className="text-md opacity-50">
        Log out
      </button>
    </div>
  )
}

export default Menu;