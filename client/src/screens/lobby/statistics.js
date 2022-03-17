const DummyStatistics = {
  games: [
    {
      label: 'Total Games',
      value: 0,
    },
    {
      label: 'Total Wins',
      value: 0,
    },
    {
      label: 'Winrate (Top 10)',
      value: '0%',
    },
  ],
  wins: [{
    label: 'Top 10',
    value: 0,
  },
  {
    label: 'Top 5',
    value: 0,
  },
  {
    label: 'Top 3',
    value: 0,
  },
  {
    label: 'Top 1',
    value: 0,
  },
  ],
  player: [
    {
      label: 'MMR',
      value: 0,
    },
    {
      label: 'Score',
      value: 0,
    },
  ],
}

const Statistics = () => {
  const generateWins = data => (
    <div className="flex justify-between w-full gap-2">
      {DummyStatistics.wins.map((s, i) => (
        <div
          key={i}
          className="flex flex-col gap-2 w-[180px]">
          <p className="text-3xl">
            {s.label}
          </p>
          <h3 className="text-4xl">
            {s.value}
          </h3>
        </div>
      ))}
    </div>
  )

  const generateGamesRecord = data => (
    <div className="flex w-full gap-2 w-full justify-between">
      {DummyStatistics.games.map((s, i) => (
        <div
          key={i}
          className="flex flex-col gap-2 w-[380px]">
          <p className="text-4xl">
            {s.label}
          </p>
          <h3 className="text-5xl">
            {s.value}
          </h3>
        </div>
      ))}
    </div>
  )

  const generatePlayerRecord = data => (
    <div className="flex justify-between w-full gap-2">
      {DummyStatistics.player.map((s, i) => (
        <div
          key={i}
          className="flex flex-col gap-2 w-1/2">
          <p className="text-4xl">
            {s.label}
          </p>
          <h3 className={`text-5xl text-blue-${100 * i + 100}`}>
            {s.value}
          </h3>
        </div>
      ))}
    </div>
  )

  return (
    <div className="flex flex-col gap-24 text-white text-center bangers opacity-60">
      {generateGamesRecord()}
      {generateWins()}
      {generatePlayerRecord()}
    </div>
  )
}

export default Statistics