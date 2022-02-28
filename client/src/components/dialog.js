import { useDialog } from "../contexts/dialog-context"

const Dialog = () => {
  const { component } = useDialog()

  return (
    component.isDialogVisible ? (
      <div
        className="w-screen h-screen bg-black bg-opacity-20">
        <div
          className="w-1/3 h-1/3 flex flex-col border rounded-xl bg-white">
          <div
            className="w-full flex flex-col text-center">
            <h2
              className="text-3xl font-semibold py-2">
              {component.model?.title}
            </h2>
            <p>
              {`Tetris Battle Royale`}
            </p>
          </div>
          <div
            className="w-full flex flex-col text-center">
            <p>
              {component.model?.content}
            </p>
          </div>
        </div>
      </div>
    ) : (
      <></>
    )
  )
}

export default Dialog