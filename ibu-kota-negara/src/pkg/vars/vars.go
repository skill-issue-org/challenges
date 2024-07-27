package vars

import . "skill-issue.org/ibu-kota-negara/pkg/utils"

// Password is the obfuscated of "simcity" string (to avoid `strings`, lol)
var Password = string(
	[]byte{
		((((EAX()<<EAX()^EAX())<<EAX()|EAX())<<EAX()<<EAX()<<EAX()^EAX())<<EAX() ^ EAX()),
		(((EAX()<<EAX()^EAX())<<EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX() ^ EAX()),
		((((EAX()<<EAX()^EAX())<<EAX()<<EAX()|EAX())<<EAX()^EAX())<<EAX()<<EAX() ^ EAX()),
		(((EAX()<<EAX()^EAX())<<EAX()<<EAX()<<EAX()<<EAX()|EAX())<<EAX() ^ EAX()),
		(((EAX()<<EAX()^EAX())<<EAX()<<EAX()|EAX())<<EAX()<<EAX()<<EAX() ^ EAX()),
		(((EAX()<<EAX()^EAX())<<EAX()|EAX())<<EAX()<<EAX() ^ EAX()) << EAX() << EAX(),
		((((EAX()<<EAX()^EAX())<<EAX()|EAX())<<EAX()^EAX())<<EAX()<<EAX()<<EAX() ^ EAX()),
	},
)
