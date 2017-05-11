package main

import (
	"strconv"
	"time"

	//math "github.com/chewxy/math32"
	cm "github.com/vova616/chipmunk"
	vect "github.com/vova616/chipmunk/vect"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
	*sf.Sprite
	*cm.Body

	keyUp    sf.KeyCode
	keyDown  sf.KeyCode
	keyLeft  sf.KeyCode
	keyRight sf.KeyCode

	lives        int
	crouch       bool
	invert       bool
	box          map[string]sf.Rectf
	isMoving     bool
	canMoveL     bool
	canMoveR     bool
	isGrounded   bool
	canJump      bool
	isWalking    bool
	isAnim       bool
	animStage    int
	queuedFrames []Frame
	vel          sf.Vector2f
}

var pTextures map[string]sf.Recti

type Frame struct {
	rect sf.Recti
	name string
}

func NewPlayer(id int, x, y float32) *Player {
	p := new(Player)
	ParsePlayerSpritesheet()
	p.lives = 3
	p.crouch = false
	p.isGrounded = false
	box := cm.NewBox(vect.Vect{0, 35}, 40, 70)
	box.SetElasticity(0)
	box.SetFriction(1)
	p.Body = cm.NewBody(1, box.Moment(1))
	p.Body.SetPosition(vect.Vect{vect.Float(x), vect.Float(-y)})
	p.Body.SetAngle(0)
	p.AddShape(box)
	p.canJump = true
	/*p.box = make(map[string]sf.Rectf)
	p.box["feet"] = sf.Rectf{p.GetPosition().X - 20, p.GetPosition().Y - 5, 40, 5}
	p.box["left"] = sf.Rectf{p.GetGlobalBounds().Left, p.GetGlobalBounds().Top + p.GetGlobalBounds().Height/4, 5, p.GetGlobalBounds().Height / 2}
	p.box["right"] = sf.Rectf{p.GetGlobalBounds().Left + p.GetGlobalBounds().Width - 5, p.GetGlobalBounds().Top + p.GetGlobalBounds().Height/4, 5, p.GetGlobalBounds().Height / 2}*/
	p.Sprite = sf.NewSprite(res.images["p"+strconv.Itoa(id)+"_spritesheet.png"])
	if id == 1 {
		p.keyUp = sf.KeyW
		p.keyDown = sf.KeyS
		p.keyLeft = sf.KeyA
		p.keyRight = sf.KeyD
	}
	if id == 3 {
		p.keyUp = sf.KeyUp
		p.keyDown = sf.KeyDown
		p.keyLeft = sf.KeyLeft
		p.keyRight = sf.KeyRight
	}
	p.SetTextureRect(pTextures["p1_stand"])
	p.Sprite.SetPosition(sf.Vector2f{x, y})
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
	p.StartAnimTimer()
	return p
}

func (p *Player) Update(dt float32) {
	if sf.KeyboardIsKeyPressed(p.keyLeft) && !sf.KeyboardIsKeyPressed(p.keyRight) {
		p.Left()
	}
	if sf.KeyboardIsKeyPressed(p.keyRight) && !sf.KeyboardIsKeyPressed(p.keyLeft) {
		p.Right()
	} /* else {
		p.isMoving = false
	}
	if !p.isMoving {
		if p.isGrounded {
			p.vel.X *= playerDecelG
		} else {
			p.vel.X *= playerDecelA
		}
	}
	if p.canJump && !p.crouch && !p.isMoving {
		p.SetTextureRect(pTextures["p1_stand"])
	}

	if p.vel.X < .15 && p.vel.X > -.15 {
		p.vel.X = 0
	}*/
	p.vel.X = clamp(p.vel.X, -playerTopSpeed, playerTopSpeed)
	if sf.KeyboardIsKeyPressed(p.keyDown) && !p.crouch {
		p.Crouch()
	} else if !sf.KeyboardIsKeyPressed(p.keyDown) && p.crouch {
		p.Uncrouch()
	}

	/*p.isGrounded = false
	p.canMoveL = true
	p.canMoveR = true
	for _, t := range level {
		rect := t.GetGlobalBounds()
		c, _ := p.box["feet"].Intersects(rect)
		if c {
			p.isGrounded = true
			v := p.GetPosition()
			p.Move(sf.Vector2f{0, -v.Y + rect.Top})
		}
		c, _ = p.box["left"].Intersects(rect)
		if c {
			v := p.GetPosition()
			prect := p.GetGlobalBounds()
			p.Move(sf.Vector2f{-(v.X - prect.Width/2) + (rect.Left + rect.Width), 0})
			p.canMoveL = false
		}
		c, _ = p.box["right"].Intersects(rect)
		if c {
			v := p.GetPosition()
			prect := p.GetGlobalBounds()
			p.Move(sf.Vector2f{-(v.X + prect.Width/2) + rect.Left, 0})
			p.canMoveR = false
		}
	}

	if p.isGrounded && p.vel.Y != 0 {
		p.vel.Y = 0
	} else if !p.isGrounded {
		p.vel.Y += gravity
	}

	if p.isGrounded && p.isMoving {
		p.Walk()
	} else {
		p.StopAnim("walk")
	}

	if !p.canMoveL && p.vel.X != 0 {
		p.vel.X = 0
	}

	if !p.canMoveR && p.vel.X != 0 {
		p.vel.X = 0
	}
	*/
	p.canJump = (p.Velocity().Y == 0)

	if sf.KeyboardIsKeyPressed(p.keyUp) {
		if p.canJump /*&& p.isGrounded*/ {
			p.Jump()
		}
	} /*
		if !sf.KeyboardIsKeyPressed(p.keyUp) && p.isGrounded {
			p.canJump = true
		}
	*/
	if p.crouch {
		p.SetTextureRect(pTextures["p1_crouch"])
		rect := p.GetGlobalBounds()
		p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
	}

	p.UpdatePos()

	//p.Move(sf.Vector2f{p.vel.X * dt, p.vel.Y * dt})
}

func (p *Player) Move(pos sf.Vector2f) {
	p.Sprite.Move(pos)
	//v := p.GetPosition()
	/*rect := p.GetGlobalBounds()
	p.box["feet"] = sf.Rectf{v.X - 20, v.Y - 5, 40, 5}
	p.box["left"] = sf.Rectf{rect.Left, rect.Top + rect.Height/4, 5, rect.Height / 2}
	p.box["right"] = sf.Rectf{rect.Left + rect.Width - 5, rect.Top + rect.Height/4, 5, rect.Height / 2}*/

}

func (p *Player) UpdatePos() {
	p.Body.SetAngle(0)
	pos := p.Body.Position()
	p.Sprite.SetPosition(sf.Vector2f{float32(pos.X), float32(-pos.Y)})
	//fmt.Println(p.Position(), p.GetPosition())
}

func (p *Player) Crouch() {
	p.StopAnim("walk")
	p.crouch = true
	p.SetTextureRect(pTextures["p1_crouch"])
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
}

func (p *Player) Uncrouch() {
	p.StopAnim("walk")
	p.crouch = false
	p.SetTextureRect(pTextures["p1_stand"])
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
}

func (p *Player) Left() {
	if !p.invert {
		p.invert = true
		p.Scale(sf.Vector2f{-1, 1})
	}

	if !p.crouch {
		p.isMoving = true
		y := float32(p.Velocity().Y)
		p.Body.AddVelocity(-playerAccel, 0)
		p.Body.SetVelocity(clamp(float32(p.Velocity().X), -playerTopSpeed, playerTopSpeed), y)
	} else {
		p.isMoving = false
	}
}

func (p *Player) Right() {
	if p.invert {
		p.invert = false
		p.Scale(sf.Vector2f{-1, 1})
	}

	if !p.crouch {
		p.isMoving = true
		y := float32(p.Velocity().Y)
		p.Body.AddVelocity(playerAccel, 0)
		p.Body.SetVelocity(clamp(float32(p.Velocity().X), -playerTopSpeed, playerTopSpeed), y)
	} else {
		p.isMoving = false
	}
}

func (p *Player) Jump() {
	p.StopAnim("walk")
	//v := p.Velocity()
	p.AddVelocity(0, playerJump)
	p.canJump = false
	p.QueueFrame(pTextures["p1_jump"], "jump")
	p.StopAnim("walk")
}

func (p *Player) Walk() {
	p.QueueFrame(pTextures["p1_walk"+strconv.Itoa(p.animStage)], "walk")
	p.animStage++
	if p.animStage > 10 {
		p.animStage = 0
	}
}

func (p *Player) StartAnimTimer() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 50)
			p.NextFrame()
		}
	}()
}

func (p *Player) NextFrame() {
	if len(p.queuedFrames) > 0 {
		p.SetTextureRect(p.queuedFrames[0].rect)
		if len(p.queuedFrames) > 0 {
			p.queuedFrames = append(p.queuedFrames[1:])
		} else {
			p.queuedFrames = []Frame{}
		}
	}
}

func (p *Player) QueueFrame(rect sf.Recti, name string) {
	p.queuedFrames = append(p.queuedFrames, Frame{rect, name})
}

func (p *Player) StopAnim(name string) {
	for i := 0; i < len(p.queuedFrames); i++ {
		if p.queuedFrames[i].name == name {
			p.queuedFrames = append(p.queuedFrames[:i], p.queuedFrames[i+1:]...)
			i--
		}
	}
}
