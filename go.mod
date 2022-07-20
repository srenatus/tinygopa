module github.com/srenatus/tinygopa

go 1.18

require github.com/open-policy-agent/opa v0.42.2

require (
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// stub out metrics that need runtime's Memstats
replace github.com/rcrowley/go-metrics => ./replacements/github.com/rcrowley/go-metrics

// stub out anything using reflect, https://github.com/tinygo-org/tinygo/issues/2950
replace gopkg.in/yaml.v2 => ./replacements/gopkg.in/yaml.v2

replace github.com/ghodss/yaml => ./replacements/github.com/ghodss/yaml
