import { AnimatePresence, motion } from "framer-motion"
import { BeatLoader, MoonLoader, PulseLoader, SyncLoader } from "react-spinners"
import { useQueue } from "../contexts/queue-context"

const QueueBox = () => {
  const {
    isInQueue,
    cancelQueue,
    elapsed
  } = useQueue()

  return (
    <AnimatePresence>
      {
        isInQueue && (
          <motion.div
            initial={{ y: 200, x: '-50%' }}
            animate={{ y: 0 }}
            exit={{ y: 200 }}
            transition={{ duration: .75, type: 'spring' }}
            className={`absolute bottom-16 left-1/2
            flex flex-col justify-center items-center py-4`}>
            {
              elapsed > 3 ? (
                <SyncLoader
                  size={10}
                  color='#19a18688'
                />
              ) : (
                <MoonLoader
                  size={32}
                  color='#19a186'
                />
              )
            }
            <p className="green-grad-text text-xl py-2">
              {
                elapsed > 3 ?
                  'Match Found! Navigating..' :
                  `Waiting for other players... ${elapsed}s`
              }
            </p>
            <p
              onClick={() => cancelQueue()}
              className="text-gray-400 opacity-50 text-sm cursor-pointer hover:opacity-80">
              Cancel
            </p>
          </motion.div>
        )
      }
    </AnimatePresence>
  )
}

export default QueueBox