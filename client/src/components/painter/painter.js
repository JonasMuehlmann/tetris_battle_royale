import Star from "./star";

const BG_RADIANS = 0.0001;
const BG_ALPHA = 0.98;
const BG_ROTATE_RADIANS = 0.0005;

export class Painter {
  backgroundRotateRadians = BG_RADIANS;
  backgroundAlpha = BG_ALPHA;
  rotateValue = BG_ROTATE_RADIANS;
  speed = 0;
  rotating = true;
  resetting = false;
  stars = [];
  drawWidth = 0;
  drawHeight = 0;

  constructor(
    context,
    canvasWidth,
    canvasHeight,
    starsCount = 100
  ) {
    this.stars = [];
    this.context = context;
    this.starsCount = starsCount
    this.canvasWidth = canvasWidth
    this.canvasHeight = canvasHeight
    this.drawWidth = (canvasWidth + 600) / 2;
    this.drawHeight = (canvasHeight + 1000) / 2;

    for (let i = 0; i < this.starsCount; i++) {
      this.stars.push(new Star(this.drawWidth, this.drawHeight, this.context));
    }

    this.stars.forEach(star => star.spawn());
  }

  update() {
    // clears canvas
    this.context.fillStyle = `rgba(20, 20, 20, ${this.backgroundAlpha})`;
    this.context.fillRect(0, 0, this.canvasWidth, this.canvasWidth);
    // rotates canvas
    this.context.save();
    // sets the pivot point
    this.context.translate(this.canvasWidth / 2, this.canvasHeight / 2);
    // takes radians (360° = 2 * Math.PI)
    this.context.rotate(this.backgroundRotateRadians);
    if (this.rotating && this.backgroundRotateRadians > 0) {
      this.backgroundRotateRadians += this.rotateValue;
    }
    if (this.backgroundRotateRadians >= 2 * Math.PI) {
      this.backgroundRotateRadians = 0;
    }
    /* 
     * beschleunigt sich, wenn sich zurücksetzt
     */
    if (this.resetting) {
      this.rotateValue *= 1.025;
      if (this.backgroundRotateRadians <= 0) {
        this.rotateValue = 0;
        this.backgroundRotateRadians = 0;
        this.rotating = false;
        this.resetting = false;
      }
    }
    /*
     * zeichnet Sterne
     */
    this.stars.forEach((star) => {
      if (star.x < -this.drawWidth) {
        star.respawn();
      } else {
        star.update();
      }
    });

    this.context.restore();
  }

  rotate() {
    if (this.rotating) {
      console.info("Canvas is in rotation");
      return;
    }

    this.rotateValue = BG_ROTATE_RADIANS;
    this.backgroundRotateRadians = BG_RADIANS;
    this.rotating = true;
  }

  reset() {
    const acceleration = 0.001;

    this.resetting = true;
    this.rotateValue = -(this.rotateValue + acceleration);
  }

  setSpeed(speed) {
    this.speed = speed;
    this.stars.forEach(star => star.velocity.x = this.speed);
  }
}