package tdata

import "fmt"

// Get returns r
func Get(s string) (string, error) {
	switch s {
	case "cfg":
		return Cfg, nil
	case "cfg2":
		return CfgPartial, nil
	case "cfg3":
		return CfgMimimal, nil
	case "fileManifest":
		return FileManifest, nil
	}
	return "", fmt.Errorf("function Get(%v) is unknown", s)
}

var (
	// Cfg contains a valid full EdgeOS blacklist configuration
	Cfg = `blacklist {
    disabled false
    dns-redirect-ip 0.0.0.0
    domains {
        dns-redirect-ip 192.168.100.1
        include adsrvr.org
        include adtechus.net
        include advertising.com
        include centade.com
        include doubleclick.net
        include free-counter.co.uk
        include intellitxt.com
        include kiosked.com
        include patoghee.in
        source malc0de {
            dns-redirect-ip 192.168.168.1
            description "List of zones serving malicious executables observed by malc0de.com/database/"
            prefix "zone "
            url http://malc0de.com/bl/ZONES
        }
        source malwaredomains.com {
            dns-redirect-ip 10.0.0.1
            description "Just domains"
            prefix ""
            url http://mirror1.malwaredomains.com/files/justdomains
        }
        source simple_tracking {
            description "Basic tracking list by Disconnect"
            prefix ""
            url https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt
        }
        source zeus {
            description "abuse.ch ZeuS domain blocklist"
            prefix ""
            url https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist
        }
    }
    exclude 1e100.net
    exclude 2o7.net
    exclude adobedtm.com
    exclude akamai.net
    exclude akamaihd.net
    exclude amazon.com
    exclude amazonaws.com
    exclude apple.com
    exclude ask.com
    exclude avast.com
    exclude bitdefender.com
    exclude cdn.visiblemeasures.com
    exclude cloudfront.net
    exclude coremetrics.com
    exclude edgesuite.net
    exclude freedns.afraid.org
    exclude github.com
    exclude githubusercontent.com
    exclude google.com
    exclude googleadservices.com
    exclude googleapis.com
    exclude googletagmanager.com
    exclude googleusercontent.com
    exclude gstatic.com
    exclude gvt1.com
    exclude gvt1.net
    exclude hb.disney.go.com
    exclude hp.com
    exclude hulu.com
    exclude images-amazon.com
    exclude live.com
    exclude microsoft.com
    exclude msdn.com
    exclude msecnd.net
    exclude paypal.com
    exclude rackcdn.com
    exclude schema.org
    exclude shopify.com
    exclude skype.com
    exclude smacargo.com
    exclude sourceforge.net
    exclude ssl-on9.com
    exclude ssl-on9.net
    exclude sstatic.net
    exclude static.chartbeat.com
    exclude storage.googleapis.com
    exclude windows.net
    exclude xboxlive.com
    exclude yimg.com
    exclude ytimg.com
    hosts {
        include beap.gemini.yahoo.com
        source openphish {
            description "OpenPhish automatic phishing detection"
            prefix http
            url https://openphish.com/feed.txt
        }
        source raw.github.com {
            description "This hosts file is a merged collection of hosts from reputable sources"
            prefix "0.0.0.0 "
            url https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts
        }
        source sysctl.org {
            dns-redirect-ip 172.16.16.1
            description "This hosts file is a merged collection of hosts from cameleon"
            prefix "127.0.0.1	 "
            url http://sysctl.org/cameleon/hosts
        }
        source tasty {
            description "File source"
            dns-redirect-ip 10.10.10.10
            file ../internal/testdata/blist.hosts.src
        }
        source volkerschatz {
            description "Ad server blacklists"
            prefix http
            url http://www.volkerschatz.com/net/adpaths
        }
        source yoyo {
            description "Fully Qualified Domain Names only - no prefix to strip"
            prefix ""
            url https://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml&showintro=1&mimetype=plaintext
        }
    }
}

	/* Warning: Do not remove the following line. */
	/* === vyatta-config-version: "config-management@1:conntrack@1:cron@1:dhcp-relay@1:dhcp-server@4:firewall@5:ipsec@5:nat@3:qos@1:quagga@2:system@4:ubnt-pptp@1:ubnt-util@1:vrrp@1:webgui@1:webproxy@1:zone-policy@1" === */
	/* Release version: v1.8.5.4884695.160608.1057 */
}`

	// CfgPartial contains a valid partial EdgeOS blacklist configuration
	CfgPartial = `blacklist {
    disabled false
    dns-redirect-ip 0.0.0.0
    domains {
        include adsrvr.org
        include adtechus.net
        include advertising.com
        include centade.com
        include doubleclick.net
        include free-counter.co.uk
        include intellitxt.com
        include kiosked.com
        source malc0de {
            description "List of zones serving malicious executables observed by malc0de.com/database/"
            prefix "zone "
            url http://malc0de.com/bl/ZONES
        }
    }
    exclude 122.2o7.net
    exclude 1e100.net
    exclude adobedtm.com
    exclude akamai.net
    exclude amazon.com
    exclude amazonaws.com
    exclude apple.com
    exclude ask.com
    exclude avast.com
    exclude bitdefender.com
    exclude cdn.visiblemeasures.com
    exclude cloudfront.net
    exclude coremetrics.com
    exclude edgesuite.net
    exclude freedns.afraid.org
    exclude github.com
    exclude githubusercontent.com
    exclude google.com
    exclude googleadservices.com
    exclude googleapis.com
    exclude googleusercontent.com
    exclude gstatic.com
    exclude gvt1.com
    exclude gvt1.net
    exclude hb.disney.go.com
    exclude hp.com
    exclude hulu.com
    exclude images-amazon.com
    exclude msdn.com
    exclude paypal.com
    exclude rackcdn.com
    exclude schema.org
    exclude skype.com
    exclude smacargo.com
    exclude sourceforge.net
    exclude ssl-on9.com
    exclude ssl-on9.net
    exclude static.chartbeat.com
    exclude storage.googleapis.com
    exclude windows.net
    exclude yimg.com
    exclude ytimg.com
    hosts {
        include beap.gemini.yahoo.com
    }
}`

	// CfgMimimal contains a valid minimal EdgeOS blacklist configuration
	CfgMimimal = `blacklist {
    disabled false
    dns-redirect-ip 0.0.0.0
    domains {
        include adsrvr.org
        include adtechus.net
        include advertising.com
        include centade.com
        include doubleclick.net
        include free-counter.co.uk
        include intellitxt.com
        include kiosked.com
        source malc0de {
            description "List of zones serving malicious executables observed by malc0de.com/database/"
            prefix "zone "
            url http://malc0de.com/bl/ZONES
        }
    }
    exclude ytimg.com
    hosts {
        include beap.gemini.yahoo.com
        source tasty {
            description "File source"
            dns-redirect-ip 10.10.10.10
            file ../internal/testdata/blist.hosts.src
        }
    }
}`

	// DisabledCfg contains a disabled valid EdgeOS blacklist configuration
	DisabledCfg = `blacklist {
        disabled true
        dns-redirect-ip 0.0.0.0
        domains {
            dns-redirect-ip 192.168.100.1
            include adsrvr.org
            include adtechus.net
            include advertising.com
            include centade.com
            include doubleclick.net
            include free-counter.co.uk
            include intellitxt.com
            include kiosked.com
            include patoghee.in
            source malc0de {
                dns-redirect-ip 192.168.168.1
                description "List of zones serving malicious executables observed by malc0de.com/database/"
                prefix "zone "
                url http://malc0de.com/bl/ZONES
            }
            source malwaredomains.com {
                dns-redirect-ip 10.0.0.1
                description "Just domains"
                prefix ""
                url http://mirror1.malwaredomains.com/files/justdomains
            }
            source simple_tracking {
                description "Basic tracking list by Disconnect"
                prefix ""
                url https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt
            }
            source zeus {
                description "abuse.ch ZeuS domain blocklist"
                prefix ""
                url https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist
            }
        }
        exclude 1e100.net
        exclude 2o7.net
        exclude adobedtm.com
        exclude akamai.net
        exclude akamaihd.net
        exclude amazon.com
        exclude amazonaws.com
        exclude apple.com
        exclude ask.com
        exclude avast.com
        exclude bitdefender.com
        exclude cdn.visiblemeasures.com
        exclude cloudfront.net
        exclude coremetrics.com
        exclude edgesuite.net
        exclude freedns.afraid.org
        exclude github.com
        exclude githubusercontent.com
        exclude google.com
        exclude googleadservices.com
        exclude googleapis.com
        exclude googletagmanager.com
        exclude googleusercontent.com
        exclude gstatic.com
        exclude gvt1.com
        exclude gvt1.net
        exclude hb.disney.go.com
        exclude hp.com
        exclude hulu.com
        exclude images-amazon.com
        exclude live.com
        exclude microsoft.com
        exclude msdn.com
        exclude msecnd.net
        exclude paypal.com
        exclude rackcdn.com
        exclude schema.org
        exclude shopify.com
        exclude skype.com
        exclude smacargo.com
        exclude sourceforge.net
        exclude ssl-on9.com
        exclude ssl-on9.net
        exclude sstatic.net
        exclude static.chartbeat.com
        exclude storage.googleapis.com
        exclude windows.net
        exclude xboxlive.com
        exclude yimg.com
        exclude ytimg.com
        hosts {
            include beap.gemini.yahoo.com
            source openphish {
                description "OpenPhish automatic phishing detection"
                prefix http
                url https://openphish.com/feed.txt
            }
            source raw.github.com {
                description "This hosts file is a merged collection of hosts from reputable sources"
                prefix "0.0.0.0 "
                url https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts
            }
            source sysctl.org {
                dns-redirect-ip 172.16.16.1
                description "This hosts file is a merged collection of hosts from cameleon"
                prefix "127.0.0.1	 "
                url http://sysctl.org/cameleon/hosts
            }
            source tasty {
                description "File source"
                dns-redirect-ip 10.10.10.10
                file ./internal/testdata/blist.hosts.src
            }
            source volkerschatz {
                description "Ad server blacklists"
                prefix http
                url http://www.volkerschatz.com/net/adpaths
            }
            source yoyo {
                description "Fully Qualified Domain Names only - no prefix to strip"
                prefix ""
                url https://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml&showintro=1&mimetype=plaintext
            }
        }
    }

        /* Warning: Do not remove the following line. */
        /* === vyatta-config-version: "config-management@1:conntrack@1:cron@1:dhcp-relay@1:dhcp-server@4:firewall@5:ipsec@5:nat@3:qos@1:quagga@2:system@4:ubnt-pptp@1:ubnt-util@1:vrrp@1:webgui@1:webproxy@1:zone-policy@1" === */
        /* Release version: v1.8.5.4884695.160608.1057 */
    }`

	// ZeroHostSourcesCfg is a valid EdgeOS blacklist configuration with zero hosts sources
	ZeroHostSourcesCfg = `blacklist {
	disabled false
	dns-redirect-ip 0.0.0.0
	domains {
			dns-redirect-ip
			include adsrvr.org
			include adtechus.net
			include advertising.com
			include centade.com
			include doubleclick.net
			include free-counter.co.uk
			include intellitxt.com
			include kiosked.com
			source malc0de {
					description "List of zones serving malicious executables observed by malc0de.com/database/"
					prefix "zone "
					url http://malc0de.com/bl/ZONES
			}
	}
	exclude 122.2o7.net
	exclude 1e100.net
	exclude adobedtm.com
	exclude akamai.net
	exclude amazon.com
	exclude amazonaws.com
	exclude apple.com
	exclude ask.com
	exclude avast.com
	exclude bitdefender.com
	exclude cdn.visiblemeasures.com
	exclude cloudfront.net
	exclude coremetrics.com
	exclude edgesuite.net
	exclude freedns.afraid.org
	exclude github.com
	exclude githubusercontent.com
	exclude google.com
	exclude googleadservices.com
	exclude googleapis.com
	exclude googleusercontent.com
	exclude gstatic.com
	exclude gvt1.com
	exclude gvt1.net
	exclude hb.disney.go.com
	exclude hp.com
	exclude hulu.com
	exclude images-amazon.com
	exclude msdn.com
	exclude paypal.com
	exclude rackcdn.com
	exclude schema.org
	exclude skype.com
	exclude smacargo.com
	exclude sourceforge.net
	exclude ssl-on9.com
	exclude ssl-on9.net
	exclude static.chartbeat.com
	exclude storage.googleapis.com
	exclude windows.net
	exclude yimg.com
	exclude ytimg.com
	hosts {
			dns-redirect-ip
			include beap.gemini.yahoo.com
	}
}`

	// NoBlacklist returns a VyOS error message
	NoBlacklist = `Configuration under specified path is empty`

	// JSONcfg is JSON formatted blacklist configuration output
	JSONcfg = "{\n  \"nodes\": [{\n    \"blacklist\": {\n      \"disabled\": \"false\",\n      \"ip\": \"0.0.0.0\",\n      \"excludes\": [\n        \"1e100.net\",\n        \"2o7.net\",\n        \"adobedtm.com\",\n        \"akamai.net\",\n        \"akamaihd.net\",\n        \"amazon.com\",\n        \"amazonaws.com\",\n        \"apple.com\",\n        \"ask.com\",\n        \"avast.com\",\n        \"bitdefender.com\",\n        \"cdn.visiblemeasures.com\",\n        \"cloudfront.net\",\n        \"coremetrics.com\",\n        \"edgesuite.net\",\n        \"freedns.afraid.org\",\n        \"github.com\",\n        \"githubusercontent.com\",\n        \"google.com\",\n        \"googleadservices.com\",\n        \"googleapis.com\",\n        \"googletagmanager.com\",\n        \"googleusercontent.com\",\n        \"gstatic.com\",\n        \"gvt1.com\",\n        \"gvt1.net\",\n        \"hb.disney.go.com\",\n        \"hp.com\",\n        \"hulu.com\",\n        \"images-amazon.com\",\n        \"live.com\",\n        \"microsoft.com\",\n        \"msdn.com\",\n        \"msecnd.net\",\n        \"paypal.com\",\n        \"rackcdn.com\",\n        \"schema.org\",\n        \"shopify.com\",\n        \"skype.com\",\n        \"smacargo.com\",\n        \"sourceforge.net\",\n        \"ssl-on9.com\",\n        \"ssl-on9.net\",\n        \"sstatic.net\",\n        \"static.chartbeat.com\",\n        \"storage.googleapis.com\",\n        \"windows.net\",\n        \"xboxlive.com\",\n        \"yimg.com\",\n        \"ytimg.com\"\n        ]\n    },\n    \"domains\": {\n      \"disabled\": \"false\",\n      \"ip\": \"192.168.100.1\",\n      \"excludes\": [],\n      \"includes\": [\n        \"adsrvr.org\",\n        \"adtechus.net\",\n        \"advertising.com\",\n        \"centade.com\",\n        \"doubleclick.net\",\n        \"free-counter.co.uk\",\n        \"intellitxt.com\",\n        \"kiosked.com\",\n        \"patoghee.in\"\n        ],\n      \"sources\": [{\n        \"malc0de\": {\n          \"disabled\": \"false\",\n          \"description\": \"List of zones serving malicious executables observed by malc0de.com/database/\",\n          \"ip\": \"192.168.168.1\",\n          \"prefix\": \"zone \",\n          \"url\": \"http://malc0de.com/bl/ZONES\",\n        },\n        \"malwaredomains.com\": {\n          \"disabled\": \"false\",\n          \"description\": \"Just domains\",\n          \"ip\": \"10.0.0.1\",\n          \"url\": \"http://mirror1.malwaredomains.com/files/justdomains\",\n        },\n        \"simple_tracking\": {\n          \"disabled\": \"false\",\n          \"description\": \"Basic tracking list by Disconnect\",\n          \"url\": \"https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt\",\n        },\n        \"zeus\": {\n          \"disabled\": \"false\",\n          \"description\": \"abuse.ch ZeuS domain blocklist\",\n          \"url\": \"https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist\",\n        }\n    }]\n    },\n    \"hosts\": {\n      \"disabled\": \"false\",\n      \"excludes\": [],\n      \"includes\": [\"beap.gemini.yahoo.com\"],\n      \"sources\": [{\n        \"openphish\": {\n          \"disabled\": \"false\",\n          \"description\": \"OpenPhish automatic phishing detection\",\n          \"prefix\": \"http\",\n          \"url\": \"https://openphish.com/feed.txt\",\n        },\n        \"raw.github.com\": {\n          \"disabled\": \"false\",\n          \"description\": \"This hosts file is a merged collection of hosts from reputable sources\",\n          \"prefix\": \"0.0.0.0 \",\n          \"url\": \"https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts\",\n        },\n        \"sysctl.org\": {\n          \"disabled\": \"false\",\n          \"description\": \"This hosts file is a merged collection of hosts from cameleon\",\n          \"ip\": \"172.16.16.1\",\n          \"prefix\": \"127.0.0.1\\t \",\n          \"url\": \"http://sysctl.org/cameleon/hosts\",\n        },\n        \"tasty\": {\n          \"disabled\": \"false\",\n          \"description\": \"File source\",\n          \"ip\": \"10.10.10.10\",\n          \"file\": \"../internal/testdata/blist.hosts.src\",\n        },\n        \"volkerschatz\": {\n          \"disabled\": \"false\",\n          \"description\": \"Ad server blacklists\",\n          \"prefix\": \"http\",\n          \"url\": \"http://www.volkerschatz.com/net/adpaths\",\n        },\n        \"yoyo\": {\n          \"disabled\": \"false\",\n          \"description\": \"Fully Qualified Domain Names only - no prefix to strip\",\n          \"url\": \"https://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml&showintro=1&mimetype=plaintext\",\n        }\n    }]\n    }\n  }]\n}"

	// JSONcfgZeroHostSources is JSON formatted blacklist configuration output with zero sources for hosts
	JSONcfgZeroHostSources = "{\n  \"nodes\": [{\n    \"blacklist\": {\n      \"disabled\": \"false\",\n      \"ip\": \"0.0.0.0\",\n      \"excludes\": [\n        \"122.2o7.net\",\n        \"1e100.net\",\n        \"adobedtm.com\",\n        \"akamai.net\",\n        \"amazon.com\",\n        \"amazonaws.com\",\n        \"apple.com\",\n        \"ask.com\",\n        \"avast.com\",\n        \"bitdefender.com\",\n        \"cdn.visiblemeasures.com\",\n        \"cloudfront.net\",\n        \"coremetrics.com\",\n        \"edgesuite.net\",\n        \"freedns.afraid.org\",\n        \"github.com\",\n        \"githubusercontent.com\",\n        \"google.com\",\n        \"googleadservices.com\",\n        \"googleapis.com\",\n        \"googleusercontent.com\",\n        \"gstatic.com\",\n        \"gvt1.com\",\n        \"gvt1.net\",\n        \"hb.disney.go.com\",\n        \"hp.com\",\n        \"hulu.com\",\n        \"images-amazon.com\",\n        \"msdn.com\",\n        \"paypal.com\",\n        \"rackcdn.com\",\n        \"schema.org\",\n        \"skype.com\",\n        \"smacargo.com\",\n        \"sourceforge.net\",\n        \"ssl-on9.com\",\n        \"ssl-on9.net\",\n        \"static.chartbeat.com\",\n        \"storage.googleapis.com\",\n        \"windows.net\",\n        \"yimg.com\",\n        \"ytimg.com\"\n        ]\n    },\n    \"domains\": {\n      \"disabled\": \"false\",\n      \"excludes\": [],\n      \"includes\": [\n        \"adsrvr.org\",\n        \"adtechus.net\",\n        \"advertising.com\",\n        \"centade.com\",\n        \"doubleclick.net\",\n        \"free-counter.co.uk\",\n        \"intellitxt.com\",\n        \"kiosked.com\"\n        ],\n      \"sources\": [{\n        \"malc0de\": {\n          \"disabled\": \"false\",\n          \"description\": \"List of zones serving malicious executables observed by malc0de.com/database/\",\n          \"prefix\": \"zone \",\n          \"url\": \"http://malc0de.com/bl/ZONES\",\n        }\n    }]\n    },\n    \"hosts\": {\n      \"disabled\": \"false\",\n      \"excludes\": [],\n      \"includes\": [\"beap.gemini.yahoo.com\"],\n      \"sources\": [{}]\n    }\n  }]\n}"

	// JSONrawcfg is JSON unformatted blacklist configuration output
	JSONrawcfg = `{"blacklist":{"data":{},"disable":false,"excludes":["122.2o7.net","1e100.net","adobedtm.com","akamai.net","amazon.com","amazonaws.com","apple.com","ask.com","avast.com","bitdefender.com","cdn.visiblemeasures.com","cloudfront.net","coremetrics.com","edgesuite.net","freedns.afraid.org","github.com","githubusercontent.com","google.com","googleadservices.com","googleapis.com","googleusercontent.com","gstatic.com","gvt1.com","gvt1.net","hb.disney.go.com","hp.com","hulu.com","images-amazon.com","msdn.com","paypal.com","rackcdn.com","schema.org","skype.com","smacargo.com","sourceforge.net","ssl-on9.com","ssl-on9.net","static.chartbeat.com","storage.googleapis.com","windows.net","yimg.com","ytimg.com"],"includes":[],"ip":"0.0.0.0"},"domains":{"data":{"malc0de":{"desc":"List of zones serving malicious executables observed by malc0de.com/database/","disabled":false,"file":"","ip":"","name":"malc0de","prefix":"zone ","type":2,"url":"http://malc0de.com/bl/ZONES"}},"disable":false,"excludes":[],"includes":["adsrvr.org","adtechus.net","advertising.com","centade.com","doubleclick.net","free-counter.co.uk","intellitxt.com","kiosked.com"],"ip":"0.0.0.0"},"hosts":{"data":{"adaway":{"desc":"Blocking mobile ad providers and some analytics providers","disabled":false,"file":"","ip":"","name":"adaway","prefix":"127.0.0.1 ","type":3,"url":"http://adaway.org/hosts.txt"},"malwaredomainlist":{"desc":"127.0.0.1 based host and domain list","disabled":false,"file":"","ip":"","name":"malwaredomainlist","prefix":"127.0.0.1 ","type":3,"url":"http://www.malwaredomainlist.com/hostslist/hosts.txt"},"openphish":{"desc":"OpenPhish automatic phishing detection","disabled":false,"file":"","ip":"","name":"openphish","prefix":"http","type":3,"url":"https://openphish.com/feed.txt"},"someonewhocares":{"desc":"Zero based host and domain list","disabled":false,"file":"","ip":"","name":"someonewhocares","prefix":"0.0.0.0","type":3,"url":"http://someonewhocares.org/hosts/zero/"},"tasty":{"desc":"File source","disabled":false,"file":"../internal/testdata/blist.hosts.src","ip":"0.0.0.0","name":"tasty","prefix":"","type":3,"url":""},"volkerschatz":{"desc":"Ad server blacklists","disabled":false,"file":"","ip":"","name":"volkerschatz","prefix":"http","type":3,"url":"http://www.volkerschatz.com/net/adpaths"},"winhelp2002":{"desc":"Zero based host and domain list","disabled":false,"file":"","ip":"0.0.0.0","name":"winhelp2002","prefix":"0.0.0.0 ","type":3,"url":"http://winhelp2002.mvps.org/hosts.txt"},"yoyo":{"desc":"Fully Qualified Domain Names only - no prefix to strip","disabled":false,"file":"","ip":"","name":"yoyo","prefix":"","type":3,"url":"http://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml\u0026showintro=1\u0026mimetype=plaintext"}},"disable":false,"excludes":[],"includes":["beap.gemini.yahoo.com"],"ip":"192.168.168.1"}}`

	// FileManifest is complete list of the blacklist config node templates
	FileManifest = `../payload/blacklist
../payload/blacklist/disabled
../payload/blacklist/disabled/node.def
../payload/blacklist/dns-redirect-ip
../payload/blacklist/dns-redirect-ip/node.def
../payload/blacklist/domains
../payload/blacklist/domains/dns-redirect-ip
../payload/blacklist/domains/dns-redirect-ip/node.def
../payload/blacklist/domains/exclude
../payload/blacklist/domains/exclude/node.def
../payload/blacklist/domains/include
../payload/blacklist/domains/include/node.def
../payload/blacklist/domains/node.def
../payload/blacklist/domains/source
../payload/blacklist/domains/source/node.def
../payload/blacklist/domains/source/node.tag
../payload/blacklist/domains/source/node.tag/description
../payload/blacklist/domains/source/node.tag/description/node.def
../payload/blacklist/domains/source/node.tag/prefix
../payload/blacklist/domains/source/node.tag/prefix/node.def
../payload/blacklist/domains/source/node.tag/url
../payload/blacklist/domains/source/node.tag/url/node.def
../payload/blacklist/exclude
../payload/blacklist/exclude/node.def
../payload/blacklist/hosts
../payload/blacklist/hosts/dns-redirect-ip
../payload/blacklist/hosts/dns-redirect-ip/node.def
../payload/blacklist/hosts/exclude
../payload/blacklist/hosts/exclude/node.def
../payload/blacklist/hosts/include
../payload/blacklist/hosts/include/node.def
../payload/blacklist/hosts/node.def
../payload/blacklist/hosts/source
../payload/blacklist/hosts/source/node.def
../payload/blacklist/hosts/source/node.tag
../payload/blacklist/hosts/source/node.tag/description
../payload/blacklist/hosts/source/node.tag/description/node.def
../payload/blacklist/hosts/source/node.tag/prefix
../payload/blacklist/hosts/source/node.tag/prefix/node.def
../payload/blacklist/hosts/source/node.tag/url
../payload/blacklist/hosts/source/node.tag/url/node.def
../payload/blacklist/node.def
`
	// Live is a working EdgeOS configuration
	Live = `blacklist {
        disabled false
        dns-redirect-ip 192.168.168.1
        domains {
            include adk2x.com
            include adsrvr.org
            include adtechus.net
            include advertising.com
            include centade.com
            include doubleclick.net
            include fastplayz.com
            include free-counter.co.uk
            include hilltopads.net
            include intellitxt.com
            include kiosked.com
            include patoghee.in
            include themillionaireinpjs.com
            include traktrafficflow.com
            include wwwpromoter.com
            source NoBitCoin {
                description "Blocking Web Browser Bitcoin Mining"
                prefix 0.0.0.0
                url https://raw.githubusercontent.com/hoshsadiq/adblock-nocoin-list/master/hosts.txt
            }
            source malc0de {
                description "List of zones serving malicious executables observed by malc0de.com/database/"
                prefix zone
                url http://malc0de.com/bl/ZONES
            }
            source malwaredomains.com {
                description "Just Domains"
                url http://mirror1.malwaredomains.com/files/justdomains
            }
            source simple_tracking {
                description "Basic tracking list by Disconnect"
                url https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt
            }
            source zeus {
                description "abuse.ch ZeuS domain blocklist"
                url https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist
            }
        source tasty {
            description "File source"
            dns-redirect-ip 10.10.10.10
            file ./internal/testdata/blist.hosts.src
            }
        }
        exclude 1e100.net
        exclude 2o7.net
        exclude adobedtm.com
        exclude akamai.net
        exclude akamaihd.net
        exclude amazon.com
        exclude amazonaws.com
        exclude apple.com
        exclude ask.com
        exclude avast.com
        exclude avira-update.com
        exclude bannerbank.com
        exclude bing.com
        exclude bit.ly
        exclude bitdefender.com
        exclude cdn.ravenjs.com
        exclude cdn.visiblemeasures.com
        exclude cloudfront.net
        exclude coremetrics.com
        exclude dropbox.com
        exclude ebay.com
        exclude edgesuite.net
        exclude evernote.com
        exclude express.co.uk
        exclude feedly.com
        exclude freedns.afraid.org
        exclude github.com
        exclude githubusercontent.com
        exclude global.ssl.fastly.net
        exclude google.com
        exclude googleads.g.doubleclick.net
        exclude googleadservices.com
        exclude googleapis.com
        exclude googletagmanager.com
        exclude googleusercontent.com
        exclude gstatic.com
        exclude gvt1.com
        exclude gvt1.net
        exclude hb.disney.go.com
        exclude herokuapp.com
        exclude hp.com
        exclude hulu.com
        exclude images-amazon.com
        exclude live.com
        exclude magnetmail1.net
        exclude microsoft.com
        exclude microsoftonline.com
        exclude msdn.com
        exclude msecnd.net
        exclude msftncsi.com
        exclude mywot.com
        exclude nsatc.net
        exclude paypal.com
        exclude pop.h-cdn.co
        exclude rackcdn.com
        exclude rarlab.com
        exclude schema.org
        exclude shopify.com
        exclude skype.com
        exclude smacargo.com
        exclude sourceforge.net
        exclude spotify.com
        exclude spotify.edgekey.net
        exclude spotilocal.com
        exclude ssl-on9.com
        exclude ssl-on9.net
        exclude sstatic.net
        exclude static.chartbeat.com
        exclude storage.googleapis.com
        exclude twimg.com
        exclude viewpoint.com
        exclude windows.net
        exclude xboxlive.com
        exclude yimg.com
        exclude ytimg.com
        hosts {
            include ads.feedly.com
            include beap.gemini.yahoo.com
            source githubSteveBlack {
                description "Blacklists adware and malware websites"
                prefix 0.0.0.0
                url https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts
            }
            source hostsfile.org {
                description "hostsfile.org bad hosts blacklist"
                prefix 127.0.0.1
                url http://www.hostsfile.org/Downloads/hosts.txt
            }
            source openphish {
                description "OpenPhish automatic phishing detection"
                prefix http
                url https://openphish.com/feed.txt
            }
            source sysctl.org {
                description "This hosts file is a merged collection of hosts from Cameleon"
                prefix 127.0.0.1
                url http://sysctl.org/cameleon/hosts
            }
        }
    }`

	// SingleSource is a single source EdgeOS configuration
	SingleSource = `interfaces {
    ethernet eth0 {
        address dhcp
        address dhcpv6
        description "External WAN"
        dhcp-options {
            default-route update
            default-route-distance 210
            name-server no-update
        }
        duplex auto
        speed auto
    }
    ethernet eth1 {
        address 192.168.150.1/24
        description "Internal LAN"
        duplex auto
        mtu 1500
        speed auto
    }
    ethernet eth2 {
        address 192.168.200.1/24
        description DMZ
        duplex auto
        mtu 1500
        speed auto
    }
    ethernet eth3 {
        duplex auto
        speed auto
    }
    ethernet eth4 {
        duplex auto
        speed auto
    }
    loopback lo {
    }
    switch switch0 {
        mtu 1500
    }
}
service {
    dhcp-server {
        disabled false
        hostfile-update enable
        shared-network-name LAN0 {
            authoritative enable
            subnet 192.168.150.0/24 {
                default-router 192.168.150.1
                dns-server 192.168.150.1
                domain-name er-x.local
                lease 86400
                start 192.168.150.100 {
                    stop 192.168.150.200
                }
            }
        }
    }
    dns {
        forwarding {
            blacklist {
                disabled false
                dns-redirect-ip 0.0.0.0
                hosts {
                    source yoyo {
                        description "Fully Qualified Domain Names only - no prefix to strip"
                        prefix "127.0.0.1 "
                        url http://pgl.yoyo.org/as/serverlist.php
                    }
                }
            }
            cache-size 150
            listen-on eth1
            listen-on eth2
            name-server 208.67.220.220
            options bogus-priv
            options domain=er-x.local
            options expand-hosts
            options listen-address=::1
            options listen-address=127.0.0.1
            options localise-queries
            options strict-order
        }
    }
    gui {
        https-port 443
        listen-address 192.168.150.1
    }
    ssh {
        disable-password-authentication
        port 22
        protocol-version v2
    }
}
system {
    conntrack {
        expect-table-size 4096
        hash-size 4096
        ignore {
            rule 10 {
                destination {
                    address 255.255.255.255
                }
            }
        }
        table-size 262144
    }
    domain-search {
        domain ashcreek.home
    }
    host-name er-x
    ip {
        override-hostname-ip 192.168.150.1
    }
    login {
        banner {
            post-login "\nWelcome to EdgeOS!\n"
            pre-login "\n\n\n\tWARNING *** WARNING *** WARNING *** WARNING *** WARNING\n\n\n\tWARNING: Criminal and civil penalties may be imposed for obtaining\n\tunauthorized access to this system or for causing intentional,\n\tunauthorized damage, deletion, alteration, or insertion of data.\n\tAny information stored, processed, or transmitted to this system\n\tmay be monitored, used, or disclosed by authorized personnel,\n\tincluding law enforcement. Email sysadmin@empirecreekcircle.com\n\tto gain access to this equipment if you need authorization.\n\n\n"
        }
        user nbnt {
            authentication {
                encrypted-password ****************
                plaintext-password ****************
                public-keys root@MacBook-Air.local {
                    key ****************
                    type ssh-rsa
                }
            }
            full-name Admin
            level admin
        }
    }
    name-server 127.0.0.1
    ntp {
        server 0.ubnt.pool.ntp.org {
        }
        server 1.ubnt.pool.ntp.org {
        }
    }
    package {
        repository wheezy {
            components "main contrib non-free"
            distribution wheezy
            url http://http.us.debian.org/debian/
        }
    }
    syslog {
        global {
            archive {
                files 10
                size 250
            }
            facility all {
                level notice
            }
            facility cron {
                level err
            }
            facility protocols {
                level debug
            }
        }
    }
    task-scheduler {
        task update_blacklists {
            executable {
                path /config/scripts/update-dnsmasq
            }
            interval 1d
        }
    }
    time-zone America/Los_Angeles
}
`
)
