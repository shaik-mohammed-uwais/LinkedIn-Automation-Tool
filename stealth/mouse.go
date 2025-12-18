package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// MoveMouseHumanLike moves the mouse in small curved steps
func MoveMouseHumanLike(
	page *rod.Page,
	startX, startY float64,
	endX, endY float64,
) {
	steps := rand.Intn(20) + 25 // 25–45 steps

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		// Smoothstep easing
		ease := t * t * (3 - 2*t)

		x := startX + (endX-startX)*ease + rand.Float64()*2
		y := startY + (endY-startY)*ease + rand.Float64()*2

		page.Mouse.MoveTo(proto.Point{
			X: x,
			Y: y,
		})

		time.Sleep(time.Duration(rand.Intn(15)+5) * time.Millisecond)
	}
}

// MoveMouseToElement moves mouse naturally to an element
func MoveMouseToElement(page *rod.Page, el *rod.Element) error {
	box := el.MustShape().Box()

	targetX := box.X + box.Width/2 + rand.Float64()*4
	targetY := box.Y + box.Height/2 + rand.Float64()*4

	startX := rand.Float64() * 800
	startY := rand.Float64() * 600

	MoveMouseHumanLike(page, startX, startY, targetX, targetY)
	return nil
}
