package fast_strings

type FastString struct {
	s []byte
	l, r int
}

func NewFastString() *FastString {
	return &FastString{
		s: make([]byte, 0),
		l: 0, r: 0,
	}
}

func (fs *FastString) Concat(s *string) {
	var pointer int = 0
	for ;fs.r < len(fs.s) && pointer < len(*s); pointer++ { // erased characters are replaced with new characters.
		fs.s[fs.r] = (*s)[pointer]
		fs.r++
	}

	for ; pointer < len(*s); pointer++ {
		fs.s = append(fs.s, (*s)[pointer])
		fs.r++
	}

	fs.Trim()
}

func (fs *FastString) ConcatFastString(s *FastString) {
	var pointer int = s.l
	for ;fs.r < len(fs.s) && pointer < s.r; pointer++ { // erased characters are replaced with new characters.
		fs.s[fs.r] = (*s).s[pointer]
		fs.r++
	}

	for ; pointer < s.r; pointer++ {
		fs.s = append(fs.s, (*s).s[pointer])
		fs.r++
	}

	fs.Trim()
}

// Trims only spaces
func (fs *FastString) Trim() {
	space := byte(' ')

	for; fs.l < fs.r; fs.l++ {
		if fs.s[fs.l] != space{
			 break
		}
	}

	for; fs.r > fs.l; fs.r-- {
		if fs.s[fs.r - 1] != space{
			break
		}
	}
}

// Cuts prefix 's' only if there is a complete prefix, and returns
// whether said prefix was found.
func (fs *FastString) CutPrefix(s *string) bool {
	pointer := 0
	for; pointer < len(*s) && fs.l + pointer < fs.r; pointer++ {
		if (*s)[pointer] != fs.s[fs.l + pointer] {
			break
		}
	}

	if pointer == len(*s){
		fs.l += pointer
		fs.Trim()
		return true
	}

	return false
}

func (fs *FastString) Size() int {
	return fs.r - fs.l
}

func (fs *FastString) GetString() *string {
	var s []byte = make([]byte, fs.r - fs.l)

	pointer := fs.l
	for; pointer < fs.r; pointer++ {
		s[pointer - fs.l] = fs.s[pointer]
	}

	ans := string(s)
	return &ans
}