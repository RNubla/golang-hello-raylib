package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxFrameSpeed = 15
	minFrameSpeed = 1
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] example - texture loading and drawing")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)
	adventurerSheet := rl.LoadTexture("assets/player/adventurer-Sheet.png") // Texture loading width:50px height:37px each frame

	currentFrame := float32(0)

	framesCounter := 0
	framesSpeed := 8 // Number of spritesheet frames shown by second

	frameWidth := float32(adventurerSheet.Width) / 7
	frameHeight := float32(adventurerSheet.Height) / 11

	sourceRect := rl.NewRectangle(0, 0, frameWidth, frameHeight)
	// sourceRect_run := rl.NewRectangle(50,)

	//centers sprite on screen and inrease its size
	destRec := rl.NewRectangle(float32(screenWidth)/2, float32(screenHeight)/2, frameWidth*3, frameHeight*3)
	origin := rl.NewVector2(frameWidth, frameHeight)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		framesCounter++

		if framesCounter >= (60 / framesSpeed) {
			framesCounter = 0
			currentFrame++

			//idle
			// if currentFrame > 3 {
			// 	currentFrame = 0
			// }

			// sourceRect.X = currentFrame * float32(adventurerSheet.Width) / 7

			if rl.IsKeyDown(rl.KeyRight) {
				sourceRectRun := rl.NewRectangle(50, 37, frameWidth, frameHeight)
				// sourceRect := rl.NewRectangle(50, frameHeight, frameWidth, frameHeight)

				if currentFrame > 6 {
					currentFrame = 1
				}

				sourceRectRun.X = currentFrame * float32(adventurerSheet.Width) / 7
				println("KeyRight is pressed down")
			} else if currentFrame > 3 {
				currentFrame = 0
			}
			sourceRect.X = currentFrame * float32(adventurerSheet.Width) / 7
		}
		// println("Current Frame: ", int(currentFrame))

		// if rl.IsKeyDown(rl.KeyRight) {
		// 	framesSpeed++
		// } else if rl.IsKeyDown(rl.KeyLeft) {
		// 	framesSpeed--
		// }

		// if framesSpeed > maxFrameSpeed {
		// 	framesSpeed = maxFrameSpeed
		// } else if framesSpeed < minFrameSpeed {
		// 	framesSpeed = minFrameSpeed
		// }

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(adventurerSheet, 15, 40, rl.White)
		rl.DrawRectangleLines(15, 40, adventurerSheet.Width, adventurerSheet.Height, rl.Lime)
		rl.DrawRectangleLines(15+int32(sourceRect.X), 40+int32(sourceRect.Y), int32(sourceRect.Width), int32(sourceRect.Height), rl.Red)

		rl.DrawTexturePro(adventurerSheet, sourceRect, destRec, origin, 0, rl.White)

		rl.DrawText("(c) texture sprite by Eiden Marsal", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.EndDrawing()
	}

	rl.UnloadTexture(adventurerSheet)

	rl.CloseWindow()
}
