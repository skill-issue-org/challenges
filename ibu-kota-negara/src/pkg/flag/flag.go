package flag

import . "skill-issue.org/ibu-kota-negara/pkg/utils"

// Get is returning "1buK07a-Ne9@ra" as flag
func Get() string {
	return string(
		[]byte{
			((EAX()<<EAX()^EAX())<<EAX()<<EAX()<<EAX()<<EAX() ^ EAX()),
			((EAX()<<EAX()^EAX())<<EAX()<<EAX()<<EAX()<<EAX() ^ EAX()) << EAX(),
			((((EAX()<<EAX()^EAX())<<EAX()^EAX())<<EAX()<<EAX()|EAX())<<EAX()<<EAX() ^ EAX()),
			(((EAX()<<EAX()<<EAX()<<EAX()^EAX())<<EAX()<<EAX()^EAX())<<EAX() | EAX()),
			(EAX()<<EAX() ^ EAX()) << EAX() << EAX() << EAX() << EAX(),
			((((EAX()<<EAX()^EAX())<<EAX()<<EAX()^EAX())<<EAX()|EAX())<<EAX() ^ EAX()),
			((EAX()<<EAX()^EAX())<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() ^ EAX()),
			(((EAX()<<EAX()<<EAX()^EAX())<<EAX()^EAX())<<EAX()<<EAX() | EAX()),
			(((EAX()<<EAX()<<EAX()<<EAX()^EAX())<<EAX()^EAX())<<EAX() | EAX()) << EAX(),
			(((EAX()<<EAX()^EAX())<<EAX()<<EAX()<<EAX()^EAX())<<EAX()<<EAX() | EAX()),
			(((EAX()<<EAX()^EAX())<<EAX()^EAX())<<EAX()<<EAX()<<EAX() | EAX()),
			EAX() << EAX() << EAX() << EAX() << EAX() << EAX() << EAX(),
			(((EAX()<<EAX()^EAX())<<EAX()^EAX())<<EAX()<<EAX()<<EAX() | EAX()) << EAX(),
			((EAX()<<EAX()^EAX())<<EAX()<<EAX()<<EAX()<<EAX()<<EAX() ^ EAX()),
		},
	)
}
