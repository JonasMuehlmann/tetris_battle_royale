import React, { useEffect, useRef, useState } from 'react'
import { Painter } from './painter/painter';

export const BackgroundCanvas = () => {
  const canvasRef = useRef(null);
  const [painter, setPainter] = useState(null);

  (function animate() {
    requestAnimationFrame(animate);
    if (painter) painter.update();
  })()

  function resize() {
    if (canvasRef.current) {
      canvasRef.current.width = window.innerWidth;
      canvasRef.current.height = window.innerHeight;

      const context = canvasRef.current.getContext('2d');
      const starsCount = 180;
      const blocksCount = 7;

      setPainter(
        new Painter(
          context,
          window.innerWidth,
          window.innerHeight,
          starsCount,
          blocksCount
        ));
    }
  }


  useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
    return () => {
      window.removeEventListener('resize', resize)
    }
  }, []);

  return (
    <canvas className='absolute top-0 left-0 z-0' ref={canvasRef}></canvas>
  )
}