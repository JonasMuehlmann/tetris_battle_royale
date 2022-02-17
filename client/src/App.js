import { useContext } from 'react'
import { ScreenContext, withScreenContext } from './contexts/screen-context'

function App() {
  const { currentScreen } = useContext(ScreenContext)

  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      {<currentScreen.component />}
    </div>
  );
}

export default withScreenContext(App);
