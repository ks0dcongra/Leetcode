module github.com/halfrost/LeetCode-Go

go 1.19

replace github.com/halfrost/LeetCode-Go/structures => ./structures

replace github.com/halfrost/LeetCode-Go/template => ./template

replace github.com/halfrost/LeetCode-Go/ctl/util => ./ctl/util

replace github.com/halfrost/LeetCode-Go/ctl/models => ./ctl/models

require (
	github.com/BurntSushi/toml v1.2.0
	github.com/halfrost/LeetCode-Go/ctl/models v0.0.0-20220910225043-e3bb5aff34d0
	github.com/halfrost/LeetCode-Go/ctl/util v0.0.0-20220910225043-e3bb5aff34d0
	github.com/halfrost/LeetCode-Go/structures v0.0.0-20220910233101-aa0e2c897b18
	github.com/halfrost/LeetCode-Go/template v0.0.0-20220910233504-e2a72e6212ce
	github.com/mozillazg/request v0.8.0
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.7.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
