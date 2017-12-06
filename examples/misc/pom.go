package misc

import "encoding/xml"

// POM was generated 2017-12-06 00:05:23 by tir on apollo.
type POM struct {
	XMLName        xml.Name `xml:"project"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	ModelVersion   struct {
		Text string `xml:",chardata"` // 4.0.0
	} `xml:"modelVersion"`
	GroupId struct {
		Text string `xml:",chardata"` // org.springframework.samples
	} `xml:"groupId"`
	ArtifactId struct {
		Text string `xml:",chardata"` // spring-mvc-showcase
	} `xml:"artifactId"`
	Name struct {
		Text string `xml:",chardata"` // spring-mvc-showcase
	} `xml:"name"`
	Packaging struct {
		Text string `xml:",chardata"` // war
	} `xml:"packaging"`
	Version struct {
		Text string `xml:",chardata"` // 1.0.0-BUILD-SNAPSHOT
	} `xml:"version"`
	Properties struct {
		Text        string `xml:",chardata"`
		JavaVersion struct {
			Text string `xml:",chardata"` // 1.7
		} `xml:"java-version"`
		OrgSpringframeworkVersion struct {
			Text string `xml:",chardata"` // 4.2.2.RELEASE
		} `xml:"org.springframework-version"`
		OrgSpringframeworkSecurityVersion struct {
			Text string `xml:",chardata"` // 4.0.1.RELEASE
		} `xml:"org.springframework.security-version"`
		OrgAspectjVersion struct {
			Text string `xml:",chardata"` // 1.8.1
		} `xml:"org.aspectj-version"`
		OrgSlf4jVersion struct {
			Text string `xml:",chardata"` // 1.7.12
		} `xml:"org.slf4j-version"`
	} `xml:"properties"`
	Dependencies struct {
		Text       string `xml:",chardata"`
		Dependency []struct {
			Text    string `xml:",chardata"`
			GroupId struct {
				Text string `xml:",chardata"` // org.springframework, org.springframework...
			} `xml:"groupId"`
			ArtifactId struct {
				Text string `xml:",chardata"` // spring-context, spring-webmvc, aspectjrt...
			} `xml:"artifactId"`
			Version struct {
				Text string `xml:",chardata"` // ${org.springframework-version}, ${org.sp...
			} `xml:"version"`
			Exclusions struct {
				Text      string `xml:",chardata"`
				Exclusion struct {
					Text    string `xml:",chardata"`
					GroupId struct {
						Text string `xml:",chardata"` // commons-logging, javax.servlet, javax.se...
					} `xml:"groupId"`
					ArtifactId struct {
						Text string `xml:",chardata"` // commons-logging, servlet-api, servlet-ap...
					} `xml:"artifactId"`
				} `xml:"exclusion"`
			} `xml:"exclusions"`
			Scope struct {
				Text string `xml:",chardata"` // runtime, runtime, runtime, provided, pro...
			} `xml:"scope"`
		} `xml:"dependency"`
	} `xml:"dependencies"`
	Repositories struct {
		Text       string `xml:",chardata"`
		Repository []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"` // org.springframework.maven.snapshot, org....
			} `xml:"id"`
			Name struct {
				Text string `xml:",chardata"` // Spring Maven Snapshot Repository, Spring...
			} `xml:"name"`
			URL struct {
				Text string `xml:",chardata"` // http://repo.spring.io/snapshot, http://r...
			} `xml:"url"`
			Releases struct {
				Text    string `xml:",chardata"`
				Enabled struct {
					Text string `xml:",chardata"` // false
				} `xml:"enabled"`
			} `xml:"releases"`
			Snapshots struct {
				Text    string `xml:",chardata"`
				Enabled struct {
					Text string `xml:",chardata"` // true, false
				} `xml:"enabled"`
			} `xml:"snapshots"`
		} `xml:"repository"`
	} `xml:"repositories"`
	Build struct {
		Text      string `xml:",chardata"`
		FinalName struct {
			Text string `xml:",chardata"` // ${project.artifactId}
		} `xml:"finalName"`
		Plugins struct {
			Text   string `xml:",chardata"`
			Plugin []struct {
				Text    string `xml:",chardata"`
				GroupId struct {
					Text string `xml:",chardata"` // org.apache.maven.plugins, org.apache.mav...
				} `xml:"groupId"`
				ArtifactId struct {
					Text string `xml:",chardata"` // maven-compiler-plugin, maven-surefire-pl...
				} `xml:"artifactId"`
				Version struct {
					Text string `xml:",chardata"` // 2.3.2, 2.12, 1.2, 2.2, 9.0.6.v20130930
				} `xml:"version"`
				Configuration struct {
					Text   string `xml:",chardata"`
					Source struct {
						Text string `xml:",chardata"` // ${java-version}, ${java-version}
					} `xml:"source"`
					Target struct {
						Text string `xml:",chardata"` // ${java-version}, ${java-version}
					} `xml:"target"`
					Includes struct {
						Text    string `xml:",chardata"`
						Include struct {
							Text string `xml:",chardata"` // **/*Tests.java
						} `xml:"include"`
					} `xml:"includes"`
					Excludes struct {
						Text    string `xml:",chardata"`
						Exclude struct {
							Text string `xml:",chardata"` // **/Abstract*.java
						} `xml:"exclude"`
					} `xml:"excludes"`
					JunitArtifactName struct {
						Text string `xml:",chardata"` // junit:junit
					} `xml:"junitArtifactName"`
					ArgLine struct {
						Text string `xml:",chardata"` // -Xmx512m
					} `xml:"argLine"`
					Outxml struct {
						Text string `xml:",chardata"` // true
					} `xml:"outxml"`
					WebApp struct {
						Text        string `xml:",chardata"`
						ContextPath struct {
							Text string `xml:",chardata"` // /${project.artifactId}
						} `xml:"contextPath"`
					} `xml:"webApp"`
				} `xml:"configuration"`
				Executions struct {
					Text      string `xml:",chardata"`
					Execution struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"` // install
						} `xml:"id"`
						Phase struct {
							Text string `xml:",chardata"` // install
						} `xml:"phase"`
						Goals struct {
							Text string `xml:",chardata"`
							Goal []struct {
								Text string `xml:",chardata"` // sources, compile, test-compile
							} `xml:"goal"`
						} `xml:"goals"`
					} `xml:"execution"`
				} `xml:"executions"`
				Dependencies struct {
					Text       string `xml:",chardata"`
					Dependency []struct {
						Text    string `xml:",chardata"`
						GroupId struct {
							Text string `xml:",chardata"` // org.aspectj, org.aspectj
						} `xml:"groupId"`
						ArtifactId struct {
							Text string `xml:",chardata"` // aspectjrt, aspectjtools
						} `xml:"artifactId"`
						Version struct {
							Text string `xml:",chardata"` // ${org.aspectj-version}, ${org.aspectj-ve...
						} `xml:"version"`
					} `xml:"dependency"`
				} `xml:"dependencies"`
			} `xml:"plugin"`
		} `xml:"plugins"`
	} `xml:"build"`
}
