import { motion } from "framer-motion"
import { FaUserFriends } from 'react-icons/fa'

const FriendsBox = () => {
  return (
    <motion.div
      initial={{ y: 500, opacity: 0, x: '-50%' }}
      animate={{ y: 0, opacity: 1 }}
      transition={{ duration: .75, type: 'spring' }}
      className={`absolute top-40 left-1/2 z-30 
        flex justify-center gap-2`}>
      <button className={`text-white text-xl transition-all
        opacity-20 hover:opacity-60 cursor-pointer
        flex gap-2 py-3 px-4 rounded-lg items-center`}>
        <FaUserFriends
          size={24}
        />
        0
      </button>
    </motion.div>
  )
}

export default FriendsBox