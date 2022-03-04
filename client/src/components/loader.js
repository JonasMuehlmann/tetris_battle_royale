import { useDialog } from "../contexts/dialog-context"
import { GridLoader, BeatLoader, SyncLoader, DotLoader } from 'react-spinners'
import { AnimatePresence, motion } from "framer-motion"

const Loader = () => {
  const { component } = useDialog()

  return (
    <AnimatePresence>
      {
        component.isDialogVisible ? (
          <motion.div
            key={component.isDialogVisible}
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.25 }}
            className={`w-screen h-screen bg-black bg-opacity-40
      z-40 absolute left-0 top-0 flex items-center justify-center`}>
            <div
              className={`w-[280px] h-[240px] flex flex-col justify-center shadow-lg
          relative p-4 rounded-3xl bg-black bg-opacity-80`}>
              <div
                className="w-full flex flex-col items-center text-center text-white">
                <DotLoader
                  color="#19a186"
                />
                <h2
                  className="text-3xl font-semibold py-6 green-grad-text">
                  {component.currentType?.title || 'Loading..'}
                </h2>
              </div>
              <div
                className="w-full flex flex-col text-center">
                <p className="text-gray-200 text-lg">
                  {component.currentType?.content || 'It won\'t take long..'}
                </p>
              </div>
            </div>
          </motion.div>
        ) : (
          <></>
        )
      }
    </AnimatePresence>

  )
}

export default Loader