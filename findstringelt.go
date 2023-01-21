package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RemoveIndex(s []string, index int)[]string {
	return append(s[:index], s[index+1:]...)
}

func RemoveDuplicate(arr []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range arr {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func main(){
	str := "1E (hex) files were added. It has been 10 (bin) years. Ready, set, go (up) ! I should stop SHOUTING (low) . Welcome to the Brooklyn bridge (cap) . This is so exciting (up,2). ' Bonjour am exactly ' what ' you ' are saying"
	s := strings.Split(str, " ")
	fmt.Println(s)

	for i :=0 ; i <= int(len(s) -1); i++{
		
		if s[i] == "(hex)" {

			hexadecimal_num := s[i-1]
			decimal_num , err := strconv.ParseInt(hexadecimal_num, 16, 64)

			if err != nil {
				panic(err)
			}

			s[i-1] = strconv.Itoa(int(decimal_num))
			RemoveIndex(s, i)

		} else if s[i]=="(bin)" {

			deximal_num_Str := s[i-1]
			deximal_num, err := strconv.Atoi(deximal_num_Str)

			if err != nil {
				panic(err)
			}

			if deximal_num == 10 {
				s[i-1] = strconv.Itoa(2)
			}
			RemoveIndex(s, i)

		} else if s[i] == "(up)" {

			lowerCase := s[i - 1]
			upperCase := strings.ToUpper(lowerCase)
			s[i - 1] = upperCase
			RemoveIndex(s, i)

		} else if s[i] == "(low)" {

			upperCase := s[i - 1]
			lowerCase := strings.ToLower(upperCase)
			s[i - 1] = lowerCase
      RemoveIndex(s, i)

		} else if s[i] == "(cap)" {

			str := s[i-1]
			cap := strings.ToUpper(str[:1]) + str[1:]
			s[i-1] = cap
			RemoveIndex(s, i)

		} else if s[i] == "(up,2)"{

			lowerCase1 := s[i - 1]
			lowerCase2 := s[i - 2]
			upperCase1 := strings.ToUpper(lowerCase1)
			upperCase2 := strings.ToUpper(lowerCase2)
			s[i - 1] = upperCase1
			s[i - 2] = upperCase2
			RemoveIndex(s, i)

		} else if s[i] =="." || s[i] == "," || s[i] == "!" ||s[i] == "?" || s[i] ==";" {
			ponctuation := s[i]
			s[i-1] = s[i-1]+ponctuation
			RemoveIndex(s, i)
			
		} 
	}
	 Result := strings.Join(RemoveDuplicate(s), " ")
	fmt.Println(Result)
}