import { AnimatePresence, motion } from "framer-motion"
import { MoonLoader } from "react-spinners"
import { useQueue } from "../contexts/queue-context"

const QueueBox = () => {
  const {
    isInQueue,
    setIsInQueue,
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
            className={`absolute bottom-14 left-1/2
            flex flex-col justify-center gap-2 items-center py-4`}>
            <MoonLoader
              size={32}
              color='#19a186'
            />
            <p className="green-grad-text text-xl">
              Waiting for other players... {elapsed}s
            </p>
            <p
              onClick={() => setIsInQueue(false)}
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