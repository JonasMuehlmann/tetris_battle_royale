import { BackgroundCanvas } from './components/background_canvas';
import Loader from './components/loader';
import { withDialogContext } from './contexts/dialog-context';
import MainScreen from './screens/main_screen';

function App() {
  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      <BackgroundCanvas />
      <MainScreen />
      <Loader />
    </div>
  );
}

export default withDialogContext(App);
