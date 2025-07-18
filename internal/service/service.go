package service

import (
	"strings"
)

func isMorse(data string) bool {
	runes := []rune(data)
	for _, c := range runes {
		if c != '.' && c != '-' && c != ' ' {
			return false
		}
	}
	return true
}

func Convert(data string) string {
	if isMorse(data) {
		return morseToText(data)
	} else {
		return textToMorse(data)
	}
}

func textToMorse(text string) string {
	morseSet := map[rune]string{
		'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
		'Е': ".", 'Ё': ".", 'Ж': "...-", 'З': "--..", 'И': "..",
		'Й': ".---", 'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.",
		'О': "---", 'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-",
		'У': "..-", 'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.",
		'Ш': "----", 'Щ': "--.-", 'Ъ': "--.--", 'Ы': "-.--", 'Ь': "-..-",
		'Э': "..-..", 'Ю': "..--", 'Я': ".-.-",
		'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
		'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
	}

	result := ""
	text = strings.ToUpper(text)

	for i, c := range text {
		code, ok := morseSet[c]
		if !ok {
			return ""
		}

		result += code
		// пробелы после перевода символа
		if i < len(text)-1 {
			result += " "
		}
	}

	return result
}

func morseToText(text string) string {
	morseSet := map[string]rune{
		".-": 'А', "-...": 'Б', ".--": 'В', "--.": 'Г', "-..": 'Д',
		".": 'Е', "...-": 'Ж', "--..": 'З', "..": 'И', ".---": 'Й',
		"-.-": 'К', ".-..": 'Л', "--": 'М', "-.": 'Н', "---": 'О',
		".--.": 'П', ".-.": 'Р', "...": 'С', "-": 'Т', "..-": 'У',
		"..-.": 'Ф', "....": 'Х', "-.-.": 'Ц', "---.": 'Ч', "----": 'Ш',
		"--.-": 'Щ', "--.--": 'Ъ', "-.--": 'Ы', "-..-": 'Ь', "..-..": 'Э',
		"..--": 'Ю', ".-.-": 'Я',
		"-----": '0', ".----": '1', "..---": '2', "...--": '3', "....-": '4',
		".....": '5', "-....": '6', "--...": '7', "---..": '8', "----.": '9',
	}

	splitText := strings.Split(text, " ")
	result := ""

	for _, code := range splitText {
		if code == "" {
			continue
		}

		char, ok := morseSet[code]
		if !ok {
			return ""
		}

		result += string(char)
	}

	return result
}
