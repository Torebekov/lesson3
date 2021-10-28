package itoa

func Itoa(i int) (s string) {
	var res string
	if i<0{
		s+="-"
		i*=-1
	}
	for {
		res+=string(byte(i%10+'0'))
		i/=10
		if i==0 {
			break
		}
	}
	for i := range res{
		s+=string(res[len(res)-i-1])
	}
	return

}