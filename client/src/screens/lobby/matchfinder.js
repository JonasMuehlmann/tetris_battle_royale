import GlowingText from "../../components/glowing_text/glowing_text"
import { useQueue } from "../../contexts/queue-context"

const OPTIONS = Object.freeze([
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
])

const Matchfinder = () => {
  const {
    requestQueue,
    queueType,
  } = useQueue()

  // #region COMPONENTS

  const MatchModes = () => (
    <>
      {
        OPTIONS.map((t, i) => (
          <li
            key={i}
            onClick={e => {
              if (queueType !== null) return
              requestQueue(t.request)
            }}
            className={`
              text-left cursor-pointer transition-all w-[480px]
              ${t.request === queueType ?
                'opacity-100 pl-10' :
                queueType === null ?
                  'opacity-30 hover:opacity-100 hover:pl-10' :
                  'opacity-10 cursor-not-allowed'}`}>
            <GlowingText
              className={`text-7xl`}
              {...t.request !== queueType && ({ glow: false })}>
              {t.text}
            </GlowingText>
            <span className="text-2xl text-gray-200">
              {t.description}
            </span>
          </li>
        ))
      }
    </>
  )

  // #endregion

  return (
    <ul className="flex flex-col gap-10 text-white bangers">
      <h2 className="text-4xl opacity-60 mb-4">
        Select Mode
      </h2>
      <MatchModes />
    </ul>
  )
}

export default Matchfinder