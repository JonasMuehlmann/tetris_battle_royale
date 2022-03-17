import { random } from "../../helpers/utils";
import Block from "./block";
import Star from "./star";

const BG_RADIANS = 0.0001;
const BG_ALPHA = 0.86;
const BG_ROTATE_RADIANS = 0.0005;

export class Painter {
  backgroundRotateRadians = BG_RADIANS;
  backgroundAlpha = BG_ALPHA;
  rotateValue = BG_ROTATE_RADIANS;
  speed = 0;
  rotating = false;
  stop = false;
  stars = [];
  drawWidth = 0;
  drawHeight = 0;

  constructor(
    context,
    canvasWidth,
    canvasHeight,
    starsCount = 100,
    blocksCount = 5,
  ) {
    this.stars = []
    this.blocks = []
    this.context = context
    this.starsCount = starsCount
    this.blocksCount = blocksCount
    this.canvasWidth = canvasWidth
    this.canvasHeight = canvasHeight
    this.drawWidth = canvasWidth + 100;
    this.drawHeight = canvasHeight;

    this.init()
  }

  init() {
    for (let i = 0; i < this.starsCount; i++) {
      this.stars.push(new Star(this.drawWidth, this.drawHeight, this.context))
    }

    for (let i = 0; i < this.blocksCount; i++) {
      this.blocks.push(new Block(
        this.drawWidth,
        this.drawHeight,
        this.context,
        `assets/block_${random(1, 3)}.png`,
        random(20, 70)))
    }

    this.stars.forEach(star => star.spawn())
    this.blocks.forEach(block => block.spawn())

    this.setSpeed(3.5);
  }

  update() {
    // clears canvas
    this.context.fillStyle = `rgba(20, 20, 20, ${this.backgroundAlpha})`;
    this.context.fillRect(0, 0, this.drawWidth, this.drawHeight);
    this.context.save()
    /*
     * zeichnet Sterne
     */
    this.stars.forEach((star) => {
      if (star.x < -this.canvasWidth) {
        star.respawn()
      } else {
        star.update()
      }
    })
    this.blocks.forEach(block => {
      if (block.x < -this.canvasWidth) {
        block.respawn()
      } else {
        block.update()
      }
    })

    if (this.stop) {
      const currentSpeed = this.speed
      if (currentSpeed > 0) {
        this.setSpeed(currentSpeed - 0.1)
      } else if (currentSpeed < 0) {
        this.setSpeed(0)
      }
    }

    this.context.restore()
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
    this.stars.forEach(star => {
      let randomized = this.speed * (Math.random() + 0.75)
      randomized = randomized > 15 ? 15 - randomized : randomized
      star.velocity.x = randomized
    });
    this.blocks.forEach(block => {
      let randomized = this.speed * (Math.random() + 0.75)
      randomized = randomized > 15 ? 15 - randomized : randomized
      block.velocity.x = randomized
    });
  }
}