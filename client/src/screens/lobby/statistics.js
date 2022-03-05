const OPTIONS = [
  {
    text: 'Statistics',
    description: 'In Development',
  },
]


const Statistics = () => {
  return (
    <ul className="flex flex-col gap-8 text-white bangers">
      {
        OPTIONS.map((t, i) => (
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
  )
}

export default Statistics