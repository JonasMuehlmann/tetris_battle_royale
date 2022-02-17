import { random } from "../../helpers/utils";

export default class Block {
  velocity = { x: 0, y: 0 };
  y = 0;
  x = 0;
  size = 36;
  color = 'rgba(0, 0, 0, 0)';
  shadow = 'rgba(0, 0, 0, 0)';

  isLoaded = false;
  isSpawned = false;

  constructor(
    drawWidth,
    drawHeight,
    context,
    source,
    size
  ) {
    this.drawHeight = drawHeight
    this.drawWidth = drawWidth
    this.context = context
    this.source = source
    this.size = size

    this.image = new Image()
    this.image.onload = () => {
      this.isLoaded = true
    }
    this.image.src = source
  }

  update() {
    this.x -= this.velocity.x;
    this.y -= this.velocity.y;
    this._draw();
  }

  spawn() {
    if (this.isSpawned) {
      this.respawn();
      return;
    }

    this.isSpawned = true;
    this.x = random(-this.drawWidth, this.drawWidth);
    this.y = random(-this.drawHeight, this.drawHeight);
  }

  respawn() {
    if (!this.isSpawned) {
      console.warn("Star cannot be respawned if not spawned");
      return;
    }

    this.x = this.drawWidth;
    this.y = random(-this.drawHeight, this.drawHeight);
  }

  _draw() {
    if (!this.isLoaded) return;
    this.context.drawImage(this.image, this.x, this.y, this.size, this.size);
  }

  _getRandomStyle() {
  }
}