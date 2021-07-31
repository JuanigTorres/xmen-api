package exception

type MatrixError string

func (err MatrixError) Error() string {
return string(err)
}