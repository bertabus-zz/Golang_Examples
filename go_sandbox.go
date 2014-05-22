package go_sandbox

//simple helloworld
func Helloworld() {
}

//python-numpy like arange function
//from http://stackoverflow.com/questions/19906544/
func Arange(start, stop, step float64) []float64 {
    N := int(math.Ceil((stop - start) / step))
    rnge := make([]float64, N)
    for x := range rnge {
        rnge[x] = start + step*float64(x)
    }
    return rnge
}
