import { useContext } from "react"
import { Screen, ScreenContext, withScreenContext } from "../contexts/screen-context"
import ErrorScreen from "./404";

const MainScreen = (props) => {
  const { currentScreen, navigate } = useContext(ScreenContext)

  return (
    currentScreen.component ?
      (
        <currentScreen.component />
      ) :
      (
        <ErrorScreen
          onNavigate={() => navigate(Screen.LogIn)}
        />
      )
  );
}

export default withScreenContext(MainScreen)