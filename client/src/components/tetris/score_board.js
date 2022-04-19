import GlowingText from "../glowing_text/glowing_text"

const DATA = [
    {
        label: "score: ", 
        points: 0,
    },
    {
        label: "tetris: ",
        points: 0,
    },
    {
        label: "rank: ",
        points: 0,
    }
]

const ScoreBoard = ({ score }) => {
    return (
        <div className="absolute top-20 left-1/2 transform -translate-x-1/2 -translate-y-1/5 flex gap-2 flex-row">
            {
                DATA.map((data,index) => {
                    return (
                        <div 
                            key={index}
                            className="flex flex-nowrap
                                       flex-row">
                            <GlowingText
                                className=" text-xl
                                            justify-evenly">
                                {data.label}
                            </GlowingText>
                            <p
                                className="text-xl
                                            pl-2 pr-4">
                                    {data.points}
                            </p>
                        </div>
                    )
                } )
            }
        </div>
    )
}

export default ScoreBoard