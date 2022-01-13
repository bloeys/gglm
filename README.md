# gglm

Fast Go OpenGL/Graphics focused Mathematics library inspired by the c++ library [glm](https://github.com/g-truc/glm).

## Notes

You can check compiler inlining decisions using `go run -gcflags "-m" .`. Some functions look a bit weird compared to similar ones
because we are trying to reduce function complexity so the compiler inlines.
