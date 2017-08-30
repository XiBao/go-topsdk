package security

import "fmt"

const (
	OP_ENCRYPT uint8 = 1
	OP_DECRYPT uint8 = 2
	OP_SEARCH  uint8 = 3
)

const (
	SEP_CHAR_PHONE  rune = '$'
	SEP_CHAR_NICK   rune = '~'
	SEP_CHAR_NORMAL rune = rune(1)
)

const (
	CRYPT_TYPE_NICK          = "nick"
	CRYPT_TYPE_SIMPLE        = "simple"
	CRYPT_TYPE_RECEIVER_NAME = "receiver_name"
	CRYPT_TYPE_SEARCH        = "search"
	CRYPT_TYPE_NORMAL        = "normal"
	CRYPT_TYPE_PHONE         = "phone"
)

var sep_char_map = map[string]rune{
	CRYPT_TYPE_NICK:          SEP_CHAR_NICK,
	CRYPT_TYPE_SIMPLE:        SEP_CHAR_NICK,
	CRYPT_TYPE_RECEIVER_NAME: SEP_CHAR_NICK,
	CRYPT_TYPE_SEARCH:        SEP_CHAR_NICK,
	CRYPT_TYPE_NORMAL:        SEP_CHAR_NORMAL,
	CRYPT_TYPE_PHONE:         SEP_CHAR_PHONE,
}

func get_sep_by_type(cryptType string) (rune, error) {
	sep, ok := sep_char_map[cryptType]
	if !ok {
		return 0, fmt.Errorf("Unknown type: %s", cryptType)
	}
	return sep, nil
}
