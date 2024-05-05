package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 3 {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n")
		fmt.Print("EX: go run . --output=<fileName.txt> something standard\n")
		return
	}
	input := ""
	fileName := getFileName(args[0])
	skip := true
	if fileName == "" {
		skip = false
		input = args[0]
	}
	if input == "" {
		input = args[1]
	}

	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 128 {
			fmt.Printf("error in input\n")
			return
		}
	}
	word := split(input)
	BANNER := banner(args[len(args)-1])
	if BANNER == "" {
		if len(args) > 2 {
			fmt.Print("the banner should be standard or shadow or thinkertoy\n")
		}
		// if there is no banner or err
		BANNER = "standard.txt"
	}

	fileContent, err := os.ReadFile(BANNER)
	if err != nil {
		fmt.Printf("error in stabdard file\n")
		return
	}
	if fileContent == nil {
		fmt.Print("no Text")
	}
	lettres := getLettres(fileContent)
	if skip {
		output(fileName, lettres, word)
		return
	}
	print(lettres, word)
}

func output(fileName string, lettres [][]string, word []string) {
	output := ""
	bl := false
	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if l != len(word)-1 && bl && word[l+1] != "\n" {
				continue
			}
			output += "\n"
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				output += lettres[word[l][j]-32][i]
			}
			output += "\n"
		}
		bl = true
	}

	file2, err2 := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0o644)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer file2.Close()
	file2.Truncate(0)
	file2.WriteString(output)
	os.Exit(0)
}

func print(lettres [][]string, word []string) {
	bl := false
	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if l == len(word)-1 {
				fmt.Print("\n")
				continue
			}
			if bl && word[l+1] != "\n" {
				continue
			}
			fmt.Printf("\n")
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				fmt.Printf(lettres[word[l][j]-32][i])
			}
			fmt.Print("\n")
		}
		bl = true
	}
}

func banner(s string) string {
	BANNER := ""
	if s == "standard" {
		BANNER = "standard.txt"
	} else if s == "thinkertoy" {
		BANNER = "thinkertoy.txt"
	} else if s == "shadow" {
		BANNER = "shadow.txt"
	}
	return BANNER
}

func getFileName(s string) string {
	fileName := ""
	commande := ""
	exten := ""
	bl := true
	for i := 0; i < len(s); i++ {

		if bl {
			commande += string(s[i])
		}

		if !bl {
			fileName += string(s[i])
		}
		if s[i] == '=' {
			bl = false
		}
		if len(s)-5 < i {
			exten += string(s[i])
		}
	}
	if exten != ".txt" {
		fileName = ""
	}
	if len(fileName) < 5 {
		fileName = ""
	}
	if commande=="--output=" && fileName==""{
		fmt.Print("error in commande\n")
		os.Exit(0)
	}
	if fileName != "" && commande != "--output=" {
		fmt.Print("error in commande\n")
		os.Exit(0)
	}
	return fileName
}

func split(str string) []string {
	word := ""
	splitedword := []string{}
	skip := false
	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			skip = true
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	if word != "" {
		splitedword = append(splitedword, word)
	}

	return splitedword
}

func getLettres(fileContent []byte) [][]string {
	Content := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 13 {
			Content = append(Content, fileContent[i])
		}
	}
	fileContent = Content
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if i != len(fileContent)-1 && fileContent[i] == '\n' && fileContent[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if fileContent[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, fileContent[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}
