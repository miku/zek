package dsrrs

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:repository.ul.hirosak...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2009-04-20T03:00:00Z, 200...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // hdl_10129_156, hdl_10129_...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text     string `xml:",chardata"`
		Iteminfo struct {
			Text string `xml:",chardata"`
			Item struct {
				Text string `xml:",chardata"` // 116, 117, 118, 119, 120, ...
			} `xml:"item"`
			Handle struct {
				Text string `xml:",chardata"` // 10129/157, 10129/158, 101...
			} `xml:"handle"`
			LastModified struct {
				Text string `xml:",chardata"` // 2009-04-20 12:00:00.0, 20...
			} `xml:"lastModified"`
			Submitter struct {
				Text string `xml:",chardata"` // å³æ¸æå ±æå½ å­¦è¡...
			} `xml:"submitter"`
			OwingCollection []struct {
				Text string `xml:",chardata"` // 10129/156, 10129/156, 101...
			} `xml:"owingCollection"`
			MappingCollection struct {
				Text string `xml:",chardata"` // 10129/141, 10129/141, 101...
			} `xml:"mappingCollection"`
		} `xml:"iteminfo"`
		Dc struct {
			Text                  string `xml:",chardata"`
			OaiDc                 string `xml:"oai_dc,attr"`
			Dc                    string `xml:"dc,attr"`
			Xsi                   string `xml:"xsi,attr"`
			SchemaLocation        string `xml:"schemaLocation,attr"`
			ContributorAuthorNull []struct {
				Text string `xml:",chardata"` // æ­£æ, åå½¦, ç®é», ç...
			} `xml:"contributor.author.null"`
			ContributorOtherEn []struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦å»å­¦é¨è§£å...
			} `xml:"contributor.other.en"`
			DateAccessionedNull struct {
				Text string `xml:",chardata"` // 2007-03-06T07:26:34Z, 200...
			} `xml:"date.accessioned.null"`
			DateAvailableNull struct {
				Text string `xml:",chardata"` // 2007-03-06T07:26:34Z, 200...
			} `xml:"date.available.null"`
			DateIssuedNull struct {
				Text string `xml:",chardata"` // 2006, 2006, 2006, 2006, 2...
			} `xml:"date.issued.null"`
			IdentifierURINull struct {
				Text string `xml:",chardata"` // http://hdl.handle.net/101...
			} `xml:"identifier.uri.null"`
			FormatExtentNull []struct {
				Text string `xml:",chardata"` // 30354944 bytes, 57344 byt...
			} `xml:"format.extent.null"`
			FormatMimetypeNull []struct {
				Text string `xml:",chardata"` // application/vnd.ms-powerp...
			} `xml:"format.mimetype.null"`
			LanguageIsoEn struct {
				Text string `xml:",chardata"` // ja, ja, ja, ja, ja, en, j...
			} `xml:"language.iso.en"`
			PublisherNullEn struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦å»å­¦é¨è§£å...
			} `xml:"publisher.null.en"`
			TitleNullEn struct {
				Text string `xml:",chardata"` // è³ã®è§£åå­¦. ç¬¬1é¨,...
			} `xml:"title.null.en"`
			TitleAlternativeEn struct {
				Text string `xml:",chardata"` // Brain Dissection. Part 1,...
			} `xml:"title.alternative.en"`
			TypeNullEn struct {
				Text string `xml:",chardata"` // Learning Object, Learning...
			} `xml:"type.null.en"`
			SubjectNiiEn struct {
				Text string `xml:",chardata"` // åºç¤å»å­¦, åºç¤, åº...
			} `xml:"subject.nii.en"`
			SubjectNdcEn struct {
				Text string `xml:",chardata"` // 491.1, 491.1, 491.1, 491....
			} `xml:"subject.ndc.en"`
			TypeNiiEn struct {
				Text string `xml:",chardata"` // Learning Material, Learni...
			} `xml:"type.nii.en"`
			RightsTextversionEn struct {
				Text string `xml:",chardata"` // author, author, author, a...
			} `xml:"rights.textversion.en"`
			TitleTranscriptionEn struct {
				Text string `xml:",chardata"` // ãã¦ ã ã«ã¤ãã¦ã...
			} `xml:"title.transcription.en"`
			ContributorOtherNull []struct {
				Text string `xml:",chardata"` // éæ£®çç«ä¸­å¤®çé¢æ...
			} `xml:"contributor.other.null"`
			IdentifierCitationEn struct {
				Text string `xml:",chardata"` // å¼åå»å­¦. 55(1), 2003...
			} `xml:"identifier.citation.en"`
			IdentifierISSNNull []struct {
				Text string `xml:",chardata"` // 0439-1721, 0287-0886, 028...
			} `xml:"identifier.issn.null"`
			DescriptionNullEn struct {
				Text string `xml:",chardata"` // å¹³æ17å¹´åº¦ï½å¹³æ18...
			} `xml:"description.null.en"`
			DescriptionAbstractEn struct {
				Text string `xml:",chardata"` // ä¸»ã«30Gy/3åå²/3æ¥é...
			} `xml:"description.abstract.en"`
			SubjectNullEn []struct {
				Text string `xml:",chardata"` // è´ç¥çµè«ç, å®ä½æ...
			} `xml:"subject.null.en"`
			SubjectNiiNull []struct {
				Text string `xml:",chardata"` // åç§ç³»è¨åºå»å­¦, è¾...
			} `xml:"subject.nii.null"`
			IdentifierNiiissnEn struct {
				Text string `xml:",chardata"` // 0439-1721, 0287-0886, 028...
			} `xml:"identifier.niiissn.en"`
			IdentifierNcidEn struct {
				Text string `xml:",chardata"` // AN00211444, AN00150918, A...
			} `xml:"identifier.ncid.en"`
			IdentifierJtitleEn struct {
				Text string `xml:",chardata"` // å¼åå»å­¦, å²å­¦ä¼èª...
			} `xml:"identifier.jtitle.en"`
			IdentifierVolumeEn struct {
				Text string `xml:",chardata"` // 55, 57, 57, 30, 31, 47, 2...
			} `xml:"identifier.volume.en"`
			IdentifierIssueEn struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			} `xml:"identifier.issue.en"`
			IdentifierSpageEn struct {
				Text string `xml:",chardata"` // 7, 1, 2, 3, 15, 21, 38, 5...
			} `xml:"identifier.spage.en"`
			IdentifierEpageEn struct {
				Text string `xml:",chardata"` // 17, 1, 2, 14, 20, 37, 55,...
			} `xml:"identifier.epage.en"`
			TitleTranscriptionNull struct {
				Text string `xml:",chardata"` // ãã¤ã¤ãã­ ãã¦ã·...
			} `xml:"title.transcription.null"`
			SortkeyNullNull struct {
				Text string `xml:",chardata"` // 002, 000, 001, 002, 003, ...
			} `xml:"sortkey.null.null"`
			ContributorAlternativeNull []struct {
				Text string `xml:",chardata"` // Saito, Takeo, Kawaguchi, ...
			} `xml:"contributor.alternative.null"`
			ContributorTranscriptionNull []struct {
				Text string `xml:",chardata"` // ãµã¤ãã¦, ã¿ã±ãª, ...
			} `xml:"contributor.transcription.null"`
			DescriptionNullNull []struct {
				Text string `xml:",chardata"` // å¹³æ12å¹´åº¦ï½15å¹´åº¦...
			} `xml:"description.null.null"`
			IdentifierCitationNull struct {
				Text string `xml:",chardata"` // 19p ; 30cm, å¼åå»å­¦....
			} `xml:"identifier.citation.null"`
			SubjectNdcNull struct {
				Text string `xml:",chardata"` // 625, 375.3, 376.1, 501.6,...
			} `xml:"subject.ndc.null"`
			RelationNullNull struct {
				Text string `xml:",chardata"` // is Version of: http://ci....
			} `xml:"relation.null.null"`
			RightsNullNull []struct {
				Text string `xml:",chardata"` // ç¤¾å£æ³äººä¸­å½ç ç©¶æ...
			} `xml:"rights.null.null"`
			SortkeyNullEn struct {
				Text string `xml:",chardata"` // 001, 002, 004, 005, 006, ...
			} `xml:"sortkey.null.en"`
			TitleAlternativeNull struct {
				Text string `xml:",chardata"` // The "May Fourth" New Cult...
			} `xml:"title.alternative.null"`
			IdentifierNiiissnNull struct {
				Text string `xml:",chardata"` // 0073-229X, 0286-9136, 028...
			} `xml:"identifier.niiissn.null"`
			LanguageNullNull struct {
				Text string `xml:",chardata"` // ja, ja, ja, ja, ja, ja, j...
			} `xml:"language.null.null"`
			ContributorAuthorEn struct {
				Text string `xml:",chardata"` // Masaki, Shinzo, Masaki, S...
			} `xml:"contributor.author.en"`
			DescriptionAbstractNull struct {
				Text string `xml:",chardata"` // Series of experiments hav...
			} `xml:"description.abstract.null"`
			PublisherNullJa struct {
				Text string `xml:",chardata"` // æ¥æ¬æè«å­¦ä¼, æ¥æ...
			} `xml:"publisher.null.ja"`
			RelationIsversionofNull struct {
				Text string `xml:",chardata"` // http://ci.nii.ac.jp/naid/...
			} `xml:"relation.isversionof.null"`
			RightsNullJa []struct {
				Text string `xml:",chardata"` // æ¥æ¬æè«å­¦ä¼, æ¬æ...
			} `xml:"rights.null.ja"`
			SubjectNiiJa struct {
				Text string `xml:",chardata"` // çç©å­¦, çç©å­¦, æ...
			} `xml:"subject.nii.ja"`
			IdentifierJtitleJa struct {
				Text string `xml:",chardata"` // æè«, æè«, å¼åå¤§...
			} `xml:"identifier.jtitle.ja"`
			ContributorAlternativeJa struct {
				Text string `xml:",chardata"` // æ­£æ¨,é²ä¸, æ­£æ¨, é...
			} `xml:"contributor.alternative.ja"`
			ContributorTranscriptionJa []struct {
				Text string `xml:",chardata"` // ããµã­, ã·ã³ã¾ã¦, ...
			} `xml:"contributor.transcription.ja"`
			ContributorAuthorJa []struct {
				Text string `xml:",chardata"` // é½è¤, å°å­, é·è°·å·...
			} `xml:"contributor.author.ja"`
			ContributorOtherJa []struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦æè²å­¦é¨å...
			} `xml:"contributor.other.ja"`
			IdentifierCitationJa struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦æè²å­¦é¨ç...
			} `xml:"identifier.citation.ja"`
			DescriptionAbstractJa struct {
				Text string `xml:",chardata"` // å­¦æ ¡ã«ãããé£æè...
			} `xml:"description.abstract.ja"`
			SubjectNullJa []struct {
				Text string `xml:",chardata"` // å­¦æ ¡æ é¤è·å¡, æ é¤...
			} `xml:"subject.null.ja"`
			TitleNullJa struct {
				Text string `xml:",chardata"` // å­¦æ ¡ã«ãããé£æè...
			} `xml:"title.null.ja"`
			TitleTranscriptionJa struct {
				Text string `xml:",chardata"` // ã¬ãã³ã¦ ããªã±ã«...
			} `xml:"title.transcription.ja"`
			ContributorAlternativeEn []struct {
				Text string `xml:",chardata"` // Saito, Hisako, Hasegawa, ...
			} `xml:"contributor.alternative.en"`
			RelationIspartofseriesEn []struct {
				Text string `xml:",chardata"` // æ¥æ¬éè¡éèç ç©¶æ...
			} `xml:"relation.ispartofseries.en"`
			TypeNullNull struct {
				Text string `xml:",chardata"` // Thesis, Research Paper, R...
			} `xml:"type.null.null"`
			ContributorAdvisorNull struct {
				Text string `xml:",chardata"` // Imai, Tamiko, Department ...
			} `xml:"contributor.advisor.null"`
			LanguageNull struct {
				Text string `xml:",chardata"` // ja, ja
			} `xml:"language.null."`
			SubjectOtherNull struct {
				Text string `xml:",chardata"` // ç¤¾ä¼ã»æå, è¨èªå­...
			} `xml:"subject.other.null"`
			AddressNullNull struct {
				Text string `xml:",chardata"` // çµæ¸å­¦, æ°´ç£å­¦, æ¬...
			} `xml:"address.null.null"`
			IdentifierNcidNull struct {
				Text string `xml:",chardata"` // BN09775946, AN10212330, A...
			} `xml:"identifier.ncid.null"`
			IdentifierSpageNull struct {
				Text string `xml:",chardata"` // 55, 99, 1, 1, 75, 199, 14...
			} `xml:"identifier.spage.null"`
			IdentifierEpageNull struct {
				Text string `xml:",chardata"` // 76, 120, 278, 5, 6, 37, 8...
			} `xml:"identifier.epage.null"`
			IdentifierISBNNull struct {
				Text string `xml:",chardata"` // 4335610025, 4642032541, 4...
			} `xml:"identifier.isbn.null"`
			IdentifierJtitleNull struct {
				Text string `xml:",chardata"` // å­£åã¢ã¹ãã£ãªã³,...
			} `xml:"identifier.jtitle.null"`
			IdentifierIssueNull struct {
				Text string `xml:",chardata"` // Supplement, 3/4, 26, 7, 2...
			} `xml:"identifier.issue.null"`
			SubjectBshEn struct {
				Text string `xml:",chardata"` // å¬çäºæ¥­, å¬çäºæ¥...
			} `xml:"subject.bsh.en"`
			SubjectNdlshEn struct {
				Text string `xml:",chardata"` // å¬çäºæ¥­, å¬ä¼æ¥­, ...
			} `xml:"subject.ndlsh.en"`
			SubjectNdlcEn struct {
				Text string `xml:",chardata"` // DH151, 453.8, 376.4, GB11...
			} `xml:"subject.ndlc.en"`
			DescriptionTableofcontentsNull []struct {
				Text string `xml:",chardata"` // ç¬¬1ç«  å¬çä¼æ¥­æ¦å¿...
			} `xml:"description.tableofcontents.null"`
			IdentifierOtherNull struct {
				Text string `xml:",chardata"` // 0287-4318, 0287-4318, 028...
			} `xml:"identifier.other.null"`
			IdentifierCitationEnUS struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦åå²ç ç©¶....
			} `xml:"identifier.citation.en_US"`
			LanguageIsoEnUS struct {
				Text string `xml:",chardata"` // ja, ja, ja, ja, ja, ja, j...
			} `xml:"language.iso.en_US"`
			PublisherNullEnUS struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦åå²ç ç©¶ä...
			} `xml:"publisher.null.en_US"`
			TitleNullEnUS struct {
				Text string `xml:",chardata"` // è¡¨ç´ã»å¥¥ä», ä»å¾³ç´...
			} `xml:"title.null.en_US"`
			TypeNullEnUS []struct {
				Text string `xml:",chardata"` // Other, Article, Thesis, O...
			} `xml:"type.null.en_US"`
			SubjectNiiEnUS struct {
				Text string `xml:",chardata"` // æ­´å²å­¦, æ­´å²å­¦, æ...
			} `xml:"subject.nii.en_US"`
			TypeNiiEnUS struct {
				Text string `xml:",chardata"` // Others, Journal Article, ...
			} `xml:"type.nii.en_US"`
			IdentifierNiiissnEnUS struct {
				Text string `xml:",chardata"` // 0287-4318, 0287-4318, 134...
			} `xml:"identifier.niiissn.en_US"`
			IdentifierNcidEnUS struct {
				Text string `xml:",chardata"` // AN00211772, AN00211772, A...
			} `xml:"identifier.ncid.en_US"`
			IdentifierJtitleEnUS struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦åå²ç ç©¶,...
			} `xml:"identifier.jtitle.en_US"`
			IdentifierIssueEnUS struct {
				Text string `xml:",chardata"` // 104, 107, 23, 23, 23, 23,...
			} `xml:"identifier.issue.en_US"`
			RightsTextversionEnUS struct {
				Text string `xml:",chardata"` // publisher, publisher, aut...
			} `xml:"rights.textversion.en_US"`
			TitleTranscriptionEnUS struct {
				Text string `xml:",chardata"` // ãã§ã¦ã· ãªã¯ãºã±...
			} `xml:"title.transcription.en_US"`
			SortkeyNullEnUS struct {
				Text string `xml:",chardata"` // 000, 001, 0462, 005, 002,...
			} `xml:"sortkey.null.en_US"`
			ContributorOtherEnUS []struct {
				Text string `xml:",chardata"` // å²©æçç«åç©é¤¨, å¼...
			} `xml:"contributor.other.en_US"`
			SubjectNdcEnUS struct {
				Text string `xml:",chardata"` // 212.2, 370, 361.3, 200, 2...
			} `xml:"subject.ndc.en_US"`
			IdentifierSpageEnUS struct {
				Text string `xml:",chardata"` // 1, 39, 29, 42, 53, 49, 60...
			} `xml:"identifier.spage.en_US"`
			IdentifierEpageEnUS struct {
				Text string `xml:",chardata"` // 18, 70, 48, 43, 54, 52, 6...
			} `xml:"identifier.epage.en_US"`
			DescriptionNullEnUS struct {
				Text string `xml:",chardata"` // æä¸å¤§å­¦:å¼åå¤§å­¦...
			} `xml:"description.null.en_US"`
			TitleAlternativeEnUS struct {
				Text string `xml:",chardata"` // The effect of urbanism on...
			} `xml:"title.alternative.en_US"`
			DescriptionAbstractEnUS struct {
				Text string `xml:",chardata"` // æ¬ç ç©¶ã®ç®çã¯ãå...
			} `xml:"description.abstract.en_US"`
			SubjectNullEnUS []struct {
				Text string `xml:",chardata"` // çåä¿è², ä¿è²çè­...
			} `xml:"subject.null.en_US"`
			IdentifierVolumeNull struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			} `xml:"identifier.volume.null"`
			IdentifierVolumeEnUS struct {
				Text string `xml:",chardata"` // 2, 2, 2, 2, 2, 2, 2, 2, 2...
			} `xml:"identifier.volume.en_US"`
			PublisherNullNull struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦äººæå­¦é¨,...
			} `xml:"publisher.null.null"`
			DateCreatedNull struct {
				Text string `xml:",chardata"` // 1982-02-25, 21 seiki kyou...
			} `xml:"date.created.null"`
			RelationWosNull struct {
				Text string `xml:",chardata"` // A1996VN62700002, 00017717...
			} `xml:"relation.wos.null"`
			RelationPmidEn struct {
				Text string `xml:",chardata"` // 11436984, 18079584, 18974...
			} `xml:"relation.pmid.en"`
			RelationDoiEnUS struct {
				Text string `xml:",chardata"` // 10.1017/S1479591409000278...
			} `xml:"relation.doi.en_US"`
			SubjectNullNull []struct {
				Text string `xml:",chardata"` // valproate, hepatic injury...
			} `xml:"subject.null.null"`
			RelationURINull struct {
				Text string `xml:",chardata"` // http://wwwsoc.nii.ac.jp/s...
			} `xml:"relation.uri.null"`
			SubjectNdlcEnUS struct {
				Text string `xml:",chardata"` // 493.1, äººæç§å­¦, 341,...
			} `xml:"subject.ndlc.en_US"`
			SubjectBshEnUS struct {
				Text string `xml:",chardata"` // è©å -- æå¤ªé, è©©(...
			} `xml:"subject.bsh.en_US"`
			DescriptionNullJa struct {
				Text string `xml:",chardata"` // å¹³æ15å¹´åº¦ï½å¹³æ17...
			} `xml:"description.null.ja"`
			SubjectBshNull struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦éå±å³æ¸é...
			} `xml:"subject.bsh.null"`
			DateNullNull struct {
				Text string `xml:",chardata"` // 2013, 2013, 2013
			} `xml:"date.null.null"`
			DateSubmittedNull struct {
				Text string `xml:",chardata"` // 2003-01-27, 2003-01-27, 2...
			} `xml:"date.submitted.null"`
			RelationPmidEnUS struct {
				Text string `xml:",chardata"` // 23114857, 23417602, 25007...
			} `xml:"relation.pmid.en_US"`
			IdentifierGrantidEnUS struct {
				Text string `xml:",chardata"` // 11101A1812, 11101A1813, 1...
			} `xml:"identifier.grantid.en_US"`
			DateGrantedEnUS struct {
				Text string `xml:",chardata"` // 2013-09-30, 2013-09-30, 2...
			} `xml:"date.granted.en_US"`
			DescriptionDegreenameEnUS struct {
				Text string `xml:",chardata"` // åå£«ï¼å»å­¦ï¼, åå£...
			} `xml:"description.degreename.en_US"`
			DescriptionGrantorEnUS struct {
				Text string `xml:",chardata"` // å¼åå¤§å­¦, å¼åå¤§å­...
			} `xml:"description.grantor.en_US"`
			RelationDoiNull struct {
				Text string `xml:",chardata"` // 10.2176/nmc.oa2012-0204, ...
			} `xml:"relation.doi.null"`
			ContributorNullNull struct {
				Text string `xml:",chardata"` // åç°, çäºº
			} `xml:"contributor.null.null"`
			TypeNiiNull struct {
				Text string `xml:",chardata"` // Thesis or Dissertation, T...
			} `xml:"type.nii.null"`
			RightsTextversionNull struct {
				Text string `xml:",chardata"` // ETD, ETD
			} `xml:"rights.textversion.null"`
			RelationPmidNull struct {
				Text string `xml:",chardata"` // 27644042
			} `xml:"relation.pmid.null"`
		} `xml:"dc"`
		Bitstreaminfo struct {
			Text        string `xml:",chardata"`
			FullTextURL []struct {
				Text string `xml:",chardata"` // http://repository.ul.hiro...
			} `xml:"fullTextURL"`
			BitstreamName []struct {
				Text string `xml:",chardata"` // Menges,Vessels,Cranial ne...
			} `xml:"bitstreamName"`
		} `xml:"bitstreaminfo"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
