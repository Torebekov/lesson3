package atoi

import (
	"errors"
)
func checkString(s string) bool {
	for i,v := range s {

		if i==0{
			if !(v>='1' && v<='9' || v=='-' && len(s)>1 && s[1]!='0'){
				return false
			}

		} else if !(v>='0' && v<='9'){
			return false
		}
	}
	return true
}
func Atoi(s string) (int, error) {
	if len(s)==1 && s[0]=='0'{
		return 0,nil
	}
	if len(s)==0{
		return 0, errors.New("Empty string provided")
	} else if !checkString(s) {
		return 0, errors.New("The string is invalid")
	}
	sign := 1
	ans := 0
	if s[0]=='-'{
		sign*=-1
		s = string(s[1:])
	}
	for _,v := range s{
		ans *= 10
		ans += int(v-'0')
	}
	ans*=sign
	return ans,nil
}