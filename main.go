package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxFrameSpeed = 15
	minFrameSpeed = 1
)

var adventurerSheet rl.Texture2D
var screenWidth int32
var screenHeight int32

var currentFrame float32
var frameCounter int
var frameSpeed int

var sourceRect rl.Rectangle
var runRect rl.Rectangle
var destRec rl.Rectangle
var origin rl.Vector2

func main() {
	screenWidth = int32(800)
	screenHeight = int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Sprite Sheet Animation")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)
	adventurerSheet = rl.LoadTexture("assets/player/adventurer-Sheet.png") // Texture loading width:50px height:37px each frame

	// currentFrame = float32(0)

	// frameCounter = 0
	frameSpeed = 8 // Number of spritesheet frames shown by second

	frameWidth := float32(adventurerSheet.Width) / 7
	frameHeight := float32(adventurerSheet.Height) / 11

	sourceRect = rl.NewRectangle(0, 0, frameWidth, frameHeight)
	runRect := rl.NewRectangle(50, 37, frameWidth, frameHeight)

	//centers sprite on screen and inrease its size
	destRec = rl.NewRectangle(float32(screenWidth)/2, float32(screenHeight)/2, frameWidth*3, frameHeight*3)
	origin = rl.NewVector2(frameWidth, frameHeight)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		frameCounter++
		if rl.IsKeyDown(rl.KeyD) {
			if frameCounter >= (60 / frameSpeed) {
				frameCounter = 0
				currentFrame++
				if currentFrame > 6 {
					currentFrame = 1
				}
			}
			runRect.X = currentFrame * float32(adventurerSheet.Width) / 7
			rl.DrawRectangleLines(15+int32(runRect.X), 40+int32(runRect.Y), int32(runRect.Width), int32(runRect.Height), rl.Red)
			rl.DrawTexturePro(adventurerSheet, runRect, destRec, origin, 0, rl.White)
		}
		if rl.IsKeyUp(rl.KeyD) {
			if frameCounter >= (60 / frameSpeed) {
				frameCounter = 0
				currentFrame++
				if currentFrame > 3 {
					currentFrame = 0
				}
				sourceRect.X = currentFrame * float32(adventurerSheet.Width) / 7
			}
			rl.DrawRectangleLines(15+int32(sourceRect.X), 40+int32(sourceRect.Y), int32(sourceRect.Width), int32(sourceRect.Height), rl.Red)

			rl.DrawTexturePro(adventurerSheet, sourceRect, destRec, origin, 0, rl.White)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(adventurerSheet, 15, 40, rl.White)
		rl.DrawRectangleLines(15, 40, adventurerSheet.Width, adventurerSheet.Height, rl.Lime)

		rl.DrawText("(c) texture sprite by Eiden Marsal", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.EndDrawing()
	}

	rl.UnloadTexture(adventurerSheet)

	rl.CloseWindow()
}
