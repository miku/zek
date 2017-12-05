package misc

import "encoding/xml"

// HugsErrors was generated 2017-12-06 00:02:23 by tir on apollo.
type HugsErrors struct {
	XMLName xml.Name `xml:"hugsErrors"`
	Text    string   `xml:",chardata"`
	Section []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		Item []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Keywords string `xml:"keywords,attr"`
			BadCode  struct {
				Text string `xml:",chardata"`
				Code struct {
					Text string `xml:",chardata"` // 3 'div' 4, data BTree a = EmptyBTree | Node a (BTree a) (BTree a), card :: BTree a -> Integer, card ...
					Nl   []struct {
						Text string `xml:",chardata"`
					} `xml:"nl"`
					Highlight []struct {
						Text string `xml:",chardata"` // ), maxFour, bool -> bool -> bool, 2b, n, f (x:xs) = 6, div, ==, a b, b a
					} `xml:"highlight"`
					Tab []struct {
						Text string `xml:",chardata"`
					} `xml:"tab"`
					Space []struct {
						Text string `xml:",chardata"`
					} `xml:"space"`
				} `xml:"code"`
			} `xml:"badCode"`
			Error struct {
				Text string `xml:",chardata"` // ERROR: Improperly terminated character constant, ERROR: Equations give different arities for "card",...
				Nl   []struct {
					Text string `xml:",chardata"`
				} `xml:"nl"`
				Tab []struct {
					Text string `xml:",chardata"`
				} `xml:"tab"`
				Space []struct {
					Text string `xml:",chardata"`
				} `xml:"space"`
				Highlight struct {
					Text string `xml:",chardata"` // test [True]
				} `xml:"highlight"`
			} `xml:"error"`
			Solution struct {
				Chardata string `xml:",chardata"`
				Text     []struct {
					Text   string `xml:",chardata"` // The problem here is the use of the wrong sort of quotes.  					  To turn a function, which is writte...
					Inline []struct {
						Text string `xml:",chardata"` // maxFour a b c d., data, Fun, double, Int -> Int, double (-4), abs, -2, 2, abs
					} `xml:"inline"`
					Nl []struct {
						Text string `xml:",chardata"`
					} `xml:"nl"`
					Tab []struct {
						Text string `xml:",chardata"`
					} `xml:"tab"`
				} `xml:"text"`
				Code []struct {
					Text      string `xml:",chardata"` // card (Node x lt rt, = (height lt) + (height rt) + 1, exOr :: Bool -> Bool -> Bool, >, fun x, >, fun ...
					Highlight struct {
						Text string `xml:",chardata"` // ), [(a,b)]
					} `xml:"highlight"`
					Tab []struct {
						Text string `xml:",chardata"`
					} `xml:"tab"`
					Nl []struct {
						Text string `xml:",chardata"`
					} `xml:"nl"`
				} `xml:"code"`
				Error []struct {
					Text string `xml:",chardata"` // ERROR "text.lhs" (line 1): Illegal Haskell 98 class constraint in inferred type, *** Expression : fu...
					Nl   []struct {
						Text string `xml:",chardata"`
					} `xml:"nl"`
					Tab []struct {
						Text string `xml:",chardata"`
					} `xml:"tab"`
					Space []struct {
						Text string `xml:",chardata"`
					} `xml:"space"`
				} `xml:"error"`
			} `xml:"solution"`
		} `xml:"item"`
	} `xml:"section"`
}
