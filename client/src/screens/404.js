const ErrorScreen = ({ onNavigate }) => (
  <div className="w-full h-full flex flex-col items-center justify-center">
    <h2 className="text-4xl">
      Invalid Screen Routing
    </h2>
    <button
      type="button"
      onClick={e => onNavigate()}
      className="border w-96 text-sm mt-2 opacity-60 hover:opacity-100">
      Back
    </button>
  </div>
)

export default ErrorScreen