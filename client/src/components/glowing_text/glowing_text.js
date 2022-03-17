import styles from './glowing_text.module.css'

const GlowingText = ({
  children,
  className,
  glow = true,
}) => (
  <p
    className={
      `${styles.text}` +
      `${glow ? ` ${styles.glowing}` : ''}` +
      ` ${className}`}>
    {children}
  </p>
)

export default GlowingText