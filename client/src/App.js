import { BackgroundCanvas } from './components/background_canvas';
import { withScreenContext } from './contexts/screen-context';
import MainScreen from './screens/main_screen';

function App() {
  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      <BackgroundCanvas />
      <MainScreen />
    </div>
  );
}

export default withScreenContext(App);
