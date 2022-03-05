import { random } from "../../helpers/utils";

export default class Star {
  velocity = { x: 0, y: 0 };
  y = 0;
  x = 0;
  radius = 0;
  color = 'rgba(0, 0, 0, 0)';
  shadow = 'rgba(0, 0, 0, 0)';

  isSpawned = false;

  constructor(
    drawWidth,
    drawHeight,
    context
  ) {
    this.drawHeight = drawHeight
    this.drawWidth = drawWidth
    this.context = context
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
    this._getRandomStyle();
  }

  respawn() {
    if (!this.isSpawned) {
      console.warn("Star cannot be respawned if not spawned");
      return;
    }

    this.x = this.drawWidth;
    this.y = random(-this.drawHeight, this.drawHeight);
    this.velocity.x = this.velocity.x;
    this._getRandomStyle();
  }

  _draw() {
    this.context.beginPath();
    this.context.arc(this.x, this.y, this.radius, 0, Math.PI * 2, false);
    this.context.shadowBlur = 10;
    this.context.shadowColor = this.shadow;
    this.context.fillStyle = this.shadow;
    this.context.fill();
    this.context.closePath();
  }

  _getRandomStyle() {
    const hue = random(70, 270);
    const opacity = random(0.01, 0.8);

    this.radius = random(0.25, 2.85);
    this.color = `rgba(${hue}, 171, 255, ${opacity})`;
    this.shadow = `rgba(${hue}, 171 , 255, 1)`;
  }
}