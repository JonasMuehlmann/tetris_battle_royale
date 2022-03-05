import { useQueue } from "../../contexts/queue-context"

const OPTIONS = [
  {
    text: 'Single Play',
    description: 'With A.I. Enemies!',
    request: 'single',
  },
  {
    text: 'Battle Royale',
    description: 'Queue with your friends!',
    request: 'multi',
  },
]

const Matchfinder = ({
}) => {
  const {
    setIsInQueue,
  } = useQueue()

  return (
    <ul className="flex flex-col gap-10 text-white bangers">
      <h2 className="text-4xl opacity-60 shadow mb-4">
        Select Mode
      </h2>
      {
        OPTIONS.map((t, i) => (
          <li
            key={i}
            onClick={e => setIsInQueue(true)}
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

export default Matchfinder