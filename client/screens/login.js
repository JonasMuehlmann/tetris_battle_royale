const styles = {
  container: 'w-full h-full flex flex-col items-center justify-center',
  header: 'flex flex-col gap-2',
  form: 'flex flex-col gap-2',
}

const LogIn = () => {
  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h2>
          Tetris Battle Royale
        </h2>
      </div>
      <form className={styles.form}>
        <label>
          Username
        </label>
        <input
          
          />
      </form>
    </div>
  )
}

export default LogIn