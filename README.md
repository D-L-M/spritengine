# spritengine

spritengine is a simple sprite engine written in Go that uses 32-bit integers to store scanline data (two integers per scanline).

Sprites are 16x16 pixels in size and can contain up to 16 different RGBA colours as defined by a colour palette. Sprites can be combined into groups to form larger sprites, and both individual sprites and groups can be easily drawn onto an image canvas.

## Example

For a basic example, please see the code in the `example/` directory.