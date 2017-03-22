# spiral
Playing with spirals and images in Go

Some idle experiments to teach me the basics of Go.

`write` traces a spiral path through an image, outputting ASCII `0` or `1` depending on the colour of the pixel at that point.

`read` draws a new image by reading those ASCII chars and setting pixels along a spiral path.

If you set the `width` and `period` parameters the same, you get a recognisable approximation of the input image. Otherwise
you get warped versions.
