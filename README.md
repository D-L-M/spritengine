# spritengine

[![GoDoc](https://godoc.org/github.com/D-L-M/spritengine?status.svg)](https://godoc.org/github.com/D-L-M/spritengine)

spritengine is a simple game engine written in Golang that provides the tools to easily build an 80s style 2D side-scrolling platformer.

## Sprites

Sprites are 16x16 pixels in size, represented by a group of 32 32-bit integers. They can contain up to 16 different RGBA colours as defined by a colour palette.

Sprites can be combined into groups to form larger sprites, and both individual sprites and groups can be easily drawn onto an image canvas.

A helper function is included to generate ready-to-use Golang sprite files from PNGs.

## Game Objects

Game objects wrap sprites and add extra properties that the game can interact with, such as mass and velocity.

## Levels

Levels represent a traditional game 'level' and contain all of the game objects required for that level. Under the hood the game delegates responsibility for repainting the canvas to the level.

## Game

The game object contains all of the levels and is responsible for creating a window and capturing user input.