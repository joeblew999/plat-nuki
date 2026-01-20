module github.com/joeblew999/plat-nuki

go 1.22

// Local development - point to .src repos for version control
replace (
	github.com/nuki-io/go-nuki => ./.src/go-nuki
	github.com/nuki-io/nuki-cli => ./.src/nuki-cli
	github.com/qvest-digital/go-nuki => ./.src/go-nuki-ble
)
