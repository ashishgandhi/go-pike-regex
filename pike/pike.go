/*
Copyright 2013 Ashish Gandhi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// An implementation of Rob Pike's regular expression matcher that
// handles the following constructs:
//
//    c    matches any literal character c
//    .    matches any single character
//    ^    matches the beginning of the input string
//    $    matches the end of the input string
//    *    matches zero or more occurrences of the previous character
//
// http://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html
package pike

import "errors"

var invalidRegexErr = errors.New("Invalid regex")

// Search for regexp anywhere in text
func Match(regexStr string, textStr string) (result bool, err error) {
	regex := []rune(regexStr)
	text := []rune(textStr)
	if len(regex) > 0 && regex[0] == '^' {
		return matchHere(regex[1:], text)
	}
	if len(text) == 0 {
		return matchHere(regex, text)
	}
	for i, _ := range text {
		r, e := matchHere(regex, text[i:])
		if r || e != nil {
			return r, e
		}
	}
	return result, err
}

// Search for regexp at beginning of text
func matchHere(regex []rune, text []rune) (result bool, err error) {
	if len(regex) == 0 {
		return true, err
	}
	if regex[0] == '*' {
		return result, invalidRegexErr
	}
	if regex[0] == '$' {
		if len(regex) > 1 {
			return result, invalidRegexErr
		}
		return len(text) == 0, err
	}
	if len(regex) > 1 && regex[1] == '*' {
		return matchStar(regex[0], regex[2:], text)
	}
	if regex[0] == '.' || regex[0] == text[0] {
		return matchHere(regex[1:], text[1:])
	}
	return result, err
}

// Search for c*regexp at beginning of text
func matchStar(c rune, regex []rune, text []rune) (result bool, err error) {
	if len(text) == 0 {
		return matchHere(regex, text)
	}
	for i, tc := range text {
		r, e := matchHere(regex, text[i:])
		if r || e != nil {
			return r, e
		}
		// Important to not check this before the first loop as
		// c* matches zero or more c
		if tc != c && c != '.' {
			break
		}
	}
	return result, err
}
