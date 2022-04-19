import GlowingText from '../../components/glowing_text/glowing_text'

const DATA = [
  {
    label: "Total Games",
    points: 0,
  },
  {
    label: "Total Wins",
    points: 0,
  },
  {
    label: "Winrate (Top 10)",
    points: 0,
  },
  {
    label: "Top 10",
    points: 0,
  },
  {
    label: "Top 5",
    points: 0,
  },
  {
    label: "Top 3",
    points: 0,
  },
  {
    label: "Top 1",
    points: 0,
  },
  {
    label: "MMR",
    points: 0,
  },
  {
    label: "Score",
    points: 0,
  },
]

const PlayerProfile = () => {
  return (
    <div className="grid grid-cols-3 grid-rows-3 w-full h-full text-white">
      {
        DATA.map((data, index) => {
          return (
            <div
              key={index}
              className="flex flex-col justify-center items-center">
              <GlowingText
                className="text-3xl">
                {data.label}
              </GlowingText>
              <p
                className="text-3xl">
                {data.points}
              </p>
            </div>
          )
        })
      }
    </div>
  )
}

export default PlayerProfile