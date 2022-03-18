import {useState} from 'react'

const OPTIONS = Object.freeze([
  {
    text: 'Profile',
    description: 'In Development',
  },
])

const PlayerProfile = () => {
  const [isOpen, setIsOpen] = useState(false)

  // #region COMPONENTS
  // #endregion
  return (
    <ul className="flex flex-col gap-8 text-white bangers">
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
      <button
        className='text-white opacity-50 transition-all hover:opacity-100 hover:text-red'>
        Open
      </button>
    </ul>
  )
}

export default PlayerProfile