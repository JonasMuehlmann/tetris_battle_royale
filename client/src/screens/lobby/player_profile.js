const OPTIONS = Object.freeze([
  {
    text: 'Profile',
    description: 'In Development',
  },
])

const PlayerProfile = () => {
  return (
    <ul className={`flex flex-col w-full h-full
      justify-center items-center gap-8 
      text-white bangers text-center`}>
      {
        OPTIONS.map((t, i) => (
          <li
            key={i}
            onClick={t.onClick}
            className={`
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
  )
}

export default PlayerProfile