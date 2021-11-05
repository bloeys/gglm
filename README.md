# gglm

An OpenGL focused Go mathematics inspired by the C++ glm (OpenGL Mathematics) library

## Notes

You can check compiler inlining decisions using `go run -gcflags "-m" .`. Some functions look a bit weird compared to similar ones
because we are trying to reduce function complexity so the compiler inlines.
