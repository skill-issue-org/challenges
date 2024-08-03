package vars

import "unsafe"

func EAX() uint8 {
	return uint8(unsafe.Sizeof(true))
}

// Flag is the obfuscated of "n1ng$etasiyun₿alapan" string (to avoid `strings`, lol).
var Flag = []byte{
	((((EAX()<<EAX()|EAX())<<EAX()<<EAX()|EAX())<<EAX()^EAX())<<EAX() ^ EAX()) << EAX(),
	((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()<<EAX() | EAX()),
	((((EAX()<<EAX()|EAX())<<EAX()<<EAX()|EAX())<<EAX()^EAX())<<EAX() ^ EAX()) << EAX(),
	((((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()|EAX())<<EAX()^EAX())<<EAX() ^ EAX()),
	(EAX()<<EAX()<<EAX()<<EAX() | EAX()) << EAX() << EAX(),
	(((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()|EAX())<<EAX()<<EAX() ^ EAX()),
	(((EAX()<<EAX()|EAX())<<EAX()|EAX())<<EAX()<<EAX() ^ EAX()) << EAX() << EAX(),
	((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() | EAX()),
	((((EAX()<<EAX()|EAX())<<EAX()|EAX())<<EAX()<<EAX()<<EAX()^EAX())<<EAX() ^ EAX()),
	(((EAX()<<EAX()|EAX())<<EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX() ^ EAX()),
	((((EAX()<<EAX()|EAX())<<EAX()|EAX())<<EAX()^EAX())<<EAX()<<EAX()<<EAX() ^ EAX()),
	((((EAX()<<EAX()|EAX())<<EAX()|EAX())<<EAX()<<EAX()^EAX())<<EAX()<<EAX() ^ EAX()),
	((((EAX()<<EAX()|EAX())<<EAX()<<EAX()|EAX())<<EAX()^EAX())<<EAX() ^ EAX()) << EAX(),
	(((EAX()<<EAX()|EAX())<<EAX()|EAX())<<EAX()<<EAX()<<EAX()<<EAX() ^ EAX()) << EAX(),
	(EAX()<<EAX()<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() | EAX()) << EAX(),
	((((((EAX()<<EAX()<<EAX()|EAX())<<EAX()|EAX())<<EAX()^EAX())<<EAX()^EAX())<<EAX()^EAX())<<EAX() | EAX()),
	((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() | EAX()),
	(((EAX()<<EAX()|EAX())<<EAX()<<EAX()|EAX())<<EAX() ^ EAX()) << EAX() << EAX(),
	((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() | EAX()),
	((EAX()<<EAX()|EAX())<<EAX() | EAX()) << EAX() << EAX() << EAX() << EAX(),
	((EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() | EAX()),
	((((EAX()<<EAX()|EAX())<<EAX()<<EAX()|EAX())<<EAX()^EAX())<<EAX() ^ EAX()) << EAX(),
}
